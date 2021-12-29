package mhttp

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"time"
)

func JsonSendPost(url string, data []byte) ([]byte, error) {
	var body []byte = nil

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return body, err
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Timeout: 2 * time.Second, Transport: tr}

	response, err := client.Do(request)
	if err != nil {
		return body, err
	}

	defer func() {
		response.Body.Close()
	}()

	body, _ = ioutil.ReadAll(response.Body)
	if err != nil {
		return body, err
	}

	return body, nil
}
