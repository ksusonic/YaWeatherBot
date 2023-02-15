package weather

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/ksusonic/YaWeatherBot/config"
)

func GetForecast(cfg *config.Forecast) (string, error) {
	r, err := requestWeather(cfg)
	if err != nil {
		return "", fmt.Errorf("could not get weather: %w", err)
	}
	var message = randomTitle()
	message += fmt.Sprintf(
		"Сейчас в %s %d, %s\n",
		cfg.ForecastPlaceNaming, r.Fact.Temp, r.Fact.ConditionReadable(),
	)
	message += fmt.Sprintf(
		"Днем %s - %s\n",
		r.Forecasts[0].Parts.DayShort.SmartRange(), r.Forecasts[0].Parts.DayShort.ConditionReadable(),
	)
	message += fmt.Sprintf(
		"Вечером %s - %s",
		r.Forecasts[0].Parts.NightShort.SmartRange(), r.Forecasts[0].Parts.NightShort.ConditionReadable(),
	)
	return message, err
}

func requestWeather(cfg *config.Forecast) (*Response, error) {
	r, err := http.Get(cfg.BaseUrl + "?" +
		"lat=" + cfg.Lat +
		"&lon=" + cfg.Lon +
		"&lang=ru_RU",
	)

	if err != nil {
		return nil, err
	}
	all, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var response Response
	err = json.Unmarshal(all, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func randomTitle() string {
	return "Присылаю прогноз погоды! " + func() string {
		rand.Seed(time.Now().UnixNano())
		emoji := [][]int{
			{128513, 128591}, // Emoticons icons
			{128640, 128704}, // Transport and map symbols
		}
		r := emoji[rand.Int()%len(emoji)]
		min := r[0]
		max := r[1]
		n := rand.Intn(max-min+1) + min
		return html.UnescapeString("&#" + strconv.Itoa(n) + ";")
	}() + "\n"
}
