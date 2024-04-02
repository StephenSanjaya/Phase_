package handlers

import (
	"Phase2/week3/day2/NGC-11-12/config"
	"Phase2/week3/day2/NGC-11-12/entity"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetAllProducts(c echo.Context) error {
	products := new([]entity.Product)

	res := config.DB.Raw("SELECT * FROM products").Scan(&products)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, "product not found")
	}
	if res.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, res.Error)
	}

	return c.JSON(http.StatusOK, products)
}

func CreateNewTransaction(c echo.Context) error {
	trans := new(entity.Transaction)

	if err := c.Bind(&trans); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// set user id manually
	trans.UserID = 1

	err := config.DB.Transaction(func(tx *gorm.DB) error {
		type tempTrans struct {
			Stock int     `json:"stock"`
			Price float64 `json:"price"`
		}
		temp_trans := new(tempTrans)
		res := tx.Raw("SELECT stock, price FROM products WHERE product_id = ?", trans.ProductID).Scan(&temp_trans)
		if res.Error != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, res.Error)
		}

		if temp_trans.Stock < trans.Quantity {
			return echo.NewHTTPError(http.StatusBadRequest, "stock is not available")
		}

		var updated_stock = temp_trans.Stock - trans.Quantity
		res = tx.Model(&entity.Product{}).Where("product_id = ?", trans.ProductID).Update("stock", updated_stock)
		if res.Error != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, res.Error)
		}

		trans.TotalAmount = float64(trans.Quantity) * temp_trans.Price

		var depo float64
		res1 := tx.Raw("SELECT deposit_amount FROM users WHERE user_id = ?", trans.UserID).Scan(&depo)
		if res1.Error != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, res1.Error)
		}

		if depo < trans.TotalAmount {
			return echo.NewHTTPError(http.StatusBadRequest, "deposit amount is not enough")
		}

		var updated_depo = int(depo) - int(trans.TotalAmount)
		res = tx.Model(&entity.User{}).Where("user_id = ?", trans.UserID).Update("deposit_amount", updated_depo)
		if res.Error != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, res.Error)
		}

		res = tx.Omit("transaction_id").Create(&trans)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, res.Error)
		}
		if res.Error != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, res.Error)
		}

		return nil
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "success create transaction")
}
