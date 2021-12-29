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
