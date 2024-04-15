package handler

import (
	"Phase2/week4/day1/preview-week3/entity"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PlayerService struct {
	db *gorm.DB
}

func NewPlayerService(db *gorm.DB) *PlayerService {
	return &PlayerService{db: db}
}

func (ps *PlayerService) GetAllPlayers(c echo.Context) error {
	players := new([]entity.Player)

	if res := ps.db.Find(&players); res.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, res.Error.Error())
	}

	return c.JSON(http.StatusOK, players)
}

func (ps *PlayerService) CreateNewPlayers(c echo.Context) error {
	players := new(entity.Player)

	players.Username = "username 1"
	players.TeamName = "team name 1"
	players.Ranking = 1
	players.Score = 99

	if res := ps.db.Create(players); res.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, res.Error.Error())
	}

	return c.JSON(http.StatusOK, players)
}

func (ps *PlayerService) UpdateNewPlayers(c echo.Context) error {
	id := c.Param("id")

	players := new(entity.Player)

	if res := ps.db.Model(&players).Omit("player_id").Where("player_id = ?", id).Updates(entity.Player{Username: "updated username", TeamName: "updated team name", Ranking: 2, Score: 98}); res.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, res.Error.Error())
	}

	return c.JSON(http.StatusOK, players)
}
