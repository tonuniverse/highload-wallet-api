/*
highload-wallet-api – API wrapper over high-load TON wallet smart contract

Copyright (C) 2021 Alexander Gapak

This file is part of highload-wallet-api.

highload-wallet-api is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

highload-wallet-api is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with highload-wallet-api.  If not, see <https://www.gnu.org/licenses/>.
*/

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
