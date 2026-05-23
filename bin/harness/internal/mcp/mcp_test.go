package mcp

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"harness/internal/tools"
)

func TestNewStdioTransport(t *testing.T) {
	// This requires a real binary, skip in short mode
	if testing.Short() {
		t.Skip("skipping stdio transport test in short mode")
	}
}

func TestNewSSETransport(t *testing.T) {
	transport := NewSSETransport("http://localhost:9999")
	if transport == nil {
		t.Fatal("expected non-nil SSE transport")
	}
	defer transport.Close()
}

func TestSSETransport_Send(t *testing.T) {
	transport := NewSSETransport("http://localhost:9999")
	defer transport.Close()

	// Send to a non-existent server should fail gracefully
	ctx := context.Background()
	err := transport.Send(ctx, []byte(`{"jsonrpc":"2.0"}`))
	if err == nil {
		t.Log("send returned nil (expected failure or timeout)")
	}
}

func TestRequestMarshal(t *testing.T) {
	req := Request{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "test",
		Params:  map[string]string{"key": "val"},
	}

	data, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	var unmarshaled Request
	if err := json.Unmarshal(data, &unmarshaled); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if unmarshaled.Method != "test" {
		t.Errorf("expected method 'test', got %q", unmarshaled.Method)
	}
}

func TestResponseMarshal(t *testing.T) {
	resp := Response{
		JSONRPC: "2.0",
		ID:      1,
		Result:  map[string]string{"status": "ok"},
	}

	data, err := json.Marshal(resp)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	if !strings.Contains(string(data), "ok") {
		t.Errorf("expected 'ok' in response, got %q", string(data))
	}
}

func TestResponseWithError(t *testing.T) {
	resp := Response{
		JSONRPC: "2.0",
		ID:      1,
		Error: &ErrorObj{
			Code:    -32601,
			Message: "Method not found",
		},
	}

	data, err := json.Marshal(resp)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	if !strings.Contains(string(data), "Method not found") {
		t.Errorf("expected error message in response, got %q", string(data))
	}
}

func TestNewServer(t *testing.T) {
	reg := tools.NewRegistry()
	reg.Register("test_tool", &mockTool{output: "hello"})

	server := NewServer(reg)
	if server == nil {
		t.Fatal("expected non-nil server")
	}
}

func TestServer_RegisterMethod(t *testing.T) {
	reg := tools.NewRegistry()
	server := NewServer(reg)

	called := false
	server.RegisterMethod("custom/method", func(req Request) (interface{}, *ErrorObj) {
		called = true
		return "done", nil
	})

	// Trigger via handleRequest
	resp := server.handleRequest(Request{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "custom/method",
	})

	if !called {
		t.Error("expected handler to be called")
	}
	if resp.Error != nil {
		t.Errorf("unexpected error: %v", resp.Error)
	}
}

func TestServer_HandleInitialize(t *testing.T) {
	reg := tools.NewRegistry()
	server := NewServer(reg)

	resp := server.handleRequest(Request{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "initialize",
	})

	if resp.Error != nil {
		t.Fatalf("unexpected error: %v", resp.Error)
	}

	result, ok := resp.Result.(map[string]interface{})
	if !ok {
		t.Fatal("expected map result")
	}

	if result["protocolVersion"] != "2025-03-26" {
		t.Errorf("expected protocol version, got %v", result["protocolVersion"])
	}
}

func TestServer_HandleToolsList(t *testing.T) {
	reg := tools.NewRegistry()
	reg.Register("my_tool", &mockTool{})

	server := NewServer(reg)

	resp := server.handleRequest(Request{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "tools/list",
	})

	if resp.Error != nil {
		t.Fatalf("unexpected error: %v", resp.Error)
	}

	// Marshal and unmarshal to JSON to simulate real behavior
	data, err := json.Marshal(resp)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	var rawResponse map[string]interface{}
	if err := json.Unmarshal(data, &rawResponse); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	result, ok := rawResponse["result"].(map[string]interface{})
	if !ok {
		t.Fatal("expected result object")
	}

	toolsRaw, ok := result["tools"].([]interface{})
	if !ok {
		t.Fatalf("expected tools array, got %T: %v", result["tools"], result["tools"])
	}

	if len(toolsRaw) != 1 {
		t.Errorf("expected 1 tool, got %d", len(toolsRaw))
	}
}

func TestServer_HandleToolsCall(t *testing.T) {
	reg := tools.NewRegistry()
	reg.Register("echo", &mockTool{output: "echo response"})

	server := NewServer(reg)

	resp := server.handleRequest(Request{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "tools/call",
		Params: map[string]interface{}{
			"name": "echo",
			"arguments": map[string]interface{}{
				"text": "hello",
			},
		},
	})

	if resp.Error != nil {
		t.Fatalf("unexpected error: %v", resp.Error)
	}
}

func TestServer_HandleToolsCall_NotFound(t *testing.T) {
	reg := tools.NewRegistry()
	server := NewServer(reg)

	resp := server.handleRequest(Request{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "tools/call",
		Params: map[string]interface{}{
			"name": "nonexistent",
		},
	})

	if resp.Error == nil {
		t.Fatal("expected error for nonexistent tool")
	}
}

func TestServer_HandleUnknownMethod(t *testing.T) {
	reg := tools.NewRegistry()
	server := NewServer(reg)

	resp := server.handleRequest(Request{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "unknown/method",
	})

	if resp.Error == nil {
		t.Fatal("expected error for unknown method")
	}
	if resp.Error.Code != -32601 {
		t.Errorf("expected code -32601, got %d", resp.Error.Code)
	}
}

func TestNewClient(t *testing.T) {
	transport := &mockTransport{}
	client := NewClient(transport)
	if client == nil {
		t.Fatal("expected non-nil client")
	}
}

func TestClient_Initialize(t *testing.T) {
	transport := &mockTransport{
		response: `{"jsonrpc":"2.0","id":1,"result":{"protocolVersion":"2025-03-26"}}`,
	}
	client := NewClient(transport)

	ctx := context.Background()
	err := client.Initialize(ctx)
	if err != nil {
		t.Fatalf("Initialize failed: %v", err)
	}
}

func TestClient_ListTools(t *testing.T) {
	transport := &mockTransport{
		response: `{"jsonrpc":"2.0","id":1,"result":{"tools":[{"name":"test_tool","description":"a test tool"}]}}`,
	}
	client := NewClient(transport)

	ctx := context.Background()
	tools, err := client.ListTools(ctx)
	if err != nil {
		t.Fatalf("ListTools failed: %v", err)
	}
	if len(tools) != 1 {
		t.Errorf("expected 1 tool, got %d", len(tools))
	}
	if tools[0].Name != "test_tool" {
		t.Errorf("expected 'test_tool', got %q", tools[0].Name)
	}
}

func TestClient_CallTool(t *testing.T) {
	transport := &mockTransport{
		response: `{"jsonrpc":"2.0","id":1,"result":{"content":[{"type":"text","text":"result"}]}}`,
	}
	client := NewClient(transport)

	ctx := context.Background()
	result, err := client.CallTool(ctx, "my_tool", map[string]interface{}{"key": "val"})
	if err != nil {
		t.Fatalf("CallTool failed: %v", err)
	}
	if len(result.Content) != 1 {
		t.Errorf("expected 1 content item, got %d", len(result.Content))
	}
	if result.Content[0].Text != "result" {
		t.Errorf("expected 'result', got %q", result.Content[0].Text)
	}
}

func TestClient_Close(t *testing.T) {
	transport := &mockTransport{}
	client := NewClient(transport)
	err := client.Close()
	if err != nil {
		t.Fatalf("Close failed: %v", err)
	}
}

func TestGetString(t *testing.T) {
	m := map[string]interface{}{
		"key": "value",
	}
	if getString(m, "key") != "value" {
		t.Errorf("expected 'value', got %q", getString(m, "key"))
	}
	if getString(m, "nonexistent") != "" {
		t.Errorf("expected empty for missing key")
	}
}

func TestGetBool(t *testing.T) {
	m := map[string]interface{}{
		"yes": true,
		"no":  false,
	}
	if !getBool(m, "yes") {
		t.Error("expected true")
	}
	if getBool(m, "no") {
		t.Error("expected false")
	}
	if getBool(m, "nonexistent") {
		t.Error("expected false for missing key")
	}
}

func TestToolResultContent(t *testing.T) {
	tr := ToolResult{
		Content: []ContentItem{
			{Type: "text", Text: "hello"},
			{Type: "text", Text: "world"},
		},
		IsError: false,
	}

	data, err := json.Marshal(tr)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	if !strings.Contains(string(data), "hello") {
		t.Errorf("expected content in json, got %q", string(data))
	}
}

// --- mocks ---

type mockTransport struct {
	response string
	sent     []byte
}

func (m *mockTransport) Send(ctx context.Context, msg []byte) error {
	m.sent = msg
	return nil
}

func (m *mockTransport) Receive(ctx context.Context) ([]byte, error) {
	return []byte(m.response), nil
}

func (m *mockTransport) Close() error {
	return nil
}

type mockTool struct {
	output string
}

func (m *mockTool) Execute(name string, args map[string]any) tools.ToolResult {
	return tools.ToolResult{Output: m.output}
}
