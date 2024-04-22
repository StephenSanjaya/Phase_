package handler

import (
	"Phase3/week1/day1/preview-phase2/dto"
	"Phase3/week1/day1/preview-phase2/entity"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type LoanService struct {
	db *gorm.DB
}

func NewLoanService(db *gorm.DB) *LoanService {
	return &LoanService{db: db}
}

func (ls *LoanService) Loan(c echo.Context) error {
	loan := new(entity.Loan)

	if err := c.Bind(&loan); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	loan.UserID = int(c.Get("user_id").(float64))

	if res := ls.db.Create(&loan); res.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, res.Error)
	}

	if res := ls.db.Model(&entity.User{}).Update("balance", loan.Limit).Where("user_id = ?", loan.UserID); res.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, res.Error)
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "success add limit loan",
		"loan":    loan,
	})
}

func (ls *LoanService) GetLimit(c echo.Context) error {
	limit := 0.0

	if res := ls.db.Select("limit").Where("user_id = ?", int(c.Get("user_id").(float64))).First(&limit); res.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, res.Error)
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "success add limit loan",
		"limit":   limit,
	})
}

func (ls *LoanService) WithdrawBalance(c echo.Context) error {
	w_balance := new(dto.WithdrawBalance)

	if err := c.Bind(&w_balance); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	saldo, err := GetSaldo(ls.db, int(c.Get("user_id").(float64)))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	saldo -= w_balance.Balance

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "success withdraw balance",
	})
}

func (ls *LoanService) PayLoanBalance(c echo.Context) error {

	user_id := int(c.Get("user_id").(float64))

	limit, err := GetLimit(ls.db, user_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if res := ls.db.Model(&entity.User{}).Update("balance", limit).Where("user_id = ?", user_id); res.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, res.Error)
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "success pay loan balance",
	})
}

func GetSaldo(db *gorm.DB, user_id int) (float64, error) {
	saldo := 0.0

	if res := db.Model(&entity.User{}).Where("user_id = ?", user_id).First(&saldo); res.Error != nil {
		return -1, res.Error
	}

	return saldo, nil
}

func GetLimit(db *gorm.DB, user_id int) (float64, error) {
	limit := 0.0

	if res := db.Model(&entity.Loan{}).Where("user_id = ?", user_id).First(&limit); res.Error != nil {
		return -1, res.Error
	}

	return limit, nil
}
