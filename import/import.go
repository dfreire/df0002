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
	Title string `json:"title"`
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
