package api

import "net/http"
import "io"

func NewRequest(token, method, url string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(method, url, body)
	request.Header.Set("Authorization", "Bearer " + token)
	return request, err
}
