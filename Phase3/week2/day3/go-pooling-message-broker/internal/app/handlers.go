package app

import (
	"Phase3/week2/day3/go-pooling-message-broker/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
)

func (a *App) getItems(c echo.Context) error {
	getItemsJson, err := a.rdb.Get(c.Request().Context(), "1").Result()
	items := new([]models.Item)
	var errGet error
	if err == redis.Nil {
		*items, errGet = a.Datastore.GetAllItems()
		if errGet != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve items")
		}
		itemsData, _ := json.Marshal(&items)
		a.rdb.Set(c.Request().Context(), "1", itemsData, 0)

		return c.JSON(http.StatusOK, items)
	} else if err != nil {
		return err
	}
	json.Unmarshal([]byte(getItemsJson), &items)

	return c.JSON(http.StatusOK, items)
}

func (a *App) getItem(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	item, err := a.Datastore.GetItemByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Item not found")
	}
	return c.JSON(http.StatusOK, item)
}

func (a *App) createItem(c echo.Context) error {
	item := new(models.Item)
	if err := c.Bind(item); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid item data")
	}
	if err := a.Datastore.InsertItem(item); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create item")
	}
	message, _ := json.Marshal(item)
	a.MessageQueue.PublishItemCreated(message)

	a.rdb.Del(c.Request().Context(), "1")

	return c.JSON(http.StatusCreated, item)
}

func (a *App) updateItem(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	item := new(models.Item)
	if err := c.Bind(item); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid item data")
	}
	if err := a.Datastore.UpdateItem(id, item); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update item")
	}
	message, _ := json.Marshal(item)
	a.MessageQueue.PublishItemUpdated(message)

	a.rdb.Del(c.Request().Context(), "1")

	return c.JSON(http.StatusOK, item)
}

func (a *App) deleteItem(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := a.Datastore.DeleteItem(id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete item")
	}
	a.MessageQueue.PublishItemDeleted(id)

	a.rdb.Del(c.Request().Context(), "1")

	return c.NoContent(http.StatusOK)
}
