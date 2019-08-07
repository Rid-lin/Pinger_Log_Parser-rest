package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func checkExistFile(fileName string) (bool, error) {
	if _, err := os.Stat(fileName); err == nil {
		// path/to/whatever exists
		return true, nil
	} else if os.IsNotExist(err) {
		// path/to/whatever does *not* exist
		return false, err
	} else {
		// Schrodinger: file may or may not exist. See err for details.
		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
		return false, err
	}
}

func getExistingFile(fileName string, names []string) string {
	if len(names) > 0 {
		for _, name := range names {
			if _, err := os.Stat(fileName); err == nil {
				// path/to/whatever exists
				return name + fileName
			} else if os.IsNotExist(err) {
				// path/to/whatever does *not* exist
				continue
			} else {
				// Schrodinger: file may or may not exist. See err for details.
				// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
				continue
			}
		}
	}
	return "./" + fileName
}

// Read the config file from the current directory and marshal
// into the conf config struct.
// Example 	servers = getConf("./config.json")
func getConf(nameFile string) Configuration {

	configFile, err := os.Open(nameFile)
	string := fmt.Sprint("Get of config failed. Config file not found on patch ", nameFile)
	chkM(string, err)
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	// conf := Configuration{}
	var conf *Configuration
	jsonParser.Decode(&conf)
	return *conf
}

// Example saveConf("./config.json", struct)
func saveConf(nameFile string, conf *Configuration) {

	j, err := json.MarshalIndent(conf, "", "    ")
	chk(err)
	fileConfig, err := os.OpenFile(nameFile, os.O_WRONLY|os.O_CREATE, 0644)
	stringE := fmt.Sprint("Save of config failed. Config file not found on patch ", nameFile)
	chkM(stringE, err)
	_, err = fileConfig.Write(j)
	chk(err)
	fileConfig.Close()
}

// Example err := backupConfig("./config.json")
func backupConfig(nameFile string) error {
	srcFolder := nameFile
	destFolder := nameFile + ".bak"
	err := os.Rename(srcFolder, destFolder)
	return err
}
