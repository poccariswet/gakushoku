package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, LINE Bot")
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
  bot, err := linebot.New(
		os.Getenv("LINE_Channel_Secret"), //linechannelsecret
		os.Getenv("LINE_Channel_Token"),  //linechanneltoken
	)
	if err != nil {
		log.Fatal(err)
	} else {
    fmt.Fprintf(w, "Maybe goooood!")
  }

	events, err := bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(111111)
		} else {
			w.WriteHeader(222222)
		}
		return
	}

	//メッセージの受信
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:

				if message.Text == "学食" {
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("update\nまたは\n'月','火','水','木','金'\nを入力してください")).Do(); err != nil {
						log.Print(err)

						for _, event := range events {
							if event.Type == linebot.EventTypeMessage {
								switch message := event.Message.(type) {
								case *linebot.TextMessage:
									switch message.Text {
									case "月":
										if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("こんにちは")).Do(); err != nil {
											log.Print(err)
										}
									}
								}
							}
						}

					}// messaeg.Text == "学食"

				} else {
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						log.Print(err)
					}
				}

			}//switch message...
		}//if event.Type...
	}//for _, event ...
}//callbackhandler...

var bot *linebot.Client

func main() {
  var err error
  bot, err = linebot.New(
    os.Getenv("LINE_Channel_Secret"), //linechannelsecret
    os.Getenv("LINE_Channel_Token"),  //linechanneltoken
  )
  if err != nil {
    log.Println(err)
  }

	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}
