// Package errutil provides error categorization and recovery for the Harness agent.
package ui

import (
	"fmt"
	"strings"
	"time"
)

// Category classifies an error for appropriate recovery strategy.
type Category int

const (
	CategoryUnknown   Category = iota
	CategoryNetwork            // Connection refused, DNS failure, timeout
	CategoryAuth               // 401, 403, invalid API key
	CategoryRateLimit          // 429, quota exceeded
	CategoryServer             // 5xx, server errors
	CategoryTool               // Tool execution failed
	CategoryParse              // JSON parse, XML parse errors
	CategoryContext            // Context canceled, deadline exceeded
)

// String returns a human-readable category name.
func (c Category) String() string {
	switch c {
	case CategoryNetwork:
		return "network"
	case CategoryAuth:
		return "auth"
	case CategoryRateLimit:
		return "rate-limit"
	case CategoryServer:
		return "server"
	case CategoryTool:
		return "tool"
	case CategoryParse:
		return "parse"
	case CategoryContext:
		return "context"
	default:
		return "unknown"
	}
}

// ErrorInfo holds detailed information about an error for recovery.
type ErrorInfo struct {
	Original     error
	Category     Category
	Message      string
	Recoverable  bool
	Suggestions  []string
	RetryAfter   time.Duration // Suggested delay before retry
	ProviderName string        // Which provider failed (if applicable)
}

// ClassifyError analyzes an error and returns structured info.
func ClassifyError(err error, providerName string) ErrorInfo {
	if err == nil {
		return ErrorInfo{Recoverable: true}
	}

	msg := err.Error()
	info := ErrorInfo{
		Original:     err,
		ProviderName: providerName,
	}

	switch {
	case strings.Contains(msg, "401") || strings.Contains(msg, "API key") || strings.Contains(msg, "Unauthorized") || strings.Contains(msg, "not set"):
		info.Category = CategoryAuth
		info.Recoverable = false
		info.Message = "Erro de autenticação — chave de API inválida ou não configurada"
		info.Suggestions = []string{
			"Verifique sua chave de API no arquivo .env",
			"Execute 'harness init' para configurar",
			"Verifique se a variável de ambiente está definida",
		}

	case strings.Contains(msg, "429") || strings.Contains(msg, "quota") || strings.Contains(msg, "rate limit") || strings.Contains(msg, "too many"):
		info.Category = CategoryRateLimit
		info.Recoverable = true
		info.RetryAfter = 5 * time.Second
		info.Message = "Limite de taxa excedido — aguardando para tentar novamente"
		info.Suggestions = []string{
			"Reduza a frequência das requisições",
			"Ative o fallback para outro provedor",
			"Verifique seu plano de API",
		}

	case strings.Contains(msg, "timeout") || strings.Contains(msg, "deadline") || strings.Contains(msg, "Client.Timeout") || strings.Contains(msg, "context deadline"):
		info.Category = CategoryContext
		info.Recoverable = true
		info.RetryAfter = 1 * time.Second
		info.Message = "Tempo limite excedido — a requisição demorou muito"
		info.Suggestions = []string{
			"Tente novamente com uma consulta mais simples",
			"Ative o streaming para respostas parciais mais rápidas",
			"Verifique sua conexão de internet",
		}

	case strings.Contains(msg, "connection") || strings.Contains(msg, "no such host") || strings.Contains(msg, "refused") || strings.Contains(msg, "EOF"):
		info.Category = CategoryNetwork
		info.Recoverable = true
		info.RetryAfter = 2 * time.Second
		info.Message = "Erro de conexão de rede"
		info.Suggestions = []string{
			"Verifique sua conexão com a internet",
			"Verifique se a API está acessível",
			"Se for Ollama local, verifique se o serviço está rodando",
		}

	case strings.Contains(msg, "5") && len(msg) > 0 && msg[0] == '5':
		info.Category = CategoryServer
		info.Recoverable = true
		info.RetryAfter = 3 * time.Second
		info.Message = "Erro no servidor da API — temporário"
		info.Suggestions = []string{
			"Tente novamente em alguns segundos",
			"Ative o fallback automático entre provedores",
		}

	case strings.Contains(msg, "parse") || strings.Contains(msg, "unmarshal") || strings.Contains(msg, "unexpected"):
		info.Category = CategoryParse
		info.Recoverable = false
		info.Message = "Erro ao processar resposta da API"
		info.Suggestions = []string{
			"Pode ser um problema de formato na resposta",
			"Tente novamente — pode ser transitório",
		}

	case strings.Contains(msg, "tool") && strings.Contains(msg, "not found"):
		info.Category = CategoryTool
		info.Recoverable = false
		info.Message = "Ferramenta não encontrada ou inválida"
		info.Suggestions = []string{
			"Verifique se a ferramenta está registrada",
			"Use /skills para listar ferramentas disponíveis",
		}

	default:
		info.Category = CategoryUnknown
		info.Recoverable = true
		info.RetryAfter = 1 * time.Second
		info.Message = fmt.Sprintf("Erro inesperado: %s", msg)
		info.Suggestions = []string{
			"Tente novamente",
			"Se o erro persistir, verifique os logs",
		}
	}

	return info
}

// RecoveryAction defines what to do after an error.
type RecoveryAction int

const (
	ActionRetry    RecoveryAction = iota // Retry the same operation
	ActionFallback                        // Try a different provider/tool
	ActionAbort                           // Stop and show error
	ActionContinue                        // Skip and continue
)

// DecideRecovery determines the best recovery action based on error info and attempt count.
func DecideRecovery(info ErrorInfo, attempt int, maxAttempts int) (RecoveryAction, time.Duration) {
	if !info.Recoverable {
		return ActionAbort, 0
	}

	if attempt >= maxAttempts {
		return ActionFallback, 0
	}

	switch info.Category {
	case CategoryRateLimit:
		delay := info.RetryAfter * time.Duration(attempt+1)
		if delay > 30*time.Second {
			delay = 30 * time.Second
		}
		return ActionRetry, delay

	case CategoryNetwork, CategoryServer:
		delay := info.RetryAfter * time.Duration(attempt+1)
		if delay > 15*time.Second {
			delay = 15 * time.Second
		}
		return ActionRetry, delay

	case CategoryContext:
		return ActionRetry, info.RetryAfter

	default:
		if attempt < 2 {
			return ActionRetry, 500 * time.Millisecond
		}
		return ActionFallback, 0
	}
}

// FormatErrorForDisplay returns a human-readable error message with recovery suggestions.
func FormatErrorForDisplay(info ErrorInfo) string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("\n\033[38;5;196m❌ %s\033[0m", info.Message))

	if info.ProviderName != "" {
		b.WriteString(fmt.Sprintf(" \033[90m[%s]\033[0m", strings.ToUpper(info.ProviderName)))
	}

	b.WriteString(fmt.Sprintf(" \033[90m(%s)\033[0m\n", info.Category))

	if info.Recoverable {
		b.WriteString("\033[38;5;208m⚠️  Este erro é recuperável\033[0m\n")
	}

	if len(info.Suggestions) > 0 {
		b.WriteString("\n\033[38;5;99m💡 Sugestões:\033[0m\n")
		for _, s := range info.Suggestions {
			b.WriteString(fmt.Sprintf("  • %s\n", s))
		}
	}

	return b.String()
}
