package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type menus struct {
	// Num 				string `json:"num"`
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

func GetmenuThr(url string) []string {
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
	var meshi []string
	url := "http://www.gakushoku.com/univ_mn2.php"
	s := bufio.NewScanner(os.Stdin)
	s.Scan()

	if s.Text() == "Monday" {
		meshi = GetmenuMon(url) //月曜日のmenu
	} else if s.Text() == "Tuesday" {
		meshi = GetmenuTue(url) //火曜日のmenu
	} else if s.Text() == "Wednesday" {
		meshi = GetmenuWen(url) //水曜日のmenu
	} else if s.Text() == "Thursday" {
		meshi = GetmenuThr(url) //木曜日のmenu
	} else if s.Text() == "Friday" {
		meshi = GetmenuFri(url) //金曜日のmenu
	} else {
		fmt.Println("You hava a mistake")
		return
	}

	// m = append(m, mon)

	for _, m := range meshi {
		fmt.Println(m)
		fmt.Println("")
	}

}
