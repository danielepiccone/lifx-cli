package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var configuration Configuration

type Configuration map[string]string

func init() {
	configuration = Get()
	if configuration["selector"] == "" {
		Set("selector", "all")
	}
}

func Get() Configuration {
	config_file, err := ioutil.ReadFile(os.Getenv("HOME") + "/.lifx/conf.json")
	if err != nil {
		panic(err)
	}
	var configuration Configuration
	json.Unmarshal(config_file, &configuration)
	return configuration
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
