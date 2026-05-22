
package main

import (
	"fmt"
	"strings"
	"time"
)

func printAssistantCard(tree *SessionTree, provider string, nodeID string, content string) {
	lines := strings.Split(content, "\n")
	repeatCount := 50 - len(provider)
	if repeatCount < 0 {
		repeatCount = 0
	}
	fmt.Printf("\n\033[38;5;99m┌── \033[1;32m🤖 ASSISTENTE (%s)\033[0m \033[38;5;99m%s\033[0m\n", strings.ToUpper(provider), strings.Repeat("─", repeatCount))
	
	for _, line := range lines {
		fmt.Printf("\033[38;5;99m│\033[0m  %s\n", line)
	}
	
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	nodeTokens := 0
	if n, ok := tree.Nodes[nodeID]; ok {
		nodeTokens = n.Tokens
	}
	totalTokens := tree.GetTotalTokens()
	cost := tree.EstimateUSD(provider)
	costStr := fmt.Sprintf("$%.4f USD", cost)
	if cost == 0.0 {
		costStr = "Free/Local"
	}
	
	footerInfo := fmt.Sprintf("nó: %s | 🪙  %dt (total: %dt) | 💰 %s | 🕒 %s", nodeID, nodeTokens, totalTokens, costStr, timestamp)
	repeatCountFooter := 75 - len(footerInfo)
	if repeatCountFooter < 0 {
		repeatCountFooter = 0
	}
	fmt.Printf("\033[38;5;99m└── \033[90m(%s)\033[0m \033[38;5;99m%s\033[0m\n", footerInfo, strings.Repeat("─", repeatCountFooter))
}

func formatBoxLine(label string, val string, valColor string) string {
	totalWidth := 70
	visibleLen := len(label) + len(val)
	padding := totalWidth - visibleLen
	if padding < 0 {
		padding = 0
	}
	return fmt.Sprintf("\033[38;5;99m│\033[0m   %s%s%s%s   \033[38;5;99m│\033[0m\n", label, valColor, val, strings.Repeat(" ", padding))
}

func startSpinner(suffix string) chan struct{} {
	stopChan := make(chan struct{})
	go func() {
		frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
		i := 0
		start := time.Now()
		for {
			select {
			case <-stopChan:
				fmt.Print("\r\033[K")
				return
			default:
				elapsed := time.Since(start).Seconds()
				fmt.Printf("\r\033[38;5;99m%s\033[0m \033[1m%s\033[0m \033[90m(%.1fs)\033[0m", frames[i%len(frames)], suffix, elapsed)
				i++
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
	return stopChan
}
