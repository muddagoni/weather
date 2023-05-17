package weather

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/weather/api"
	"github.com/weather/internal/service"

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

	params := map[string]string{
		"latitude":  fmt.Sprintf("%f", lat),
		"longitude": fmt.Sprintf("%f", log),
	}
	if cw {
		params["current_weather"] = "true"
	}

	res := service.GetForecast(params)

	api.JsonResponse(w, http.StatusOK, res)

}
