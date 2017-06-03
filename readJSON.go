package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type menus struct {
	Num        string `json:"num"`
	Higawari   string `json:"higawari"`
	Donmono    string `json:"donmono"`
	Fish       string `json:"fish"`
	Salad      string `json:"salada"`
	Dessert    string `json:"desert"`
	Onepin     string `json:"onepin"`
	Pasta      string `json:"pasta"`
	Menrui     string `json:"menrui"`
	EveningSet string `json:"eveningset"`
}

func main() {
	file, err := ioutil.ReadFile("json/meshi.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var me []menus
	jsonerr := json.Unmarshal(file, &me)
	if err != nil {
		fmt.Println("Format Error: ", jsonerr)
	}

	fmt.Println(me[1].Fish)
}
