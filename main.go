package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const provider_token = "284685063:TEST:YTVhNzM2YmU4ZDdj"

func updateHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message != nil { // If we got a message
		//payam := update.Message.Text
		id := update.Message.From.ID

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				{
					msg := tgbotapi.NewMessage(id, "/newinvoice")
					bot.Send(msg)
				}
			case "newinvoice":
				{
					invoice := tgbotapi.NewInvoice(id, "تیشرت سفید", "با کیفیت و ارزان به رنگ سفید است", "123", provider_token, "", "USD", []tgbotapi.LabeledPrice{tgbotapi.LabeledPrice{Label: "USD", Amount: 2300}})
					invoice.SuggestedTipAmounts = []int{}
					invoice.NeedName = true
					invoice.NeedEmail = true
					invoice.NeedShippingAddress = true
					invoice.IsFlexible = true
					bot.Send(invoice)
				}
			}

		}

	}

	if update.ShippingQuery != nil {
		if update.ShippingQuery.ShippingAddress.CountryCode == "IR" {
			msg := tgbotapi.ShippingConfig{ShippingQueryID: update.ShippingQuery.ID, OK: false, ErrorMessage: "ارسال به ایران نداریم"}
			bot.Send(msg)
			return
		}
		msg := tgbotapi.ShippingConfig{ShippingQueryID: update.ShippingQuery.ID, OK: true, tgbotapi}
		bot.Send(msg)
	}
}

func main() {
	fmt.Println("")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("BOT_TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		go updateHandler(bot, update)
	}
}
