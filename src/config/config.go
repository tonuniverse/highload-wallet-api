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
