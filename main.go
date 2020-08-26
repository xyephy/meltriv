package main
//import "fmt"

import (
	"log"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

const token = "TELEGRAM_TOKEN"

func main() {

	// Connect to Bot via Token
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "NOT FOUND"))
			continue
		}

	 //var userID = update.Message.From.ID
	//var user = userID

	//log.Printf("[%d] %s", userID, update.Message.Text)
	//log.Printf("Command: %s", update.Message.Command())

	switch update.Message.Command() {
		case "start":
			//cache[userID] = UserConfigurations{page:1}
			//cache[userID].resetFilter(userID)
			name := update.Message.From.UserName
			if update.Message.From.FirstName != "" {
				name = update.Message.From.FirstName
			}
			reply := "Hey " + name + ", welcome to Telegram. karibu sana"
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
			bot.Send(msg)
		}
	}
}
