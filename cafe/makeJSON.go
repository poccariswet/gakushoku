package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type menus struct {
	Week       string `json:"week"`
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

func GetmenuMon(url string) []string {
	doc, _ := goquery.NewDocument(url)
	var m []string
	doc.Find("div > div > section > table > tbody > tr").Each(func(_ int, s *goquery.Selection) {
		t := s.Find("td").First().Text()
		m = append(m, t)
	})

	return m
}

func GetmenuTue(url string) []string {
	doc, _ := goquery.NewDocument(url)
	var m []string
	doc.Find("div > div > section > table > tbody > tr").Each(func(_ int, s *goquery.Selection) {
		t := s.Find("td").First().Next().Text()
		m = append(m, t)
	})

	return m
}

func GetmenuWen(url string) []string {
	doc, _ := goquery.NewDocument(url)
	var m []string
	doc.Find("div > div > section > table > tbody > tr").Each(func(_ int, s *goquery.Selection) {
		t := s.Find("td").First().Next().Next().Text()
		m = append(m, t)
	})

	return m
}

func GetmenuThu(url string) []string {
	doc, _ := goquery.NewDocument(url)
	var m []string
	doc.Find("div > div > section > table > tbody > tr").Each(func(_ int, s *goquery.Selection) {
		t := s.Find("td").First().Next().Next().Next().Text()
		m = append(m, t)
	})

	return m
}

func GetmenuFri(url string) []string {
	doc, _ := goquery.NewDocument(url)
	var m []string
	doc.Find("div > div > section > table > tbody > tr").Each(func(_ int, s *goquery.Selection) {
		t := s.Find("td").First().Next().Next().Next().Next().Text()
		m = append(m, t)
	})

	return m
}

func main() {
	var menu []menus //menuのjson型構造体
	url := "http://www.gakushoku.com/univ_mn2.php"

	mon := GetmenuMon(url)
	tue := GetmenuTue(url)
	wen := GetmenuWen(url)
	thr := GetmenuThu(url)
	fri := GetmenuFri(url)

	menu = append(menu, menus{
		Week:       "Monday",
		Higawari:   mon[1],
		Donmono:    mon[2],
		Fish:       mon[3],
		Salad:      mon[4],
		Dessert:    mon[5],
		Onepin:     mon[6],
		Pasta:      mon[7],
		Menrui:     mon[8],
		EveningSet: mon[9],
	})

	menu = append(menu, menus{
		Week:       "Tuesday",
		Higawari:   tue[1],
		Donmono:    tue[2],
		Fish:       tue[3],
		Salad:      tue[4],
		Dessert:    tue[5],
		Onepin:     tue[6],
		Pasta:      tue[7],
		Menrui:     tue[8],
		EveningSet: tue[9],
	})

	menu = append(menu, menus{
		Week:       "Wednesday",
		Higawari:   wen[1],
		Donmono:    wen[2],
		Fish:       wen[3],
		Salad:      wen[4],
		Dessert:    wen[5],
		Onepin:     wen[6],
		Pasta:      wen[7],
		Menrui:     wen[8],
		EveningSet: wen[9],
	})

	menu = append(menu, menus{
		Week:       "Thursday",
		Higawari:   thr[1],
		Donmono:    thr[2],
		Fish:       thr[3],
		Salad:      thr[4],
		Dessert:    thr[5],
		Onepin:     thr[6],
		Pasta:      thr[7],
		Menrui:     thr[8],
		EveningSet: thr[9],
	})

	menu = append(menu, menus{
		Week:       "Friday",
		Higawari:   fri[1],
		Donmono:    fri[2],
		Fish:       fri[3],
		Salad:      fri[4],
		Dessert:    fri[5],
		Onepin:     fri[6],
		Pasta:      fri[7],
		Menrui:     fri[8],
		EveningSet: fri[9],
	})

	bytes, _ := json.Marshal(menu)
	ioutil.WriteFile("./json/meshi.json", bytes, os.ModePerm)
}
