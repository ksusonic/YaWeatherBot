package weather

import (
	"fmt"
	"strconv"
)

type Response struct {
	Now  uint64 `json:"now"`
	Info struct {
		Url           string `json:"url"`
		DefPressureMm int    `json:"def_pressure_mm"`
	} `json:"info"`
	Yesterday struct {
		Temp int `json:"temp"`
	} `json:"yesterday"`
	Fact      Fact       `json:"fact"`
	Forecasts []Forecast `json:"forecasts"`
}

type Fact struct {
	Temp       int     `json:"temp"`       // Температура (°C)
	FeelsLike  int     `json:"feels_like"` // Ощущаемая температура (°C).
	Condition  string  `json:"condition"`
	WindSpeed  float64 `json:"wind_speed"` // Скорость ветра (в м/с).
	WindDir    string  `json:"wind_dir"`
	WindGust   float64 `json:"wind_gust"`
	PressureMm int     `json:"pressure_mm"` // Давление (в мм рт. ст.).
	PressurePa int     `json:"pressure_pa"` // Давление (в гектопаскалях).
	Humidity   int     `json:"humidity"`    // Влажность воздуха (в процентах).
}

func (r Fact) WindDirReadable() string {
	return windDirReadable(r.WindDir)
}
func (r Fact) ConditionReadable() string {
	return conditionReadable(r.Condition)
}

type Forecast struct {
	Date    string `json:"date"`
	Sunrise string `json:"sunrise"`
	Sunset  string `json:"sunset"`
	Parts   struct {
		DayShort   part `json:"day_short"`
		NightShort part `json:"night_short"`
	} `json:"parts"`
}

type part struct {
	TempMin     *int    `json:"temp_min"`
	TempMax     *int    `json:"temp_max"`
	Temp        int     `json:"temp"`
	WindSpeed   float32 `json:"wind_speed"`
	WindDir     string  `json:"wind_dir"`
	PressureMm  int     `json:"pressure_mm"`
	PressurePa  int     `json:"pressure_pa"`
	Humidity    int     `json:"humidity"`
	PrecMm      float32 `json:"prec_mm"`
	PrecProb    int     `json:"prec_prob"`
	PrecPeriod  int     `json:"prec_period"`
	PrecType    int     `json:"prec_type"`
	Icon        string  `json:"icon"`
	Condition   string  `json:"condition"`
	FeelsLike   int     `json:"feels_like"`
	Polar       bool    `json:"polar"`
	FreshSnowMm float32 `json:"fresh_snow_mm"`
}

func (p part) SmartRange() string {
	if p.TempMin != nil && p.TempMax != nil && *p.TempMin != *p.TempMax {
		return fmt.Sprintf("от %d до %d", *p.TempMin, *p.TempMax)
	} else {
		return strconv.Itoa(p.Temp)
	}
}

func (p part) WindDirReadable() string {
	return windDirReadable(p.WindDir)
}
func (p part) ConditionReadable() string {
	return conditionReadable(p.Condition)
}

func windDirReadable(windDir string) string {
	switch windDir {
	case "nw":
		return "северо-западный"
	case "n":
		return "северный"
	case "ne":
		return "северо-восточный"
	case "e":
		return "восточный"
	case "se":
		return "юго-восточный"
	case "s":
		return "южный"
	case "sw":
		return "юго - западное"
	case "w":
		return "западное"
	case "c":
		return "штиль"
	default:
		return ""
	}
}

func conditionReadable(condition string) string {
	switch condition {
	case "clear":
		return "ясно ☀️"
	case "partly-cloudy":
		return "малооблачно 🌤"
	case "cloudy":
		return "облачно с прояснениями ⛅️"
	case "overcast":
		return "пасмурно ☁️"
	case "drizzle":
		return "морось 🌧"
	case "light-rain":
		return "небольшой дождь ☔️"
	case "rain":
		return "дождь ☔☔️🌧"
	case "moderate-rain":
		return "умеренно сильный дождь ☔️🌧"
	case "heavy-rain":
		return "сильный дождь ☔️🌧🌧"
	case "continuous-heavy-rain":
		return "длительный сильный дождь ☔️🌧🌧🌧"
	case "showers":
		return "ливень ☔️🌧🌧🌧🌧"
	case "wet-snow":
		return "дождь со снегом 🌧🌨"
	case "light-snow":
		return "небольшой снег 🌨"
	case "snow":
		return "снег 🌨🌨"
	case "snow-showers":
		return "снегопад 🌨🌨🌨"
	case "hail":
		return "град 🌨"
	case "thunderstorm":
		return "гроза 🌩"
	case "thunderstorm-with-rain":
		return "дождь с грозой 🌧🌩"
	case "thunderstorm-with-hail":
		return "гроза с градом 🌩🌧"
	default:
		return ""
	}
}
