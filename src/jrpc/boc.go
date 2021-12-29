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

package jrpc

import (
	"encoding/base64"
	"encoding/json"
	"highload-wallet-api/src/mhttp"
	"io/ioutil"
)

func SendBocFromFile(jsonRpcURL string, filename string) (string, error) {

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	var b64boc string = base64.StdEncoding.EncodeToString(bytes)

	var data = struct {
		Jsonrpc string `json:"jsonrpc"`
		Method  string `json:"method"`
		Params  struct {
			Boc string `json:"boc"`
		} `json:"params"`
	}{
		Jsonrpc: "2.0",
		Method:  "sendBoc",
	}

	data.Params.Boc = b64boc

	b, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	resp, err := mhttp.JsonSendPost(jsonRpcURL, b)
	if err != nil {
		return "", err
	}

	return string(resp), nil
}
