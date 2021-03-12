package daemon

import (
	"log"
	"os"
	"reflect"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/thiago-scherrer/hall9000/internal/config"
	"github.com/thiago-scherrer/hall9000/pkg/news"
	"github.com/thiago-scherrer/hall9000/pkg/tv"
	"github.com/thiago-scherrer/hall9000/pkg/volume"
	"github.com/thiago-scherrer/hall9000/pkg/weather"
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

		u := update.Message.From.UserName
		if reflect.DeepEqual(u, "thiago42") {
			log.Println(u)
		} else if reflect.DeepEqual(u, "karinaspd") {
			log.Println(u)
		} else {
			break
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
			case "clima":
				go weather.Start()
			case "canal":
				go tv.Canal(update.Message.CommandArguments())
			case "tvi":
				go tv.Tvi()
			case "tvd":
				go tv.Tvd()
			default:
				msg.Text = `Comandos validos: 
				/news (brasil, tec, meio, geral)
				/stop
				/volume (+/-)
				/clima
				/canal (1..500)
				/tvi (aumentar volume)
				/tvd (diminuir volume)
				`
			}
			bot.Send(msg)
		}

	}

}
