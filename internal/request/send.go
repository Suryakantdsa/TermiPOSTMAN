package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type RequestData struct {
	URL    string
	Method string
	Header map[string]string
	Body   string
}

func SendRequest(data RequestData) (string, error) {
	start := time.Now()
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
	duration := time.Since(start)
	var body bytes.Buffer
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		return "", err
	}

	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, body.Bytes(), "", "  ")
	if err != nil {
		return "", err
	}
	result := fmt.Sprintf(
		"Status: %d | Time: %s\n\n\n%s",
		resp.StatusCode,
		duration.String(),
		prettyJSON.String(),
	)
	// bodyBytes, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return "", err
	// }
	return string(result), nil
}
