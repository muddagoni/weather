package weather

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/weather/api"
	"github.com/weather/utils"

	"github.com/gorilla/mux"
)

func RegisterHandlers(r *mux.Router) {

	r.HandleFunc("/forecast", GetForecast).Methods("GET")

}

func GetForecast(w http.ResponseWriter, r *http.Request) {

	latitude := r.URL.Query().Get("latitude")

	longitude := r.URL.Query().Get("longitude")

	current_weather := r.URL.Query().Get("current_weather")
	currr_weather := strings.ToLower(current_weather)

	lat, err := strconv.ParseFloat(latitude, 64)
	if err != nil {
		api.JsonResponse(w, http.StatusInternalServerError, api.NewResponse(false, "Value of type 'Float' required for key 'longitude'.", fmt.Errorf("value of type 'Float' required for key 'latitude'")))
		return

	}
	if lat < -90 && lat > 90 {
		api.JsonResponse(w, http.StatusInternalServerError, api.NewResponse(false, "Latitude must be in range of -180 to 180°", fmt.Errorf("latitude must be in range of -90 to 90")))
		return
	}
	log, err := strconv.ParseFloat(longitude, 64)
	if err != nil {
		api.JsonResponse(w, http.StatusInternalServerError, api.NewResponse(false, "Value of type 'Float' required for key 'longitude'.", fmt.Errorf("value of type 'Float' required for key 'longitude'")))
		return

	}
	if log < -180 && lat < 180 {
		api.JsonResponse(w, http.StatusInternalServerError, api.NewResponse(false, "Longitude must be in range of -180 to 180°", fmt.Errorf("longitude must be in range of -180 to 180°")))
		return
	}

	cw, err := strconv.ParseBool(currr_weather)
	if err != nil {
		api.JsonResponse(w, http.StatusInternalServerError, api.NewResponse(false, "Value of type 'Bool' required for key 'current weather'.", fmt.Errorf("value of type 'Bool' required for key 'current weather'")))
		return
	}

	var res ForecastResponse

	params := map[string]string{
		"latitude":  fmt.Sprintf("%f", lat),
		"longitude": fmt.Sprintf("%f", log),
	}
	if cw {
		params["current_weather"] = "true"
	}

	url := "https://api.open-meteo.com/v1/forecast"
	utils.DoHTTPRequest(context.Background(), "GET", url, nil, params, nil, &res)

	api.JsonResponse(w, http.StatusOK, res)

}

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
