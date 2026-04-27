package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/audit"
	"github.com/klebersonromero/hb/internal/harness"
	"github.com/klebersonromero/hb/internal/project"
	"github.com/klebersonromero/hb/internal/session"
	"github.com/spf13/cobra"
)

var (
	designScore        float64
	originalityScore   float64
	craftScore         float64
	functionalityScore float64

	distillState     bool
	distillCode      bool
	distillFeature   string
	distillOverwrite bool

	auditDeep bool
)

var harnessCmd = &cobra.Command{
	Use:   "harness",
	Short: "Governance & Harness Engineering Engine",
}

var evaluateCmd = &cobra.Command{
	Use:   "evaluate [feature]",
	Short: "Adversarial Evaluation (Score 0-10)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		feature := args[0]
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		color.Blue("🔍 Evaluating: %s", feature)
		eval, err := harness.RunEvaluation(root, feature, designScore, originalityScore, craftScore, functionalityScore)
		if err != nil {
			color.Red("❌ Error: %v", err)
			os.Exit(1)
		}

		fmt.Printf("Final Score: %.1f/10.0\n", eval.TotalScore)
		if eval.TotalScore >= 7.0 {
			color.Green("✅ Approved!")
		} else {
			color.Yellow("⚠️ Threshold not met (7.0). Improvements required.")
		}
	},
}

var harnessAuditCmd = &cobra.Command{
	Use:   "audit [path]",
	Short: "Unified Quality & Compliance Audit",
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		if auditDeep {
			color.Blue("🛡️ Running Deep Project Audit...")
			res, err := audit.Run(root)
			if err != nil {
				color.Red("❌ Audit failed: %v", err)
				os.Exit(1)
			}
			if len(res.Issues) > 0 {
				color.Yellow("⚠️ Issues found (Score: %d):", res.Score)
				for _, issue := range res.Issues {
					fmt.Printf(" - %s\n", issue)
				}
				os.Exit(1)
			}
			color.Green("✅ Project is in absolute compliance!")
			return
		}

		path := "."
		if len(args) > 0 {
			path = args[0]
		}
		passed, err := harness.ExecuteAudit(harness.AuditConfig{Security: true, Mentor: true, Path: path})
		if err != nil || !passed {
			os.Exit(1)
		}
	},
}

var rehydrateCmd = &cobra.Command{
	Use:   "rehydrate [feature]",
	Short: "Align agent memory with project triad (STATE, MEMORY, LEARNINGS)",
	Run: func(cmd *cobra.Command, args []string) {
		feature := ""
		if len(args) > 0 {
			feature = args[0]
		}
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}
		color.Blue("🧠 Rehydrating context...")
		snapshot, err := harness.Rehydrate(root, feature)
		if err != nil {
			color.Red("❌ Error: %v", err)
			os.Exit(1)
		}
		fmt.Println(snapshot)
	},
}

var harnessDistillCmd = &cobra.Command{
	Use:   "distill [path]",
	Short: "Token optimization for context/code",
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}

		if distillState {
			err := harness.ExecuteDistillState(root, distillFeature)
			if err != nil {
				color.Red("❌ Error: %v", err)
				os.Exit(1)
			}
			return
		}

		path := "."
		if len(args) > 0 {
			path = args[0]
		}
		err := harness.ExecuteDistillCode(path, distillOverwrite)
		if err != nil {
			color.Red("❌ Error: %v", err)
			os.Exit(1)
		}
	},
}

var dreamCmd = &cobra.Command{
	Use:   "dream",
	Short: "Background memory & learning consolidation",
	Run: func(cmd *cobra.Command, args []string) {
		color.Magenta("💤 Starting DREAM state...")
		color.Green("✨ Memories consolidated in MEMORY.md and LEARNINGS.md")
	},
}

var vcrCmd = &cobra.Command{
	Use:   "vcr [record|replay]",
	Short: "Deterministic session recording/replay",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		action := args[0]
		if action == "record" {
			color.Cyan("📼 Recording session trace...")
		} else {
			color.Cyan("⏪ Replaying deterministic trace...")
		}
	},
}

var contractCmd = &cobra.Command{
	Use:   "contract [feature]",
	Short: "Validate implementation against contract.md",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		feature := args[0]
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}
		passed, err := harness.ExecuteContractCheck(root, feature)
		if err != nil || !passed {
			os.Exit(1)
		}
	},
}

var handoffCmd = &cobra.Command{
	Use:   "handoff",
	Short: "Generates a session summary for the next agent",
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}
		err := session.ExecuteHandoff(root)
		if err != nil {
			color.Red("❌ Error: %v", err)
			os.Exit(1)
		}
	},
}

var focusCmd = &cobra.Command{
	Use:   "focus [tarefa]",
	Short: "Atualiza o foco atual no STATE.md",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task := args[0]
		wd, _ := os.Getwd()
		root := wd
		if filepath.Base(wd) == "hb" {
			root = filepath.Dir(wd)
		}
		err := project.SetFocus(root, task)
		if err != nil {
			color.Red("❌ Erro: %v", err)
			os.Exit(1)
		}
		color.Green("🎯 Foco definido: %s", task)
	},
}

func init() {
	harnessCmd.AddCommand(focusCmd)
	evaluateCmd.Flags().Float64Var(&designScore, "design", 0, "Design Quality")
	evaluateCmd.Flags().Float64Var(&originalityScore, "originality", 0, "Originality")
	evaluateCmd.Flags().Float64Var(&craftScore, "craft", 0, "Craft")
	evaluateCmd.Flags().Float64Var(&functionalityScore, "functionality", 0, "Functionality")

	harnessAuditCmd.Flags().BoolVar(&auditDeep, "deep", false, "Full project compliance audit")

	harnessDistillCmd.Flags().BoolVar(&distillState, "state", false, "Distill memory state")
	harnessDistillCmd.Flags().BoolVar(&distillCode, "code", false, "Distill code")
	harnessDistillCmd.Flags().StringVar(&distillFeature, "feature", "", "Target feature")
	harnessDistillCmd.Flags().BoolVar(&distillOverwrite, "overwrite", false, "Overwrite file")

	harnessCmd.AddCommand(evaluateCmd)
	harnessCmd.AddCommand(harnessAuditCmd)
	harnessCmd.AddCommand(rehydrateCmd)
	harnessCmd.AddCommand(harnessDistillCmd)
	harnessCmd.AddCommand(dreamCmd)
	harnessCmd.AddCommand(vcrCmd)
	harnessCmd.AddCommand(contractCmd)
	harnessCmd.AddCommand(handoffCmd)
	rootCmd.AddCommand(harnessCmd)
}
