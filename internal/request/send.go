package request

import (
	"bytes"
	"io"
	"net/http"
)

type RequestData struct {
	URL    string
	Method string
	Header map[string]string
	Body   string
}

func SendRequest(data RequestData) (string, error) {

	client := &http.Client{}
	req, err := http.NewRequest(data.Method, data.URL, bytes.NewBufferString(data.Body))
	if err != nil {
		return "", err
	}
	for k, v := range data.Header {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bodyBytes), nil
}
