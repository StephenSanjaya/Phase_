package externalapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetWeatherData(c echo.Context, longitude string, latitude string) interface{} {

	url := fmt.Sprintf("https://weather-by-api-ninjas.p.rapidapi.com/v1/weather?lat=%s&lon=%s&city=Jakarta&country=indonesia", latitude, longitude)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "110cf3aea9mshc617cd8f71c3120p1e16b1jsn925d5e0089f2")
	req.Header.Add("X-RapidAPI-Host", "weather-by-api-ninjas.p.rapidapi.com")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err.Error()
	}
	defer res.Body.Close()

	var body map[string]interface{}

	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		return err.Error()
	}

	return body

}
