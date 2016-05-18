package main

import (
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
	fmt.Print(string(dat))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
