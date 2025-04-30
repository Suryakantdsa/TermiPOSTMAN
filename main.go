package main

import (
	"fmt"
	"os"

	"github.com/Suryakantdsa/postman-tui/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(tui.New())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
}
