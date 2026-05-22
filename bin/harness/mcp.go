
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

// ---------------------------------------------------------------------------
// MCP Protocol Types (JSON-RPC 2.0)
// ---------------------------------------------------------------------------

type MCPRequest struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
}

type MCPResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id"`
	Result  interface{} `json:"result,omitempty"`
	Error   *MCPError   `json:"error,omitempty"`
}

type MCPError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type MCPNotification struct {
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
}

type MCPTool struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	InputSchema interface{} `json:"inputSchema"`
}

type MCPToolCall struct {
	Name       string                 `json:"name"`
	Arguments  map[string]interface{} `json:"arguments"`
}

type MCPToolResult struct {
	Content []MCPContent `json:"content"`
	IsError bool         `json:"isError"`
}

type MCPContent struct {
	Type string `json:"type"` // "text", "image", "resource"
	Text string `json:"text,omitempty"`
}

// ---------------------------------------------------------------------------
// MCP Server (stdio transport)
// ---------------------------------------------------------------------------

type MCPServer struct {
	tools map[string]ToolHandler
}

type ToolHandler func(args map[string]interface{}) (MCPToolResult, error)

func NewMCPServer() *MCPServer {
	return &MCPServer{
		tools: make(map[string]ToolHandler),
	}
}

func (s *MCPServer) RegisterTool(name string, handler ToolHandler) {
	s.tools[name] = handler
}

func (s *MCPServer) Start() {
	scanner := bufio.NewScanner(os.Stdin)
	// MCP uses length-prefixed JSON over stdin/stdout
	// Simplified: read each line as a JSON message (works with \n-delimited JSON)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		// Try as request first
		var req MCPRequest
		if err := json.Unmarshal([]byte(line), &req); err == nil && req.Method != "" {
			resp := s.handleRequest(req)
			respBytes, _ := json.Marshal(resp)
			fmt.Println(string(respBytes))
			continue
		}

		// Try as notification
		var notif MCPNotification
		if err := json.Unmarshal([]byte(line), &notif); err == nil && notif.Method != "" {
			s.handleNotification(notif)
			continue
		}

		// Invalid message
		errResp := MCPResponse{
			JSONRPC: "2.0",
			ID:      nil,
			Error: &MCPError{
				Code:    -32700,
				Message: "Parse error: invalid JSON-RPC message",
			},
		}
		respBytes, _ := json.Marshal(errResp)
		fmt.Println(string(respBytes))
	}
}

func (s *MCPServer) handleRequest(req MCPRequest) MCPResponse {
	var result interface{}
	var err *MCPError

	switch req.Method {
	case "initialize":
		result = map[string]interface{}{
			"protocolVersion": "2025-03-26",
			"capabilities": map[string]interface{}{
				"tools": map[string]interface{}{},
			},
			"serverInfo": map[string]interface{}{
				"name":    "harness-mcp-server",
				"version": "1.0.0",
			},
		}

	case "tools/list":
		toolList := make([]MCPTool, 0, len(s.tools))
		for name, handler := range s.tools {
			_ = handler // just to use the variable
			toolList = append(toolList, MCPTool{
				Name:        name,
				Description: getToolDescription(name),
				InputSchema: getToolInputSchema(name),
			})
		}
		result = map[string]interface{}{"tools": toolList}

	case "tools/call":
		params, ok := req.Params.(map[string]interface{})
		if !ok {
			err = &MCPError{Code: -32602, Message: "Invalid params: expected object"}
			break
		}
		toolName, _ := params["name"].(string)
		arguments, _ := params["arguments"].(map[string]interface{})

		handler, exists := s.tools[toolName]
		if !exists {
			err = &MCPError{Code: -32601, Message: fmt.Sprintf("Tool not found: %s", toolName)}
			break
		}

		toolResult, handlerErr := handler(arguments)
		if handlerErr != nil {
			err = &MCPError{Code: -32000, Message: handlerErr.Error()}
			break
		}
		result = toolResult

	case "notifications/initialized":
		// No response needed for notifications, but since we handle here as request, just return empty
		return MCPResponse{JSONRPC: "2.0", ID: req.ID, Result: nil}

	default:
		err = &MCPError{Code: -32601, Message: fmt.Sprintf("Method not found: %s", req.Method)}
	}

	if err != nil {
		return MCPResponse{JSONRPC: "2.0", ID: req.ID, Error: err}
	}
	return MCPResponse{JSONRPC: "2.0", ID: req.ID, Result: result}
}

func (s *MCPServer) handleNotification(notif MCPNotification) {
	switch notif.Method {
	case "notifications/initialized":
		// Handshake complete, no action needed
	}
}

// ---------------------------------------------------------------------------
// MCP Client (stdio and SSE transports)
// ---------------------------------------------------------------------------

type MCPClient struct {
	transport MCPTransport
	nextID    int
	mu        sync.Mutex
}

type MCPTransport interface {
	Send(message []byte) error
	Receive() ([]byte, error)
	Close() error
}

// StdioTransport connects via subprocess stdin/stdout
type StdioTransport struct {
	cmd    *exec.Cmd
	stdin  io.WriteCloser
	stdout io.ReadCloser
	reader *bufio.Scanner
}

func NewStdioTransport(command string, args ...string) (*StdioTransport, error) {
	cmd := exec.Command(command, args...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stdin pipe: %w", err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stdout pipe: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start MCP server process: %w", err)
	}

	return &StdioTransport{
		cmd:    cmd,
		stdin:  stdin,
		stdout: stdout,
		reader: bufio.NewScanner(stdout),
	}, nil
}

func (t *StdioTransport) Send(message []byte) error {
	_, err := t.stdin.Write(append(message, '\n'))
	return err
}

func (t *StdioTransport) Receive() ([]byte, error) {
	if t.reader.Scan() {
		return t.reader.Bytes(), nil
	}
	if err := t.reader.Err(); err != nil {
		return nil, fmt.Errorf("MCP transport read error: %w", err)
	}
	return nil, io.EOF
}

func (t *StdioTransport) Close() error {
	t.stdin.Close()
	return t.cmd.Wait()
}

// SSETransport uses HTTP POST + SSE for remote MCP servers
type SSETransport struct {
	url    string
	client *http.Client
}

func NewSSETransport(url string) *SSETransport {
	return &SSETransport{
		url:    url,
		client: &http.Client{Timeout: 30 * time.Second},
	}
}

func (t *SSETransport) Send(message []byte) error {
	resp, err := t.client.Post(t.url, "application/json", strings.NewReader(string(message)))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// For SSE transport, response is typically an event stream, but we handle only single request for now
	return nil
}

func (t *SSETransport) Receive() ([]byte, error) {
	return nil, fmt.Errorf("SSE transport receive not implemented for non-streaming mode")
}

func (t *SSETransport) Close() error {
	return nil
}

// NewMCPClient creates an MCP client with the given transport
func NewMCPClient(transport MCPTransport) *MCPClient {
	return &MCPClient{
		transport: transport,
		nextID:    1,
	}
}

func (c *MCPClient) sendRequest(method string, params interface{}) (MCPResponse, error) {
	c.mu.Lock()
	id := c.nextID
	c.nextID++
	c.mu.Unlock()

	req := MCPRequest{
		JSONRPC: "2.0",
		ID:      id,
		Method:  method,
		Params:  params,
	}

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return MCPResponse{}, fmt.Errorf("failed to marshal request: %w", err)
	}

	if err := c.transport.Send(reqBytes); err != nil {
		return MCPResponse{}, fmt.Errorf("failed to send request: %w", err)
	}

	respBytes, err := c.transport.Receive()
	if err != nil {
		return MCPResponse{}, fmt.Errorf("failed to receive response: %w", err)
	}

	var resp MCPResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return MCPResponse{}, fmt.Errorf("failed to parse response: %w", err)
	}

	if resp.Error != nil {
		return resp, fmt.Errorf("MCP error (code %d): %s", resp.Error.Code, resp.Error.Message)
	}

	return resp, nil
}

// ListTools requests available tools from the MCP server
func (c *MCPClient) ListTools() ([]MCPTool, error) {
	resp, err := c.sendRequest("tools/list", nil)
	if err != nil {
		return nil, err
	}

	result, ok := resp.Result.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected response format for tools/list")
	}

	toolsRaw, ok := result["tools"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("tools field missing or not an array")
	}

	tools := make([]MCPTool, 0, len(toolsRaw))
	for _, t := range toolsRaw {
		tMap, ok := t.(map[string]interface{})
		if !ok {
			continue
		}
		tool := MCPTool{
			Name:        getString(tMap, "name"),
			Description: getString(tMap, "description"),
			InputSchema: tMap["inputSchema"],
		}
		tools = append(tools, tool)
	}

	return tools, nil
}

// CallTool invokes a tool on the MCP server
func (c *MCPClient) CallTool(name string, args map[string]interface{}) (MCPToolResult, error) {
	params := map[string]interface{}{
		"name":      name,
		"arguments": args,
	}

	resp, err := c.sendRequest("tools/call", params)
	if err != nil {
		return MCPToolResult{}, err
	}

	result, ok := resp.Result.(map[string]interface{})
	if !ok {
		return MCPToolResult{}, fmt.Errorf("unexpected response format for tools/call")
	}

	toolResult := MCPToolResult{
		IsError: getBool(result, "isError"),
	}

	contentRaw, ok := result["content"].([]interface{})
	if ok {
		for _, c := range contentRaw {
			cMap, ok := c.(map[string]interface{})
			if !ok {
				continue
			}
			content := MCPContent{
				Type: getString(cMap, "type"),
				Text: getString(cMap, "text"),
			}
			toolResult.Content = append(toolResult.Content, content)
		}
	}

	return toolResult, nil
}

// Close closes the MCP client transport
func (c *MCPClient) Close() error {
	return c.transport.Close()
}

// ---------------------------------------------------------------------------
// Helper functions
// ---------------------------------------------------------------------------

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

// getToolDescription returns a description for native Harness tools
func getToolDescription(name string) string {
	descriptions := map[string]string{
		"read_file":     "Read the contents of a file at the specified path. Returns file content as text. Respects workspace security boundaries.",
		"write_file":    "Write or overwrite a file at the specified path with the given content. Creates parent directories if needed.",
		"patch_file":    "Apply a surgical text replacement in a file. Finds an exact target block and replaces it with new content. Requires unique match.",
		"search_files":  "Search for files by name pattern and/or text content. Returns file paths with line numbers for matches.",
		"execute_command": "Execute a bash command on the local system and return its output. Runs in the workspace root directory.",
	}
	if desc, ok := descriptions[name]; ok {
		return desc
	}
	return "No description available"
}

// getToolInputSchema returns JSON Schema for each native tool
func getToolInputSchema(name string) interface{} {
	schemas := map[string]interface{}{
		"read_file": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"path": map[string]interface{}{"type": "string", "description": "Relative or absolute path to the file"},
			},
			"required": []interface{}{"path"},
		},
		"write_file": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"path":    map[string]interface{}{"type": "string", "description": "Path to write the file"},
				"content": map[string]interface{}{"type": "string", "description": "Content to write to the file"},
			},
			"required": []interface{}{"path", "content"},
		},
		"patch_file": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"path":        map[string]interface{}{"type": "string", "description": "Path to the file to patch"},
				"target":      map[string]interface{}{"type": "string", "description": "Exact text block to find (must be unique)"},
				"replacement": map[string]interface{}{"type": "string", "description": "Replacement text block"},
			},
			"required": []interface{}{"path", "target", "replacement"},
		},
		"search_files": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"pattern": map[string]interface{}{"type": "string", "description": "Glob pattern for filename (e.g., *.go)"},
				"query":   map[string]interface{}{"type": "string", "description": "Text to search within files"},
			},
			"required": []interface{}{},
		},
		"execute_command": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"command": map[string]interface{}{"type": "string", "description": "Bash command to execute"},
			},
			"required": []interface{}{"command"},
		},
	}
	if schema, ok := schemas[name]; ok {
		return schema
	}
	return map[string]interface{}{"type": "object"}
}

// runMCPServer is the entry point for `harness mcp-server`
func runMCPServer() {
	server := NewMCPServer()

	// Register native Harness tools as MCP tools
	server.RegisterTool("read_file", func(args map[string]interface{}) (MCPToolResult, error) {
		path := getString(args, "path")
		if path == "" {
			return MCPToolResult{IsError: true, Content: []MCPContent{{Type: "text", Text: "Missing required argument: path"}}}, nil
		}
		content, err := ReadFile(path)
		if err != nil {
			return MCPToolResult{IsError: true, Content: []MCPContent{{Type: "text", Text: err.Error()}}}, nil
		}
		return MCPToolResult{Content: []MCPContent{{Type: "text", Text: content}}}, nil
	})

	server.RegisterTool("write_file", func(args map[string]interface{}) (MCPToolResult, error) {
		path := getString(args, "path")
		content := getString(args, "content")
		if path == "" {
			return MCPToolResult{IsError: true, Content: []MCPContent{{Type: "text", Text: "Missing required argument: path"}}}, nil
		}
		if err := WriteFile(path, content); err != nil {
			return MCPToolResult{IsError: true, Content: []MCPContent{{Type: "text", Text: err.Error()}}}, nil
		}
		return MCPToolResult{Content: []MCPContent{{Type: "text", Text: "File written successfully."}}}, nil
	})

	server.RegisterTool("patch_file", func(args map[string]interface{}) (MCPToolResult, error) {
		path := getString(args, "path")
		target := getString(args, "target")
		replacement := getString(args, "replacement")
		if path == "" || target == "" {
			return MCPToolResult{IsError: true, Content: []MCPContent{{Type: "text", Text: "Missing required arguments: path, target, replacement"}}}, nil
		}
		if err := PatchFile(path, target, replacement); err != nil {
			return MCPToolResult{IsError: true, Content: []MCPContent{{Type: "text", Text: err.Error()}}}, nil
		}
		return MCPToolResult{Content: []MCPContent{{Type: "text", Text: "Patch applied successfully."}}}, nil
	})

	server.RegisterTool("search_files", func(args map[string]interface{}) (MCPToolResult, error) {
		pattern := getString(args, "pattern")
		query := getString(args, "query")
		result, err := SearchFiles(pattern, query)
		if err != nil {
			return MCPToolResult{IsError: true, Content: []MCPContent{{Type: "text", Text: err.Error()}}}, nil
		}
		return MCPToolResult{Content: []MCPContent{{Type: "text", Text: result}}}, nil
	})

	server.RegisterTool("execute_command", func(args map[string]interface{}) (MCPToolResult, error) {
		command := getString(args, "command")
		if command == "" {
			return MCPToolResult{IsError: true, Content: []MCPContent{{Type: "text", Text: "Missing required argument: command"}}}, nil
		}
		out, code, err := ExecuteCommand(command)
		if err != nil || code != 0 {
			errMsg := fmt.Sprintf("Command failed (exit code %d): %s\n%s", code, err, out)
			return MCPToolResult{IsError: true, Content: []MCPContent{{Type: "text", Text: errMsg}}}, nil
		}
		return MCPToolResult{Content: []MCPContent{{Type: "text", Text: out}}}, nil
	})

	fmt.Fprintf(os.Stderr, "Harness MCP Server started (stdio transport)\n")
	server.Start()
}

// runMCPClient connects to an external MCP server and interactively lists/calls tools
func runMCPClient(transportType string, address string) {
	var transport MCPTransport
	var err error

	switch transportType {
	case "stdio":
		parts := strings.Split(address, " ")
		cmd := parts[0]
		args := parts[1:]
		transport, err = NewStdioTransport(cmd, args...)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to connect to MCP server: %v\n", err)
			return
		}
	case "sse":
		transport = NewSSETransport(address)
	default:
		fmt.Fprintf(os.Stderr, "Unknown transport type: %s (use 'stdio' or 'sse')\n", transportType)
		return
	}

	client := NewMCPClient(transport)
	defer client.Close()

	// Initialize handshake
	_, err = client.sendRequest("initialize", map[string]interface{}{
		"protocolVersion": "2025-03-26",
		"capabilities":    map[string]interface{}{},
		"clientInfo": map[string]interface{}{
			"name":    "harness-mcp-client",
			"version": "1.0.0",
		},
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "MCP initialization failed: %v\n", err)
		return
	}

	// List tools
	tools, err := client.ListTools()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to list tools: %v\n", err)
		return
	}

	fmt.Printf("\n🧰 MCP Server Tools (%d):\n", len(tools))
	for _, tool := range tools {
		fmt.Printf("  • \033[32m%s\033[0m: %s\n", tool.Name, tool.Description)
	}
	fmt.Println()

	// Interactive tool call loop
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("🔧 MCP Client Interactive Shell")
	fmt.Println("Commands: /list, /call <toolname> <json-args>, /exit")
	for {
		fmt.Print("mcp-client ❯ ")
		if !scanner.Scan() {
			break
		}
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, " ", 3)
		switch parts[0] {
		case "/exit", "/quit":
			return
		case "/list":
			tools, err := client.ListTools()
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}
			for _, tool := range tools {
				fmt.Printf("  • \033[32m%s\033[0m: %s\n", tool.Name, tool.Description)
			}
		case "/call":
			if len(parts) < 3 {
				fmt.Println("Usage: /call <toolname> <json-args>")
				continue
			}
			toolName := parts[1]
			var args map[string]interface{}
			if err := json.Unmarshal([]byte(parts[2]), &args); err != nil {
				fmt.Printf("Invalid JSON args: %v\n", err)
				continue
			}
			result, err := client.CallTool(toolName, args)
			if err != nil {
				fmt.Printf("Error calling tool: %v\n", err)
				continue
			}
			for _, content := range result.Content {
				if result.IsError {
					fmt.Printf("\033[31m%s\033[0m\n", content.Text)
				} else {
					fmt.Println(content.Text)
				}
			}
		default:
			fmt.Println("Unknown command. Use /list, /call, /exit")
		}
	}
}
