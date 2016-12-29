package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Configuration map[string]string

var configuration = make(Configuration)

func InitializeConfig() {
	os.Mkdir(os.Getenv("HOME")+"/.lifx", 0700)
	Set("selector", "all")
}

func Get() (Configuration, error) {
	config_file, err := ioutil.ReadFile(os.Getenv("HOME") + "/.lifx/conf.json")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	json.Unmarshal(config_file, &configuration)
	return configuration, nil
}

func Set(k string, v string) Configuration {
	configuration[k] = v

	b, err := json.Marshal(configuration)

	if err != nil {
		panic(err)
	}

	ioutil.WriteFile(os.Getenv("HOME")+"/.lifx/conf.json", b, 0644)
	return configuration
}
