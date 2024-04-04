package handlers

import (
	"Phase2/week3/day2/NGC-11-12-13/config"
	"Phase2/week3/day2/NGC-11-12-13/dto"
	externalapi "Phase2/week3/day2/NGC-11-12-13/external-api"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllStores(c echo.Context) error {
	products := new([]dto.Store)

	if res := config.DB.Omit("longitude", "latitude", "rating").Find(&products); res.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, res.Error.Error())
	}

	return c.JSON(http.StatusOK, products)
}

func GetStoreById(c echo.Context) error {
	id := c.Param("id")

	type Detail struct {
		WeatherData interface{} `json:"weather_data"`
		StoreDetail dto.Store   `json:"store_detail"`
	}

	detail := new(Detail)

	if res := config.DB.Where("store_id = ?", id).Find(&detail.StoreDetail); res.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, res.Error.Error())
	}

	detail.WeatherData = externalapi.GetWeatherData(c, detail.StoreDetail.Longitude, detail.StoreDetail.Latitude)

	return c.JSON(http.StatusOK, detail)
}
