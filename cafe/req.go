package cafe

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var me []menus

func fileGet() {
	file, err := ioutil.ReadFile("json/meshi.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	jsonerr := json.Unmarshal(file, &me)
	if err != nil {
		fmt.Println("Format Error: ", jsonerr)
	}
}

func RequestCafeMon() menus { //[]string {
	fileGet()
	return me[0]
}

func RequestCafeTue() menus { //[]string {
	fileGet()
	return me[1]
}

func RequestCafeWen() menus { //[]string {
	fileGet()
	return me[2]
}

func RequestCafeThu() menus { //[]string {
	fileGet()
	return me[3]
}

func RequestCafeFri() menus { //[]string {
	fileGet()
	return me[4]
}
