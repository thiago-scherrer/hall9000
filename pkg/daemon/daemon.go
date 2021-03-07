package daemon

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/thiago-scherrer/hall9000/internal/config"
	"github.com/thiago-scherrer/hall9000/pkg/news"
	"github.com/thiago-scherrer/hall9000/pkg/volume"
)

func Start() {
	tgkey := config.GetKey()
	bot, err := tgbotapi.NewBotAPI(tgkey)

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = config.GetDebug()

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = config.GetTimeOut()

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.From.UserName == "thiago42" || update.Message.From.UserName == "karinas" {
			continue
		}

		if update.Message.IsCommand() {

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "news":
				go news.Start(update.Message.CommandArguments())
			case "stop":
				os.Setenv("CONTROL", "true")
			case "volume":
				go volume.Start(update.Message.CommandArguments())
			case "withArgument":
				msg.Text = "You supplied the following argument: " + update.Message.CommandArguments()
			case "html":
				msg.ParseMode = "html"
				msg.Text = "This will be interpreted as HTML, click <a href=\"https://www.example.com\">here</a>"
			default:
				msg.Text = "I don't know that command"
			}
			bot.Send(msg)
		}

	}

}
