package cmd

import (
	"log"

	"github.com/Suryakantdsa/postman-tui/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the TUI interface",
	Run: func(cmd *cobra.Command, args []string) {
		model := tui.New()
		if err := tea.NewProgram(model).Start(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
