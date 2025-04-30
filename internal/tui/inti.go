package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func New() Model {
	url := textinput.New()
	url.Placeholder = "http://api.example.com/user"
	url.Focus()

	method := textinput.New()
	method.Placeholder = "GET"

	header := textinput.New()
	header.Placeholder = "Content-type: Application"

	body := textinput.New()
	body.Placeholder = `{"name":"surya"}`

	return Model{
		URIinput:    url,
		MethodInput: method,
		HeaderInput: header,
		BodyInput:   body,
		FocusIndex:  0,
	}

}
func (m Model) Init() tea.Cmd {
	return textinput.Blink
}
