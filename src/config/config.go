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

package config

import (
	"encoding/json"
	"log"
	"os"
)

type FileConfig struct {
	Server struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"server"`

	TonNet struct {
		JsonRpcURL string `json:"json_rpc_url"`
	} `json:"ton_net"`

	Fift struct {
		Path   string `json:"path"`
		Binary string `json:"binary"`
	} `json:"fift"`

	Contract struct {
		NewOrderFif string `json:"new_order_fif"`
	} `json:"contract"`

	TempPath struct {
		Orders string `json:"orders"`
		Bocs   string `json:"bocs"`
	} `json:"temp_path"`

	Wallet struct {
		Path        string `json:"path"`
		Name        string `json:"name"`
		SubwalletID string `json:"subwallet_id"`
	} `json:"wallet"`
}

var Cfg FileConfig

func Configure() {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer jsonFile.Close()

	if err := json.NewDecoder(jsonFile).Decode(&Cfg); err != nil {
		log.Fatal(err)
	}

}
