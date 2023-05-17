package service

import (
	"context"

	"github.com/weather/internal/service/weather"
	"github.com/weather/utils"
)

func GetForecast(p map[string]string) weather.ForecastResponse {

	var res weather.ForecastResponse

	url := "https://api.open-meteo.com/v1/forecast"
	utils.DoHTTPRequest(context.Background(), "GET", url, nil, p, nil, &res)
	return res
}
