package util

import (
	"fmt"
	"net/http"
)

func CallURLGetHeader(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	header := resp.Header.Get("Content-Type")
	if header == "" {
		return "", fmt.Errorf("response has no content type header")
	}
	return header, nil
}
