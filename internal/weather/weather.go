package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ksusonic/YaWeatherBot/config"
)

func GetForecast(cfg *config.Forecast) (string, error) {
	r, err := requestWeather(cfg)
	if err != nil {
		return "", fmt.Errorf("could not get weather: %w", err)
	}
	return fmt.Sprintf(
		"Сейчас в %s 🦉 %d, %s.\n"+
			"Облачность: %f\n"+
			"Давление: %d мм\n",
		cfg.ForecastPlaceNaming,
		r.Temp,
		r.ConditionHumanReadable,
		r.Cloudness,
		r.PressureMm,
	), nil
}

func requestWeather(cfg *config.Forecast) (*FactResponse, error) {
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

	var response FactResponse
	err = json.Unmarshal(all, &response)
	if err != nil {
		return nil, err
	}

	response.ConditionHumanReadable = conditionHumanReadable[response.Condition]
	return &response, nil
}
