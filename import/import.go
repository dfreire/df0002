package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
