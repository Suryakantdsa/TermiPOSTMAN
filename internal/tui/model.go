package tui

import (
	"net/http"

	"github.com/charmbracelet/bubbles/textinput"
)

type Model struct {
	URIinput    textinput.Model
	MethodInput textinput.Model
	HeaderInput textinput.Model
	BodyInput   textinput.Model
	FocusIndex  int
}

type ResponseMsg struct {
	Body *http.Response
	Err  error
}
