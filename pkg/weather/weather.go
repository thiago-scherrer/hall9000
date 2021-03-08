package weather

import (
	"fmt"
	"log"

	owm "github.com/briandowns/openweathermap"
	"github.com/thiago-scherrer/hall9000/internal/config"
	"github.com/thiago-scherrer/hall9000/internal/voice"
)

func Start() {
	apiKey := config.GetClimaKey()

	w, err := owm.NewCurrent("C", "pt", apiKey)
	if err != nil {
		log.Fatalln(err)
	}

	w.CurrentByID(3458611)

	temp := fmt.Sprintf("%f", w.Main.Temp)
	tempMax := fmt.Sprintf("%f", w.Main.TempMax)
	tempMin := fmt.Sprintf("%f", w.Main.TempMin)

	voice.Start("A temperatura agora é de " + temp + " graus, a mínima é de " + tempMin + " graus e a máxima é de " + tempMax + " graus.")

}
