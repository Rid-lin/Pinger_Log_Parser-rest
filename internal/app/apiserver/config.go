package apiserver

import (
	"encoding/json"
	"os"
)

//Config ...
type Config struct {
	BindAddr string `JSON:"bind_addr"`
}

//NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
	}
}

//GetConf ...
// Read the config file from the current directory and marshal
// into the conf config struct.
// Example 	servers = getConf("./config.json")
func (conf *Config) GetConf(nameFile string) error {

	configFile, err := os.Open(nameFile)
	if err != nil {
		return err
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(conf)
	return nil
}
