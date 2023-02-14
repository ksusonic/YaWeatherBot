package weather

type FactResponse struct {
	PrecProb               int     `json:"prec_prob"`
	SoilTemp               int     `json:"soil_temp"`
	Polar                  bool    `json:"polar"`
	PressurePa             int     `json:"pressure_pa"`
	PrecType               int     `json:"prec_type"`
	Uptime                 int     `json:"uptime"`
	SoilMoisture           float64 `json:"soil_moisture"`
	Daytime                string  `json:"daytime"`
	UvIndex                int     `json:"uv_index"`
	Source                 string  `json:"source"`
	Condition              string  `json:"condition"`
	ConditionHumanReadable string  `json:"-"`
	IsThunder              bool    `json:"is_thunder"`
	Season                 string  `json:"season"`
	Cloudness              float64 `json:"cloudness"`
	PrecStrength           int     `json:"prec_strength"`
	PressureMm             int     `json:"pressure_mm"`
	WindDir                string  `json:"wind_dir"`
	Icon                   string  `json:"icon"`
	WindSpeed              float64 `json:"wind_speed"`
	ObsTime                int     `json:"obs_time"`
	WindGust               float64 `json:"wind_gust"`
	Temp                   int     `json:"temp"`
	Humidity               int     `json:"humidity"`
	FeelsLike              int     `json:"feels_like"`
}

var conditionHumanReadable = map[string]string{
	"clear":                  "ясно",
	"partly-cloudy":          "малооблачно",
	"cloudy":                 "облачно с прояснениями",
	"overcast":               "пасмурно",
	"drizzle":                "морось",
	"light-rain":             "небольшой дождь",
	"rain":                   "дождь",
	"moderate-rain":          "умеренно сильный дождь",
	"heavy-rain":             "сильный дождь",
	"continuous-heavy-rain":  "длительный сильный дождь",
	"showers":                "ливень",
	"wet-snow":               "дождь со снегом",
	"light-snow":             "небольшой снег",
	"snow":                   "снег",
	"snow-showers":           "снегопад",
	"hail":                   "град",
	"thunderstorm":           "гроза",
	"thunderstorm-with-rain": "дождь с грозой",
	"thunderstorm-with-hail": "гроза с градом",
}
