package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func HandlerCountry(ctx *gin.Context) {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error while load env file : %s", err)
	}

	url := os.Getenv("URL_NINJAS")
	ApiKey := os.Getenv("API_KEY_NINJAS")
	ApiHost := os.Getenv("API_HOST_NINJAS")

	urlNew := fmt.Sprintf("%s/v1/country?name=United%%20States", url)
	fmt.Println(urlNew)
	req, _ := http.NewRequest("GET", urlNew, nil)

	req.Header.Add("X-RapidAPI-Key", ApiKey)
	req.Header.Add("X-RapidAPI-Host", ApiHost)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	defer res.Body.Close()

	fmt.Println(res)
	var resData []map[string]interface{}

	if err := json.NewDecoder(res.Body).Decode(&resData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": resData,
	})
}
