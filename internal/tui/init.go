package tui

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func New() *Model {
	uri := textinput.New()
	uri.Placeholder = "https://api.example.com"
	uri.Focus()
	uri.CharLimit = 200
	uri.Width = 100

	method := textinput.New()
	method.Placeholder = "GET"
	method.Width = 40

	headers := textinput.New()
	headers.Placeholder = "Content-Type: application/json"
	headers.Width = 40

	body := textinput.New()
	body.Placeholder = `{"name": "surya"}`
	body.Width = 40

	responseArea := textarea.New()
	responseArea.Placeholder = "Response will appear here.."
	responseArea.SetHeight(10)
	responseArea.SetWidth(100)
	responseArea.ShowLineNumbers = true
	responseArea.FocusedStyle.CursorLine = responseArea.FocusedStyle.CursorLine.Bold(true)

	return &Model{
		URIinput:     uri,
		MethodInput:  method,
		HeaderInput:  headers,
		BodyInput:    body,
		ResponseArea: responseArea,
		FocusIndex:   0,
	}

}
func (m *Model) Init() tea.Cmd {
	return textinput.Blink
}
