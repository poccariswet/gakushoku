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

func RequestCafeMon() menus { //月曜日
	fileGet()
	return me[0]
}

func RequestCafeTue() menus { //火曜日
	fileGet()
	return me[1]
}

func RequestCafeWen() menus { //水曜日
	fileGet()
	return me[2]
}

func RequestCafeThu() menus { //木曜日
	fileGet()
	return me[3]
}

func RequestCafeFri() menus { //金曜日
	fileGet()
	return me[4]
}
