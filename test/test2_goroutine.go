package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

//[]string
func GetmenuMon(url string, c1 chan []string) {
	doc, _ := goquery.NewDocument(url)
	var m []string
	doc.Find("div > div > section > table > tbody > tr").Each(func(_ int, s *goquery.Selection) {
		t := s.Find("td").First().Text()
		m = append(m, t)
	})
	c1 <- m
	// return m
}

func GetmenuTue(url string, c2 chan []string) {
	doc, _ := goquery.NewDocument(url)
	var m []string
	doc.Find("div > div > section > table > tbody > tr").Each(func(_ int, s *goquery.Selection) {
		t := s.Find("td").First().Next().Text()
		m = append(m, t)
	})
	// time.Sleep(100 * time.Millisecond) //Mondayのgetを待ってる
	c2 <- m
	// return m
}

func GetmenuWen(url string, c3 chan []string) {
	doc, _ := goquery.NewDocument(url)
	var m []string
	doc.Find("div > div > section > table > tbody > tr").Each(func(_ int, s *goquery.Selection) {
		t := s.Find("td").First().Next().Next().Text()
		m = append(m, t)
	})
	// time.Sleep(200 * time.Millisecond) //Mondayのgetを待ってる
	c3 <- m
	// return m
}

func GetmenuThu(url string, c4 chan []string) {
	doc, _ := goquery.NewDocument(url)
	var m []string
	doc.Find("div > div > section > table > tbody > tr").Each(func(_ int, s *goquery.Selection) {
		t := s.Find("td").First().Next().Next().Next().Text()
		m = append(m, t)
	})
	// time.Sleep(300 * time.Millisecond) //Mondayのgetを待ってる
	c4 <- m
	// return m
}

func GetmenuFri(url string, c5 chan []string) {
	doc, _ := goquery.NewDocument(url)
	var m []string
	doc.Find("div > div > section > table > tbody > tr").Each(func(_ int, s *goquery.Selection) {
		t := s.Find("td").First().Next().Next().Next().Next().Text()
		m = append(m, t)
	})
	// time.Sleep(400 * time.Millisecond) //Mondayのgetを待ってる
	c5 <- m
	// return m
}

//並列処理でそれぞれの曜日のmenuをとってくる
func main() {
	url := "http://www.gakushoku.com/univ_mn2.php"
	c1 := make(chan []string)
	c2 := make(chan []string)
	c3 := make(chan []string)
	c4 := make(chan []string)
	c5 := make(chan []string)
	// defer close(c)
	go GetmenuMon(url, c1)
	go GetmenuTue(url, c2)
	go GetmenuWen(url, c3)
	go GetmenuThu(url, c4)
	go GetmenuFri(url, c5)
	res1, res2, res3, res4, res5 := <-c1, <-c2, <-c3, <-c4, <-c5

	// fmt.Println(res1[0])
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

	defer close(c1)
	defer close(c2)
	defer close(c3)
	defer close(c4)
	defer close(c5)
}
