package app

import (
	"Phase3/week2/day3/go-pooling-message-broker/internal/app/interfaces"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
)

const (
	redisAddr = "localhost:6379"
)

type App struct {
	Datastore    interfaces.Datastore
	MessageQueue interfaces.MessageQueue
	Echo         *echo.Echo
	rdb          *redis.Client
}

func SetupRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	return rdb
}

func NewApp(db interfaces.Datastore, mq interfaces.MessageQueue) *App {
	e := echo.New()
	rdb := SetupRedis()
	app := &App{
		Datastore:    db,
		MessageQueue: mq,
		Echo:         e,
		rdb:          rdb,
	}
	app.registerRoutes()
	return app
}

func (a *App) Start(port string) error {
	return a.Echo.Start(port)
}

func (a *App) registerRoutes() {
	a.Echo.GET("/items", a.getItems)
	a.Echo.GET("/items/:id", a.getItem)
	a.Echo.POST("/items", a.createItem)
	a.Echo.PUT("/items/:id", a.updateItem)
	a.Echo.DELETE("/items/:id", a.deleteItem)
}
