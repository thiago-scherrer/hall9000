package weather

import (
	"bytes"
	"fmt"
	"log"
	"text/template"

	owm "github.com/briandowns/openweathermap"
	"github.com/thiago-scherrer/hall9000/internal/config"
	"github.com/thiago-scherrer/hall9000/internal/voice"
)

const forecastTemplate = `A previsão para amanhã em {{.City.Name}} é:
{{range .List}}
{{range .Weather}} {{.Description}}{{end}}
Temperatura:        {{.Main.Temp}}  graus
Máxima:        {{.Main.TempMax}} graus
Mínima:         {{.Main.TempMin}} graus
{{end}}
`

func Start() {
	apiKey := config.GetClimaKey()

	temp(apiKey)
	prev(apiKey)

}

func temp(apiKey string) {
	w, err := owm.NewCurrent("C", "pt", apiKey)
	if err != nil {
		log.Fatalln(err)
	}

	w.CurrentByID(3458611)

	temp := fmt.Sprintf("%.2f", w.Main.Temp)
	tempMax := fmt.Sprintf("%.2f", w.Main.TempMax)
	tempMin := fmt.Sprintf("%.2f", w.Main.TempMin)

	voice.Start("A temperatura agora é de " + temp + " graus, a mínima é de " + tempMin + " graus e a máxima é de " + tempMax + " graus.")
}

func prev(apiKey string) {
	w, err := owm.NewForecast("5", "C", "pt", apiKey)
	fmt.Println(apiKey)
	if err != nil {
		log.Fatalln(err)
	}

	w.DailyByID(3458611, 1)

	data, _ := w.ForecastWeatherJson.(*owm.Forecast5WeatherData)

	tmpl, err := template.New("forecast").Parse(forecastTemplate)
	if err != nil {
		log.Fatalln(err)
	}

	var tpl bytes.Buffer
	tmpl.Execute(&tpl, data)

	voice.Start(tpl.String())

}
