package cmd

import (
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/compressor"
	"github.com/spf13/cobra"
)

var compressCmd = &cobra.Command{
	Use:   "compress [file]",
	Short: "Comprime arquivos Markdown para economizar tokens",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		// Se nenhum arquivo for passado, tenta comprimir o STATE.md (lógica check-state)
		if len(args) == 0 {
			statePath := filepath.Join(root, ".specs", "project", "STATE.md")
			info, err := os.Stat(statePath)
			if err != nil {
				color.Red("❌ STATE.md não encontrado.")
				return
			}

			// Heurística: comprimir se maior que 5KB (~1000 tokens)
			if info.Size() > 5000 {
				color.Blue("📝 STATE.md está grande (%d bytes). Comprimindo...", info.Size())
				err = compressor.CompressFile(statePath)
				if err != nil {
					color.Red("❌ Erro ao comprimir: %v", err)
				}
			} else {
				color.Green("✅ O tamanho do STATE.md está otimizado (%d bytes).", info.Size())
			}
			return
		}

		// Comprimir arquivo específico
		target := args[0]
		err := compressor.CompressFile(target)
		if err != nil {
			color.Red("❌ Erro ao comprimir %s: %v", target, err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(compressCmd)
}
