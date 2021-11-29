package telegram

import (
	"log"

    api "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendToBot(text string) error {

	bot, err := api.NewBotAPI("2082807550:AAGMLA7Vuy2e8jad0mgVw1Oz6IPkK60ry18")
    
    if err != nil {
        log.Fatalf("Didn't connection with bot: %v", err)
        return err
    }
    
    bot.Debug = true
    u := api.NewUpdate(0)
    u.Timeout = 60

    id := int64(-1001689708334)

    msg := api.NewMessage(id, text)
	
    if _, err := bot.Send(msg); err != nil {
        log.Fatalf("Problem with sending message: %v",err)
        return err
    }

    return err
}