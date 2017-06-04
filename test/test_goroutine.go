package main

import (
	"fmt"
	"time"

	"github.com/PuerkitoBio/goquery"
)

//[]string
func GetmenuMon(url string, c chan []string) {
	doc, _ := goquery.NewDocument(url)
	var m []string
	doc.Find("div > div > section > table > tbody > tr").Each(func(_ int, s *goquery.Selection) {
		t := s.Find("td").First().Text()
		m = append(m, t)
	})
	c <- m
	// return m
}

func GetmenuTue(url string, c chan []string) {
	doc, _ := goquery.NewDocument(url)
	var m []string
	doc.Find("div > div > section > table > tbody > tr").Each(func(_ int, s *goquery.Selection) {
		t := s.Find("td").First().Next().Text()
		m = append(m, t)
	})
	time.Sleep(100 * time.Millisecond) //Mondayのgetを待ってる
	c <- m
	// return m
}

func GetmenuWen(url string, c chan []string) {
	doc, _ := goquery.NewDocument(url)
	var m []string
	doc.Find("div > div > section > table > tbody > tr").Each(func(_ int, s *goquery.Selection) {
		t := s.Find("td").First().Next().Next().Text()
		m = append(m, t)
	})
	time.Sleep(200 * time.Millisecond) //Mondayのgetを待ってる
	c <- m
	// return m
}

func GetmenuThu(url string, c chan []string) {
	doc, _ := goquery.NewDocument(url)
	var m []string
	doc.Find("div > div > section > table > tbody > tr").Each(func(_ int, s *goquery.Selection) {
		t := s.Find("td").First().Next().Next().Next().Text()
		m = append(m, t)
	})
	time.Sleep(300 * time.Millisecond) //Mondayのgetを待ってる
	c <- m
	// return m
}

func GetmenuFri(url string, c chan []string) {
	doc, _ := goquery.NewDocument(url)
	var m []string
	doc.Find("div > div > section > table > tbody > tr").Each(func(_ int, s *goquery.Selection) {
		t := s.Find("td").First().Next().Next().Next().Next().Text()
		m = append(m, t)
	})
	time.Sleep(400 * time.Millisecond) //Mondayのgetを待ってる
	c <- m
	// return m
}

//並列処理でそれぞれの曜日のmenuをとってくる
func main() {
	url := "http://www.gakushoku.com/univ_mn2.php"
	c := make(chan []string)
	defer close(c)
	go GetmenuMon(url, c)
	go GetmenuTue(url, c)
	go GetmenuWen(url, c)
	go GetmenuThu(url, c)
	go GetmenuFri(url, c)
	res1, res2, res3, res4, res5 := <-c, <-c, <-c, <-c, <-c

	// fmt.Println(res4[4])
	for _, m1 := range res1 {
		fmt.Println(m1)
		// time.Sleep(2 * time.Second)
	}

	for _, m2 := range res2 {
		fmt.Println(m2)
	}

	for _, m3 := range res3 {
		fmt.Println(m3)
	}

	for _, m4 := range res4 {
		fmt.Println(m4)
	}

	for _, m5 := range res5 {
		fmt.Println(m5)
	}
}