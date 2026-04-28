package cmd

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/klebersonromero/hb/internal/ui/layout"
	"github.com/spf13/cobra"
)

var layoutTestCmd = &cobra.Command{
	Use:   "layout-test",
	Short: "Demonstra o motor de layout Flexbox no terminal",
	Run: func(cmd *cobra.Command, args []string) {
		// Define the layout
		root := layout.NewNode(layout.Style{
			Direction: layout.Column,
			Padding:   layout.EdgeValues{Top: 1, Bottom: 1, Left: 2, Right: 2},
		})

		header := layout.NewNode(layout.Style{
			Direction: layout.Row,
			Justify:   layout.JustifySpaceBetween,
			FlexGrow:  1,
		})
		
		title := layout.NewNode(layout.Style{FlexGrow: 1})
		version := layout.NewNode(layout.Style{FlexGrow: 1})
		
		header.AddChild(title)
		header.AddChild(version)

		body := layout.NewNode(layout.Style{
			FlexGrow: 5,
			Padding:  layout.EdgeValues{Top: 1, Bottom: 1},
		})

		footer := layout.NewNode(layout.Style{
			FlexGrow: 1,
			Direction: layout.Row,
		})

		root.AddChild(header)
		root.AddChild(body)
		root.AddChild(footer)

		// Calculate for a 80x20 terminal
		root.Calculate(80, 15)

		// Render (Simple ASCII representation)
		fmt.Println(color.CyanString("┌" + strings.Repeat("─", 78) + "┐"))
		
		renderNode(root, "Root")
		
		fmt.Println(color.CyanString("└" + strings.Repeat("─", 78) + "┘"))
	},
}

func renderNode(n *layout.Node, name string) {
	fmt.Printf("%s: [X:%d Y:%d W:%d H:%d]\n", 
		color.YellowString(name), 
		n.Result.X, n.Result.Y, n.Result.Width, n.Result.Height)
	
	for i, child := range n.Children {
		renderNode(child, fmt.Sprintf("  Child %d", i))
	}
}

func init() {
	uiCmd.AddCommand(layoutTestCmd)
}
