//keywordを打ち込むとそのmethodの実行 "学食更新 or update cafe"
package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func update() {
	fmt.Println("アップデートします")
	time.Sleep(time.Second)
	fmt.Println("done update")
}

func main() {
	// url := "http://www.gakushoku.com/univ_mn2.php"
	s := bufio.NewScanner(os.Stdin)
	s.Scan()

	key := s.Text()
	if key == "学食更新" {
		update()
	} else if key == "update cafe" {
    update()
  } else {
    fmt.Println("終わります")
  }
}
