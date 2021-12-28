package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
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
}

var Cfg Config

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
