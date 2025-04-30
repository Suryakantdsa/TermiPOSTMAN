package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	box := BoxStyle(m.Width - 10)
	title := TitleStyle(m.Width - 10)
	// bordedArea := BorderStyle(m.Width)

	return box.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			title.Render("🔁 Terminal based POSTMAN"),
			labelStyle.Render("URL: ")+m.URIinput.View(),
			labelStyle.Render("Method: ")+m.MethodInput.View(),
			labelStyle.Render("Header: ")+m.HeaderInput.View(),
			labelStyle.Render("Body: ")+m.BodyInput.View(),
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
