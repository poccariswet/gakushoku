package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
  fmt.Println(t.Month())
	// fmt.Println(t.Day() + 7)
  fmt.Println(t.Day())
  fmt.Println(t.Hour())
  fmt.Println(t.Minute())
}
