package client

/*

TODO

factory
mark selected in List()

*/

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lifx/config"
	"net/http"
	"os"
)

const base_path = "https://api.lifx.com/v1"

type Product struct {
	Name            string `json:"name"`
	Identifier      string `json:"identifier"`
	LastSeen        string `json:"last_seen"`
	SecondSinceSeen string `json:'second_since_seen"`
}

type Light struct {
	Id      string  `json:"id"`
	Product Product `json:"product"`
}

type Lights []Light

var configuration map[string]string

func init() {
	configuration = config.Get()

	if configuration["token"] == "" {
		fmt.Println("Set your token in ~/.lifx/conf.json")
		os.Exit(1)
	}
}

func Verbose(debug bool, body []byte, err error) {
	if err != nil {
		panic(err)
	}
	if debug {
		fmt.Println(string(body))
	}
}

func Config() {
	fmt.Printf("Configuration:\n\n")
	for k, v := range configuration {
		fmt.Printf("%-16s\t%s\n", k, v)
	}
	fmt.Printf("\n")
}

func SelectLight(debug bool, selector string) {
	if selector != "" {
		configuration = config.Set("selector", selector)
	}
	fmt.Printf("Selected:\t%s\n", configuration["selector"])
}

func List(debug bool) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", base_path+"/lights/all", nil)
	req.Header.Add("Authorization", "Bearer "+configuration["token"])
	res, err := client.Do(req)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var lights Lights

	json.Unmarshal(body, &lights)

	for i := 0; i < len(lights); i++ {
		fmt.Printf("Id:\t%s\n", lights[i].Id)
		fmt.Printf("Name:\t%s\n", lights[i].Product.Name)
		fmt.Printf("\n")
	}
}

func Toggle(debug bool) {
	client := &http.Client{}
	s := configuration["selector"]
	req, err := http.NewRequest("POST", base_path+"/lights/"+s+"/toggle", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "Bearer "+configuration["token"])
	res, err := client.Do(req)
	body, err := ioutil.ReadAll(res.Body)
	Verbose(debug, body, err)
}

func On(debug bool) {
	client := &http.Client{}
	s := configuration["selector"]
	payload := bytes.NewBuffer([]byte(`{"power":"on"}`))
	req, err := http.NewRequest("PUT", base_path+"/lights/"+s+"/state", payload)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "Bearer "+configuration["token"])
	res, err := client.Do(req)
	body, err := ioutil.ReadAll(res.Body)
	Verbose(debug, body, err)
}

func Off(debug bool) {
	client := &http.Client{}
	s := configuration["selector"]
	payload := bytes.NewBuffer([]byte(`{"power":"off"}`))
	req, err := http.NewRequest("PUT", base_path+"/lights/"+s+"/state", payload)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "Bearer "+configuration["token"])
	res, err := client.Do(req)
	body, err := ioutil.ReadAll(res.Body)
	Verbose(debug, body, err)
}

func Brightness(debug bool, b string) {
	client := &http.Client{}
	s := configuration["selector"]
	payload := bytes.NewBuffer([]byte(`{"brightness":` + b + `}`))
	req, err := http.NewRequest("PUT", base_path+"/lights/"+s+"/state", payload)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "Bearer "+configuration["token"])
	res, err := client.Do(req)
	body, err := ioutil.ReadAll(res.Body)
	Verbose(debug, body, err)
}

func Color(debug bool, c string) {
	client := &http.Client{}
	s := configuration["selector"]
	payload := bytes.NewBuffer([]byte(`{"color":"` + c + `"}`))
	req, err := http.NewRequest("PUT", base_path+"/lights/"+s+"/state", payload)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "Bearer "+configuration["token"])
	res, err := client.Do(req)
	body, err := ioutil.ReadAll(res.Body)
	Verbose(debug, body, err)
}
