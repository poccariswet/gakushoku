package cafe

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type menus struct {
	Week       string `json:"week"`
	Higawari   string `json:"higawari"`
	Donmono    string `json:"donmono"`
	Fish       string `json:"fish"`
	Salada     string `json:"salada"`
	Dessert    string `json:"desert"`
	Onepin     string `json:"onepin"`
	Pasta      string `json:"pasta"`
	Menrui     string `json:"menrui"`
	EveningSet string `json:"eveningset"`
}

const url = "http://www.gakushoku.com/univ_mn2.php"

var menu []menus

func getCafe() {
	weeks := []string{"Monday\n", "Tuesday\n", "Wednesday\n", "Thursday\n", "Friday\n"}
	doc, _ := goquery.NewDocument(url)
	var meshi []string
	doc.Find("td").Each(func(_ int, s *goquery.Selection) {
		data := s.First().Text()
		meshi = append(meshi, data+"\n")
	})
	for i := 0; i < 5; i++ {
		menu = append(menu, menus{
			Week:       weeks[i],
			Higawari:   meshi[i],
			Donmono:    meshi[i+5],
			Fish:       meshi[i+10],
			Salad:      meshi[i+15],
			Dessert:    meshi[i+20],
			Onepin:     meshi[i+25],
			Pasta:      meshi[i+30],
			Menrui:     meshi[i+35],
			EveningSet: meshi[i+40],
		})
	}
}

func UpdateCafe() {
	getCafe()
	bytes, _ := json.Marshal(menu)
	ioutil.WriteFile("json/meshi.json", bytes, os.ModePerm)
	fmt.Println("success")
}
