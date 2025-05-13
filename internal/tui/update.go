package tui

import (
	"github.com/Suryakantdsa/postman-tui/internal/request"
	tea "github.com/charmbracelet/bubbletea"
)

// func (m *Model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
// 	var cmds []tea.Cmd

// 	switch message := message.(type) {
// 	case tea.KeyMsg:
// 		switch message.String() {
// 		case "tab":
// 			m.FocusIndex = (m.FocusIndex + 1) % 4
// 		case "shift+tab":
// 			m.FocusIndex = (m.FocusIndex - 1 + 4) % 4
// 		case "ctrl+c", "q":
// 			return m, tea.Quit
// 		case "enter":
// 			return m, func() tea.Msg {
// 				body, err := request.SendRequest(request.RequestData{
// 					URL:    m.URIinput.Value(),
// 					Method: m.MethodInput.Value(),
// 					Header: map[string]string{
// 						"Content-type": "application/json",
// 					},
// 					Body: m.BodyInput.Value(),
// 				})

// 				return ResponseMsg{Body: body, Err: err}
// 			}
// 		case "up", "down", "pgup", "pgdown":
// 			var cmd tea.Cmd
// 			m.ResponseArea, cmd = m.ResponseArea.Update(message)
// 			return m, cmd
// 		}
// 	case ResponseMsg:
// 		if message.Err != nil {
// 			m.ResponseArea.SetValue("error: " + message.Err.Error())
// 		} else {
// 			m.ResponseArea.SetValue(message.Body)

// 		}
// 		m.ResponseArea.CursorDown()
// 		return m, nil
// 	case tea.WindowSizeMsg:
// 		m.Width = message.Width
// 		m.Height = message.Height
// 		m.ResponseArea.SetWidth(m.Width - 10)
// 		return m, nil
// 	}

// 	inputs := []*textinput.Model{
// 		&m.URIinput,
// 		&m.MethodInput,
// 		&m.HeaderInput,
// 		&m.BodyInput,
// 	}

// 	for i, input := range inputs {
// 		if i == m.FocusIndex {
// 			input.Focus()
// 		} else {
// 			input.Blur()
// 		}

// 		updatedModel, cmd := input.Update(message)
// 		*input = updatedModel
// 		cmds = append(cmds, cmd)
// 	}
// 	return m, tea.Batch(cmds...)
// }

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch message := msg.(type) {
	case tea.KeyMsg:
		switch message.String() {
		case "tab":
			m.FocusIndex = (m.FocusIndex + 1) % 5
		case "shift+tab":
			m.FocusIndex = (m.FocusIndex - 1 + 4) % 4
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			if m.FocusIndex == 1 {
				m.MethodIndex = (m.MethodIndex - 1 + len(m.MethodInput)) % len(m.MethodInput)
			} else if m.FocusIndex == 3 {
				m.BodyInput.CursorUp()
			}
		case "down":
			if m.FocusIndex == 1 {
				m.MethodIndex = (m.MethodIndex + 1) % len(m.MethodInput)
			} else if m.FocusIndex == 3 {
				m.BodyInput.CursorDown()
			}
		case "enter":
			if m.FocusIndex == 3 && m.MethodInput[m.MethodIndex] == "POST" {
				var cmd tea.Cmd
				m.BodyInput, cmd = m.BodyInput.Update(message)
				return m, cmd
			}
			return m, func() tea.Msg {
				body, err := request.SendRequest(request.RequestData{
					URL:    m.URIinput.Value(),
					Method: m.MethodInput[m.MethodIndex],
					Header: map[string]string{
						"Content-type": "application/json",
					},
					Body: m.BodyInput.Value(),
				})
				return ResponseMsg{Body: body, Err: err}
			}
		case "ctrl+enter":
			return m, func() tea.Msg {
				body, err := request.SendRequest(request.RequestData{
					URL:    m.URIinput.Value(),
					Method: m.MethodInput[m.MethodIndex],
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
		m.ResponseArea.CursorDown()
		return m, nil
	case tea.WindowSizeMsg:
		m.Width = message.Width
		m.Height = message.Height
		m.ResponseArea.SetWidth(m.Width - 10)
		return m, nil
	}

	// Handle focus and update
	if m.FocusIndex == 0 {
		var cmd tea.Cmd
		m.URIinput, cmd = m.URIinput.Update(msg)
		cmds = append(cmds, cmd)
	} else if m.FocusIndex == 2 {
		var cmd tea.Cmd
		m.HeaderInput, cmd = m.HeaderInput.Update(msg)
		cmds = append(cmds, cmd)
	} else if m.FocusIndex == 3 {
		if m.MethodInput[m.MethodIndex] == "POST" {
			m.BodyInput.Focus()
			var cmd tea.Cmd
			m.BodyInput, cmd = m.BodyInput.Update(msg)
			cmds = append(cmds, cmd)
		} else {
			m.BodyInput.Blur()
		}
	}

	return m, tea.Batch(cmds...)
}
