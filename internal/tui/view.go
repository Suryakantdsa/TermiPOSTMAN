package tui

import "github.com/charmbracelet/lipgloss"

func (m Model) View() string {
	box := BoxStyle(m.Width)
	title := TitleStyle(m.Width)

	return box.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			title.Render("🔁 REST API TESTER"),
			labelStyle.Render("URL: ")+m.URIinput.View(),
			labelStyle.Render("Method: ")+m.MethodInput.View(),
			labelStyle.Render("Header: ")+m.HeaderInput.View(),
			labelStyle.Render("Body: ")+m.BodyInput.View(),
			buttonStyle.Render("[ Send Request ]"),
			"\n"+m.ResponseArea.View(),
			labelStyle.Render("")+"[tab] to switch | [enter] to send | [q] to quit",
		),
		labelStyle.Render(),
	)
}
