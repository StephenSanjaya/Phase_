package controller

import (
	"Phase3/week1/day2/go-dependency-injection-04/model"
	"Phase3/week1/day2/go-dependency-injection-04/repository"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type ProductController struct {
	Repository repository.Product
}

func NewProductController(p repository.Product) ProductController {
	return ProductController{
		Repository: p,
	}
}

func (pc ProductController) Create(c echo.Context) error {
	var productReq model.Product
	//logic get data from body payload

	if productReq.Name != "Kursi" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"code":    400,
			"message": "harus selain kursi",
		})
	}

	RequestProduct := model.Product{
		Name: productReq.Name,
	}

	err := pc.Repository.Create(&RequestProduct)

	if err != nil {
		log.Fatalf("Error nih")
	}

	return c.JSON(http.StatusOK, echo.Map{
		"code":    200,
		"message": "success",
	})

}
