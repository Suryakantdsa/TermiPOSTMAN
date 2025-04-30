package main

import (
	"log"

	"github.com/Suryakantdsa/postman-tui/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	model := tui.New()
	if err := tea.NewProgram(model).Start(); err != nil {
		log.Fatal(err)
	}
}
