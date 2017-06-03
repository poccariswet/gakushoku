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

//並列処理でそれぞれの曜日のmenuをとってくる
func main() {
	url := "http://www.gakushoku.com/univ_mn2.php"
	c := make(chan []string)
	go GetmenuMon(url, c)
	go GetmenuTue(url, c)
	res1, res2 := <-c, <-c

	// fmt.Println(res1)
	for _, m1 := range res1 {
		fmt.Println(m1)
    // time.Sleep(2 * time.Second)
	}

  for _, m2 := range res2 {
	fmt.Println(m2)
  }
}
