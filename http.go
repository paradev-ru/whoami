package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpGet(url string) (string, error) {
	var body string

	resp, err := http.Get(url)
	if err != nil {
		return body, err
	}

	if resp.StatusCode != http.StatusOK {
		return body, fmt.Errorf("Error getting url: %s (%s)", url, resp.Status)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}

	err = resp.Body.Close()
	if err != nil {
		return body, err
	}

	body = string(b)
	body = strings.TrimSpace(body)

	return body, nil
}
