package infraestucture

import (
	"encoding/json"
	"os"
)



type Config struct {
	Host   		string `json: "host"`
	User    	string `json: "user"`
	Password    string `json: "password"`
	DbName  	string `json:"dbname"`
	Sslmode 	string `json: "sslmode"`
	Port    	string  `json:port`
}

func LoadConfig(configPath string) (*Config, error) {
	file, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config

	// Decodifica o JSON no struct Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil

}
