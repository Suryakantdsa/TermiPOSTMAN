package tui

import "github.com/charmbracelet/lipgloss"

// var (
// 	boxStyle = lipgloss.NewStyle().
// 			Border(lipgloss.RoundedBorder()).
// 			BorderForeground(lipgloss.Color("63")).
// 			Padding(1, 2).
// 			Width(60)

// 	titleStyle = lipgloss.NewStyle().
// 			Bold(true).
// 			Foreground(lipgloss.Color("63")).
// 			Align(lipgloss.Center).
// 			Width(60)

// 	labelStyle = lipgloss.NewStyle().
// 			Bold(true).
// 			Width(12)

// 	buttonStyle = lipgloss.NewStyle().
// 			Padding(0, 1).
// 			Foreground(lipgloss.Color("15")).
// 			Background(lipgloss.Color("63")).
// 			Bold(true).
// 			MarginTop(1).
// 			Align(lipgloss.Center)
// )

var (
	labelStyle = lipgloss.NewStyle().
			Bold(true).
			Width(12)

	buttonStyle = lipgloss.NewStyle().
			Padding(0, 1).
			Foreground(lipgloss.Color("15")).
			Background(lipgloss.Color("63")).
			Bold(true).
			MarginTop(1).
			Align(lipgloss.Center)
)

func BoxStyle(width int) lipgloss.Style {
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("63")).
		Padding(1, 2).
		Width(width - 4)
}

func TitleStyle(width int) lipgloss.Style {
	return lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("63")).
		Align(lipgloss.Center).
		PaddingBottom(1).
		Width(width - 4)
}
