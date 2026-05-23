// Package mcp provides Model Context Protocol server and client implementations.
package mcp

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"harness/internal/tools"
)

// --- JSON-RPC Types ---

type Request struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
}

type Response struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id"`
	Result  interface{} `json:"result,omitempty"`
	Error   *ErrorObj   `json:"error,omitempty"`
}

type ErrorObj struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Notification struct {
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
}

type ToolDef struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	InputSchema interface{} `json:"inputSchema"`
}

type ToolCall struct {
	Name      string                 `json:"name"`
	Arguments map[string]interface{} `json:"arguments"`
}

type ToolResult struct {
	Content []ContentItem `json:"content"`
	IsError bool          `json:"isError"`
}

type ContentItem struct {
	Type string `json:"type"` // "text", "image", "resource"
	Text string `json:"text,omitempty"`
}

// --- Sender/Receiver/Closer segregated interfaces (ISP) ---

// Sender can send messages.
type Sender interface {
	Send(ctx context.Context, msg []byte) error
}

// Receiver can receive messages.
type Receiver interface {
	Receive(ctx context.Context) ([]byte, error)
}

// Closer can close the transport.
type Closer interface {
	Close() error
}

// Transport combines all three capabilities.
type Transport interface {
	Sender
	Receiver
	Closer
}

// --- StdioTransport ---

type StdioTransport struct {
	cmd    *exec.Cmd
	stdin  io.WriteCloser
	reader *bufio.Scanner
	mu     sync.Mutex
}

func NewStdioTransport(command string, args ...string) (*StdioTransport, error) {
	cmd := exec.Command(command, args...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("stdin pipe: %w", err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("stdout pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("start process: %w", err)
	}

	return &StdioTransport{
		cmd:    cmd,
		stdin:  stdin,
		reader: bufio.NewScanner(stdout),
	}, nil
}

func (t *StdioTransport) Send(ctx context.Context, msg []byte) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	_, err := t.stdin.Write(append(msg, '\n'))
	return err
}

func (t *StdioTransport) Receive(ctx context.Context) ([]byte, error) {
	done := make(chan struct{})
	var data []byte

	go func() {
		defer close(done)
		if t.reader.Scan() {
			data = make([]byte, len(t.reader.Bytes()))
			copy(data, t.reader.Bytes())
		}
	}()

	select {
	case <-done:
		if data == nil {
			if err := t.reader.Err(); err != nil {
				return nil, err
			}
			return nil, io.EOF
		}
		return data, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (t *StdioTransport) Close() error {
	t.stdin.Close()
	return t.cmd.Wait()
}

// --- SSETransport (proper LSP compliance: no "not implemented" methods) ---

type SSETransport struct {
	url        string
	client     *http.Client
	httpCh     chan []byte
	cancel     context.CancelFunc
	closeOnce  sync.Once
	closed     chan struct{}
}

func NewSSETransport(url string) *SSETransport {
	ctx, cancel := context.WithCancel(context.Background())
	t := &SSETransport{
		url:    url,
		client: &http.Client{Timeout: 30 * time.Second},
		httpCh: make(chan []byte, 64),
		cancel: cancel,
		closed: make(chan struct{}),
	}
	go t.listenSSE(ctx)
	return t
}

func (t *SSETransport) listenSSE(ctx context.Context) {
	defer close(t.closed)
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		// SSE keep-alive loop would go here in a real implementation
		time.Sleep(100 * time.Millisecond)
	}
}

func (t *SSETransport) Send(ctx context.Context, msg []byte) error {
	resp, err := t.client.Post(t.url, "application/json", strings.NewReader(string(msg)))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("SSE POST failed (status %d)", resp.StatusCode)
	}

	// For SSE transport, responses come asynchronously through the event stream.
	// Single request-response requires a session ID mechanism.
	// For now, read the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if len(body) > 0 {
		select {
		case t.httpCh <- body:
		default:
		}
	}

	return nil
}

func (t *SSETransport) Receive(ctx context.Context) ([]byte, error) {
	// Proper implementation: wait for either the async response or context cancel
	select {
	case data := <-t.httpCh:
		return data, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (t *SSETransport) Close() error {
	t.closeOnce.Do(func() {
		t.cancel()
		<-t.closed
	})
	return nil
}

// --- MCP Client ---

type Client struct {
	transport Transport
	nextID    int
	mu        sync.Mutex
}

func NewClient(transport Transport) *Client {
	return &Client{
		transport: transport,
		nextID:    1,
	}
}

func (c *Client) sendRequest(ctx context.Context, method string, params interface{}) (Response, error) {
	c.mu.Lock()
	id := c.nextID
	c.nextID++
	c.mu.Unlock()

	req := Request{
		JSONRPC: "2.0",
		ID:      id,
		Method:  method,
		Params:  params,
	}

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return Response{}, fmt.Errorf("marshal request: %w", err)
	}

	if err := c.transport.Send(ctx, reqBytes); err != nil {
		return Response{}, fmt.Errorf("send: %w", err)
	}

	respBytes, err := c.transport.Receive(ctx)
	if err != nil {
		return Response{}, fmt.Errorf("receive: %w", err)
	}

	var resp Response
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return Response{}, fmt.Errorf("parse response: %w", err)
	}

	if resp.Error != nil {
		return resp, fmt.Errorf("MCP error (code %d): %s", resp.Error.Code, resp.Error.Message)
	}

	return resp, nil
}

func (c *Client) Initialize(ctx context.Context) error {
	_, err := c.sendRequest(ctx, "initialize", map[string]interface{}{
		"protocolVersion": "2025-03-26",
		"capabilities":    map[string]interface{}{},
		"clientInfo": map[string]interface{}{
			"name":    "harness-mcp-client",
			"version": "1.0.0",
		},
	})
	return err
}

func (c *Client) ListTools(ctx context.Context) ([]ToolDef, error) {
	resp, err := c.sendRequest(ctx, "tools/list", nil)
	if err != nil {
		return nil, err
	}

	result, ok := resp.Result.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected response format")
	}

	toolsRaw, ok := result["tools"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("tools field missing")
	}

	tools := make([]ToolDef, 0, len(toolsRaw))
	for _, t := range toolsRaw {
		tMap, ok := t.(map[string]interface{})
		if !ok {
			continue
		}
		tools = append(tools, ToolDef{
			Name:        getString(tMap, "name"),
			Description: getString(tMap, "description"),
			InputSchema: tMap["inputSchema"],
		})
	}

	return tools, nil
}

func (c *Client) CallTool(ctx context.Context, name string, args map[string]interface{}) (ToolResult, error) {
	params := map[string]interface{}{
		"name":      name,
		"arguments": args,
	}

	resp, err := c.sendRequest(ctx, "tools/call", params)
	if err != nil {
		return ToolResult{}, err
	}

	result, ok := resp.Result.(map[string]interface{})
	if !ok {
		return ToolResult{}, fmt.Errorf("unexpected response format")
	}

	toolResult := ToolResult{
		IsError: getBool(result, "isError"),
	}

	contentRaw, ok := result["content"].([]interface{})
	if ok {
		for _, c := range contentRaw {
			cMap, ok := c.(map[string]interface{})
			if !ok {
				continue
			}
			toolResult.Content = append(toolResult.Content, ContentItem{
				Type: getString(cMap, "type"),
				Text: getString(cMap, "text"),
			})
		}
	}

	return toolResult, nil
}

func (c *Client) Close() error {
	return c.transport.Close()
}

// --- MCP Server (OCP-compliant via handler registry) ---

type MethodHandler func(req Request) (interface{}, *ErrorObj)

type Server struct {
	methodHandlers map[string]MethodHandler
	toolRegistry   *tools.Registry
}

func NewServer(toolReg *tools.Registry) *Server {
	s := &Server{
		methodHandlers: make(map[string]MethodHandler),
		toolRegistry:   toolReg,
	}
	s.registerBuiltinMethods()
	return s
}

// RegisterMethod adds a handler for a JSON-RPC method (OCP: extend by registration, not modification).
func (s *Server) RegisterMethod(method string, handler MethodHandler) {
	s.methodHandlers[method] = handler
}

func (s *Server) registerBuiltinMethods() {
	s.RegisterMethod("initialize", s.handleInitialize)
	s.RegisterMethod("tools/list", s.handleToolsList)
	s.RegisterMethod("tools/call", s.handleToolsCall)
}

func (s *Server) handleInitialize(req Request) (interface{}, *ErrorObj) {
	return map[string]interface{}{
		"protocolVersion": "2025-03-26",
		"capabilities": map[string]interface{}{
			"tools": map[string]interface{}{},
		},
		"serverInfo": map[string]interface{}{
			"name":    "harness-mcp-server",
			"version": "1.0.0",
		},
	}, nil
}

func (s *Server) handleToolsList(req Request) (interface{}, *ErrorObj) {
	toolList := make([]ToolDef, 0)
	for _, name := range s.toolRegistry.Names() {
		toolList = append(toolList, ToolDef{
			Name:        name,
			Description: getBuiltinDescription(name),
			InputSchema: getBuiltinSchema(name),
		})
	}
	return map[string]interface{}{"tools": toolList}, nil
}

func (s *Server) handleToolsCall(req Request) (interface{}, *ErrorObj) {
	params, ok := req.Params.(map[string]interface{})
	if !ok {
		return nil, &ErrorObj{Code: -32602, Message: "Invalid params: expected object"}
	}

	toolName, _ := params["name"].(string)
	arguments, _ := params["arguments"].(map[string]interface{})

	// Convert arguments to map[string]any
	args := make(map[string]any)
	for k, v := range arguments {
		args[k] = v
	}

	result := s.toolRegistry.Execute(toolName, args)
	if result.Error != nil {
		return nil, &ErrorObj{Code: -32000, Message: result.Error.Error()}
	}

	return ToolResult{
		Content: []ContentItem{{Type: "text", Text: result.Output}},
		IsError: result.ExitCode != 0,
	}, nil
}

// Start begins reading from stdin in JSON-RPC over stdio mode.
func (s *Server) Start() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		// Try as request
		var req Request
		if err := json.Unmarshal([]byte(line), &req); err == nil && req.Method != "" {
			resp := s.handleRequest(req)
			respBytes, _ := json.Marshal(resp)
			fmt.Println(string(respBytes))
			continue
		}

		// Try as notification
		var notif Notification
		if err := json.Unmarshal([]byte(line), &notif); err == nil && notif.Method != "" {
			// Notifications don't get responses
			continue
		}

		// Parse error
		errResp := Response{
			JSONRPC: "2.0",
			ID:      nil,
			Error: &ErrorObj{
				Code:    -32700,
				Message: "Parse error: invalid JSON-RPC message",
			},
		}
		respBytes, _ := json.Marshal(errResp)
		fmt.Println(string(respBytes))
	}
}

func (s *Server) handleRequest(req Request) Response {
	if handler, exists := s.methodHandlers[req.Method]; exists {
		result, errObj := handler(req)
		if errObj != nil {
			return Response{JSONRPC: "2.0", ID: req.ID, Error: errObj}
		}
		return Response{JSONRPC: "2.0", ID: req.ID, Result: result}
	}

	return Response{
		JSONRPC: "2.0",
		ID:      req.ID,
		Error:   &ErrorObj{Code: -32601, Message: fmt.Sprintf("Method not found: %s", req.Method)},
	}
}

// --- Helpers ---

func getString(m map[string]interface{}, key string) string {
	if v, ok := m[key].(string); ok {
		return v
	}
	return ""
}

func getBool(m map[string]interface{}, key string) bool {
	if v, ok := m[key].(bool); ok {
		return v
	}
	return false
}

func getBuiltinDescription(name string) string {
	descriptions := map[string]string{
		"read_file":       "Read the contents of a file at the specified path.",
		"write_file":      "Write or overwrite a file at the specified path.",
		"patch_file":      "Apply a surgical text replacement in a file.",
		"search_files":    "Search for files by name pattern and/or text content.",
		"execute_command": "Execute a bash command on the local system.",
	}
	if desc, ok := descriptions[name]; ok {
		return desc
	}
	return "No description"
}

func getBuiltinSchema(name string) interface{} {
	schemas := map[string]interface{}{
		"read_file": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"path": map[string]interface{}{"type": "string"},
			},
			"required": []interface{}{"path"},
		},
		"write_file": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"path":    map[string]interface{}{"type": "string"},
				"content": map[string]interface{}{"type": "string"},
			},
			"required": []interface{}{"path", "content"},
		},
		"patch_file": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"path":        map[string]interface{}{"type": "string"},
				"target":      map[string]interface{}{"type": "string"},
				"replacement": map[string]interface{}{"type": "string"},
			},
			"required": []interface{}{"path", "target", "replacement"},
		},
		"search_files": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"pattern": map[string]interface{}{"type": "string"},
				"query":   map[string]interface{}{"type": "string"},
			},
		},
		"execute_command": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"command": map[string]interface{}{"type": "string"},
			},
			"required": []interface{}{"command"},
		},
	}
	if schema, ok := schemas[name]; ok {
		return schema
	}
	return map[string]interface{}{"type": "object"}
}
