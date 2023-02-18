package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type RawConfig struct {
	DNSAddress            string `json:"dns_address"`
	DNSPort               uint16 `json:"dns_port"`
	StatisticsChannelPort uint16 `json:"statistics_port"`
	ListenAddress         string `json:"listen_address"`
	TsigName              string `json:"tsig_key_name"`
	TsigSecret            string `json:"tsig_key_secret"`
}

type Tsig struct {
	Map  map[string]string
	Name string
}

type Config struct {
	StatisticsURL string
	DNSServer     string
	ListenAddress string
	TsigKey       Tsig
}

func LoadConfiguration() Config {
	var rawConfig RawConfig
	configFile, err := os.Open("config.json")
	if err != nil {
		panic("Could not open config.json")
	}
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&rawConfig)
	config := Config{
		StatisticsURL: fmt.Sprintf("%s:%d", rawConfig.DNSAddress, rawConfig.StatisticsChannelPort),
		DNSServer:     fmt.Sprintf("%s:%d", rawConfig.DNSAddress, rawConfig.DNSPort),
		ListenAddress: rawConfig.ListenAddress,
		TsigKey:       Tsig{Map: map[string]string{rawConfig.TsigName: rawConfig.TsigSecret}, Name: rawConfig.TsigName},
	}
	return config
}
