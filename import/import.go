package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type exported struct {
	Sections []section `json:"sections"`
}

type section struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Wines       []wine `json:"wines"`
}

type wine struct {
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Image       string      `json:"image"`
	Region      string      `json:"region"`
	TechSheets  []techSheet `json:"techSheets"`
}

type techSheet struct {
	Title        string `json:"title"`
	Vinification string `json:"vinification"`
}

func main() {
	dat, err := ioutil.ReadFile("../exported/iOS_en.json")
	check(err)

	exp := exported{}
	err = json.Unmarshal(dat, &exp)
	check(err)

	fmt.Println(exp)

	m, err := yaml.Marshal(&exp)
	check(err)
	fmt.Printf(string(m))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
