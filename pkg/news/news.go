package news

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/thiago-scherrer/hall9000/internal/config"
	"github.com/thiago-scherrer/hall9000/internal/voice"
	"golang.org/x/net/html"
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
	case "meio":
		go canalmeio()
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

func canalmeio() {
	response, err := http.Get("https://www.canalmeio.com.br/ultima-edicao/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	textTags := []string{
		"a",
		"p", "span", "em", "string", "blockquote", "q", "cite",
		"h1", "h2", "h3", "h4", "h5", "h6",
	}

	tag := ""
	enter := false

	tokenizer := html.NewTokenizer(response.Body)
	for {
		if config.GetControl() {
			break
		}

		tt := tokenizer.Next()
		token := tokenizer.Token()

		err := tokenizer.Err()
		if err == io.EOF {
			break
		}

		switch tt {
		case html.ErrorToken:
			log.Fatal(err)
		case html.StartTagToken, html.SelfClosingTagToken:
			enter = false

			tag = token.Data
			for _, ttt := range textTags {
				if tag == ttt {
					enter = true
					break
				}
				if config.GetControl() {
					break
				}

			}
		case html.TextToken:
			if enter {
				data := strings.TrimSpace(token.Data)

				if len(data) > 0 {
					voice.Start(data)
				}
			}
		}
	}
}
