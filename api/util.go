package api

import "net/http"
import "io"
import "log"

var LichessUrl string = "https://lichess.org"

func checkErr(err error) {
	if (err != nil) {
		log.Fatal(err)
	}
}

func NewRequest(token, method, url string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(method, url, body)
	request.Header.Set("Authorization", "Bearer " + token)
	return request, err
}

func NewClient() (*http.Client) {
	var transport *http.Transport = &http.Transport {
		IdleConnTimeout: 0,
	}
	var client *http.Client = &http.Client{Transport : transport}
	return client
}
