package tui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("63"))
	blurredStyle = lipgloss.NewStyle()
)

func (m Model) View() string {
	box := BoxStyle(m.Width - 10)
	title := TitleStyle(m.Width - 10)
	// bordedArea := BorderStyle(m.Width)

	return box.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			title.Render("🔁 Terminal based POSTMAN"),
			labelStyle.Render("URL: ")+styleInput(m.URIinput.View(), m.FocusIndex == 0),
			// labelStyle.Render("Method: ")+m.MethodInput.,
			labelStyle.Render("Method: ")+styleInput(m.MethodInput[m.MethodIndex], m.FocusIndex == 1),
			labelStyle.Render("Header: ")+styleInput(m.HeaderInput.View(), m.FocusIndex == 2),
			labelStyle.Render("Body: ")+
				styleInput(m.BodyInput.View(), m.FocusIndex == 3),
			buttonStyle.Render("[ Send Request ]"),
			"\n"+m.ResponseArea.View(),
			lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("ff")).
				Align(lipgloss.Center).
				PaddingBottom(1).
				Width(m.Width-10).Render("[tab] to switch | [enter] to send | [q] to quit"),
		),
		labelStyle.Render(),
	)
}

func styleInput(content string, focused bool) string {
	if focused {
		return focusedStyle.Render(content)
	}
	return blurredStyle.Render(content)
}
