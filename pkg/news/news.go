package news

import (
	"context"
	"fmt"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/thiago-scherrer/hall9000/internal/config"
	"github.com/thiago-scherrer/hall9000/internal/voice"
)

func Start(font string) {

	config.GetControl()

	switch font {
	case "brasil":
		go brasil()
	case "mundo":
		go mundo()
	case "tec":
		go tec()
	case "geral":
		go geral()
	default:
		go geral()
	}

}

func geral() {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURLWithContext("http://pox.globo.com/rss/g1/", ctx)

	for _, it := range feed.Items {
		if config.GetControl() {
			break
		}

		fmt.Println(it.Title)
		voice.Start(it.Title)
	}
}

func brasil() {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURLWithContext("http://pox.globo.com/rss/g1/brasil/", ctx)

	for _, it := range feed.Items {
		if config.GetControl() {
			break
		}

		fmt.Println(it.Title)
		voice.Start(it.Title)
	}
}

func tec() {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURLWithContext("http://pox.globo.com/rss/g1/tecnologia/", ctx)

	for _, it := range feed.Items {
		if config.GetControl() {
			break
		}
		fmt.Println(it.Title)
		voice.Start(it.Title)
	}
}

func mundo() {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURLWithContext("http://pox.globo.com/rss/g1/mundo/", ctx)

	for _, it := range feed.Items {
		if config.GetControl() {
			break
		}
		fmt.Println(it.Title)
		voice.Start(it.Title)
	}
}
