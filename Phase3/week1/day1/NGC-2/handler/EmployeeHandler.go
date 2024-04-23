package handler

import (
	"Phase3/week1/day1/NGC-2/entity"
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type EmployeeService struct {
	coll *mongo.Collection
}

func NewEmployeeService(coll *mongo.Collection) *EmployeeService {
	return &EmployeeService{coll: coll}
}

// Create Data Person
func (es *EmployeeService) CreateEmployee(c echo.Context) error {
	e := new(entity.Employee)
	if err := c.Bind(e); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// proses insert data ke mongo
	result, err := es.coll.InsertOne(context.Background(), e)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}

// Get All Data Person
func (es *EmployeeService) GetAllEmployee(c echo.Context) error {
	page := c.Request().URL.Query().Get("page")
	limit := c.Request().URL.Query().Get("limit")
	sort := c.Request().URL.Query().Get("sort")

	var cursor *mongo.Cursor
	var errCursor error
	var sorting primitive.D
	if page != "" && limit != "" && sort != "" {
		if sort == "desc" {
			sorting = bson.D{{Key: "name", Value: -1}}
		} else {
			sorting = bson.D{{Key: "name", Value: 1}}
		}

		page, _ := strconv.ParseInt(page, 10, 32)
		limit, _ := strconv.ParseInt(limit, 10, 32)

		l := int64(limit)
		skip := int64(page*limit - limit)
		fOpt := options.FindOptions{Limit: &l, Skip: &skip, Sort: &sorting}

		cursor, errCursor = es.coll.Find(context.Background(), bson.D{{}}, &fOpt)
	} else {
		cursor, errCursor = es.coll.Find(context.Background(), bson.D{{}})
	}
	var datas []entity.Employee
	if errCursor != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errCursor.Error())
	}

	// Looping Cursor
	for cursor.Next(context.Background()) {
		var data entity.Employee
		if err := cursor.Decode(&data); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// proses append "data" ke dalam "datas"
		datas = append(datas, data)
	}

	if err := cursor.Err(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, datas)
}

// Get Data based on Id
func (es *EmployeeService) GetEmployeeById(c echo.Context) error {
	// untuk mengubah tipe data string pada id menjadi Primitive.ObjectID -> WAJIB
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var data entity.Employee

	// Proses get all data
	err = es.coll.FindOne(context.Background(), bson.M{"_id": id}).Decode(&data) // bson.M -> digunakan untuk filter data based on ...
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, data)
}

// Update Data Person
func (es *EmployeeService) UpdateDataEmployee(c echo.Context) error {
	// untuk mengubah tipe data string pada id menjadi Primitive.ObjectID -> WAJIB
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	p := new(entity.Employee)
	if err := c.Bind(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Proses Update Data
	result, err := es.coll.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		bson.M{"$set": p},
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}

// Delete data Person
func (es *EmployeeService) DeleteDataEmployee(c echo.Context) error {
	// untuk mengubah tipe data string pada id menjadi Primitive.ObjectID -> WAJIB
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Proses Delete Data
	result, err := es.coll.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Untuk return 404
	if result.DeletedCount == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Data Not Found")
	}

	return c.JSON(http.StatusCreated, result)
}
