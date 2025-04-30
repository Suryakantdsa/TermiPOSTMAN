package tui

import "github.com/charmbracelet/lipgloss"

func (m Model) View() string {
	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("285"))
	labelStyle := lipgloss.NewStyle().Bold(true).MarginRight(3)

	return lipgloss.JoinVertical(lipgloss.Left,
		titleStyle.Render("🔁 REST API TESTER"),
		labelStyle.Render("URL: ")+m.URIinput.View(),
		labelStyle.Render("Method: ")+m.MethodInput.View(),
		labelStyle.Render("Header: ")+m.HeaderInput.View(),
		labelStyle.Render("Body: ")+m.BodyInput.View(),
		"\n[tab] to switch | [enter] to send | [q] to quit",
	)
}
