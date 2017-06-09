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

func RequestCafeMon() string { //月曜日
	fileGet()
	meshi := fmt.Sprintf("曜日: %s日替わり: %s丼もの: %s魚: %sサラダ: %sデザート: %s一品: %sパスタ: %s麺類: %s夕方セット: %s", me[0].Week, me[0].Higawari, me[0].Donmono, me[0].Fish, me[0].Salada, me[0].Dessert, me[0].Onepin, me[0].Pasta, me[0].Menrui, me[0].EveningSet)
	return meshi
}

func RequestCafeTue() string { //火曜日
	fileGet()
	meshi := fmt.Sprintf("曜日: %s日替わり: %s丼もの: %s魚: %sサラダ: %sデザート: %s一品: %sパスタ: %s麺類: %s夕方セット: %s", me[1].Week, me[1].Higawari, me[1].Donmono, me[1].Fish, me[1].Salada, me[1].Dessert, me[1].Onepin, me[1].Pasta, me[1].Menrui, me[1].EveningSet)
	return meshi
}

func RequestCafeWen() string { //水曜日
	fileGet()
	meshi := fmt.Sprintf("曜日: %s日替わり: %s丼もの: %s魚: %sサラダ: %sデザート: %s一品: %sパスタ: %s麺類: %s夕方セット: %s", me[2].Week, me[2].Higawari, me[2].Donmono, me[2].Fish, me[2].Salada, me[2].Dessert, me[2].Onepin, me[2].Pasta, me[2].Menrui, me[2].EveningSet)
	return meshi
}

func RequestCafeThu() string { //木曜日
	fileGet()
	meshi := fmt.Sprintf("曜日: %s日替わり: %s丼もの: %s魚: %sサラダ: %sデザート: %s一品: %sパスタ: %s麺類: %s夕方セット: %s", me[3].Week, me[3].Higawari, me[3].Donmono, me[3].Fish, me[3].Salada, me[3].Dessert, me[3].Onepin, me[3].Pasta, me[3].Menrui, me[3].EveningSet)
	return meshi
}

func RequestCafeFri() string { //金曜日
	fileGet()
	meshi := fmt.Sprintf("曜日: %s日替わり: %s丼もの: %s魚: %sサラダ: %sデザート: %s一品: %sパスタ: %s麺類: %s夕方セット: %s", me[4].Week, me[4].Higawari, me[4].Donmono, me[4].Fish, me[4].Salada, me[4].Dessert, me[4].Onepin, me[4].Pasta, me[4].Menrui, me[4].EveningSet)
	return meshi
}
