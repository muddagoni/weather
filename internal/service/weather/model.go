package weather

type ForecastResponse struct {
	Latitude             float64        `json:"latitude"`
	Longitude            float64        `json:"longitude"`
	GenerationtimeMs     float64        `json:"generationtime_ms"`
	UtcOffsetSeconds     int            `json:"utc_offset_seconds"`
	Timezone             string         `json:"timezone"`
	TimezoneAbbreviation string         `json:"timezone_abbreviation"`
	Elevation            float64        `json:"elevation"`
	CurrentWeather       CurrentWeather `json:"current_weather,omitempty"`
}

type CurrentWeather struct {
	Temperature   float64 `json:"temperature,omitempty"`
	Windspeed     float64 `json:"windspeed,omitempty"`
	Winddirection float64 `json:"winddirection,omitempty"`
	Weathercode   int     `json:"weathercode,omitempty"`
	IsDay         int     `json:"is_day,omitempty"`
	Time          string  `json:"time,omitempty"`
}
