package tui

import (
	"github.com/Suryakantdsa/postman-tui/internal/request"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch message := message.(type) {
	case tea.KeyMsg:
		switch message.String() {
		case "tab":
			m.FocusIndex = (m.FocusIndex + 1) % 4
		case "shift+tab":
			m.FocusIndex = (m.FocusIndex - 1 + 4) % 4
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			return m, func() tea.Msg {
				body, err := request.SendRequest(request.RequestData{
					URL:    m.URIinput.Value(),
					Method: m.MethodInput.Value(),
					Header: map[string]string{
						"Content-type": "application/json",
					},
					Body: m.BodyInput.Value(),
				})

				return ResponseMsg{Body: body, Err: err}
			}

		}
	case ResponseMsg:
		if message.Err != nil {
			m.ResponseArea.SetValue("error: " + message.Err.Error())
		} else {
			m.ResponseArea.SetValue(message.Body)
		}
		return m, nil
	case tea.WindowSizeMsg:
		m.Width = message.Width
		m.Height = message.Height
		return m, nil
	}
	inputs := []*textinput.Model{
		&m.URIinput,
		&m.MethodInput,
		&m.HeaderInput,
		&m.BodyInput,
	}

	for i, input := range inputs {
		if i == m.FocusIndex {
			input.Focus()
		} else {
			input.Blur()
		}

		updatedModel, cmd := input.Update(message)
		*input = updatedModel // ← THIS LINE IS CRUCIAL
		cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}
