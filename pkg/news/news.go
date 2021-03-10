package news

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/thiago-scherrer/hall9000/internal/config"
	"github.com/thiago-scherrer/hall9000/internal/voice"
	"golang.org/x/net/html"
)

const mfile string = "/tmp/.meio"

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
		log.Println(err)
	}

	defer response.Body.Close()

	f, err := os.OpenFile(mfile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		log.Println(err)
	}

	defer f.Close()

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
			log.Println(err)
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
				myRegex, _ := regexp.Compile(`(O que é o Meio)|(No que acreditamos)|(Consultar edições passadas)|(Sobre o Meio)|(O mundo transformado pelo 5G)(Curadoria de vídeos)|(Conversas com o Meio)|(Ponto de Partida)|(Colunas do Tony)|(Edições Premium)|(Todas as Edições)|(Última edição)|(Edições)|(Curadoria de vídeos)|(O mundo transformado pelo 5G)|(Quem somos)|(Política de privacidade)|(Assinantes)|(Acessar Premium)|(Benefícios da assinatura premium)|(Assinar Premium)|(Pioneiros)|(Monitor)|(Acessar Monitor)|(Como funciona o Monitor)|(Painel das Bolhas)|(Ajuda)|(Não recebi minha edição)|(Como cancelo o recebimento do Meio\?)|(Dúvidas sobre assinatura premium)|(Outras perguntas)|(Fale conosco)|(Acessar Premium)|(")`)
				altered := myRegex.ReplaceAllString(data, "")
				if len(altered) > 0 {
					w, err := f.WriteString(data)
					log.Println("got:", data)
					if err != nil {
						log.Println(w, err)
						break
					}
				}
			}
		}
	}

	content, err := ioutil.ReadFile(mfile)
	if err != nil {
		log.Println(err)
	}

	voice.Start(string(content))

	w := os.Remove(mfile)
	if w != nil {
		log.Println(w)
	}
}
