package tui

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
)

type Model struct {
	URIinput textinput.Model
	// MethodInput  textinput.Model
	MethodInput  []string
	HeaderInput  textinput.Model
	BodyInput    textarea.Model
	FocusIndex   int
	MethodIndex  int
	ResponseArea textarea.Model
	Width        int
	Height       int
}

type ResponseMsg struct {
	Body string
	Err  error
}
