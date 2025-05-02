// package main

// import (
// 	"log"

// 	"github.com/Suryakantdsa/postman-tui/internal/tui"
// 	tea "github.com/charmbracelet/bubbletea"
// )

//	func main() {
//		model := tui.New()
//		if err := tea.NewProgram(model).Start(); err != nil {
//			log.Fatal(err)
//		}
//	}
package main

import "github.com/Suryakantdsa/postman-tui/cmd"

func main() {
	cmd.Execute()
}
