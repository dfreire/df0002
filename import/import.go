package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Exported struct {
	Sections []Section `json:"sections"`
}

type Section struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Wines       []Wine `json:"wines"`
}

type Wine struct {
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Image       string      `json:"image"`
	Region      string      `json:"region"`
	TechSheets  []TechSheet `json:"techSheets"`
}

type TechSheet struct {
	Title        string `json:"title"`
	Vinification string `json:"vinification"`
}

func main() {
	dat, err := ioutil.ReadFile("../exported/iOS_en.json")
	check(err)

	exported := Exported{}
	err = json.Unmarshal(dat, &exported)
	check(err)

	fmt.Println(exported)

	m, err := yaml.Marshal(&exported)
	check(err)
	fmt.Printf(string(m))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
