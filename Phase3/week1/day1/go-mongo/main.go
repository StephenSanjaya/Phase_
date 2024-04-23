package main

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Person struct
type Person struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `json:"name"`
	Age       int                `json:"age"`
	Address   string             `json:"address,omitempty"`
	CreatedAt time.Time          `json:"created_at"`
}

// Connection Database
func ConnectionDatabase(ctx context.Context) (*mongo.Collection, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}

	collection := client.Database("fieldrental").Collection("people")
	return collection, nil
}

// Create Data Person
func CreatePerson(c echo.Context) error {
	collection, err := ConnectionDatabase(context.Background())
	if err != nil {
		return err
	}

	p := new(Person)
	err = c.Bind(p)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	p.CreatedAt = time.Now()

	// proses insert data ke mongo
	result, err := collection.InsertOne(context.Background(), p)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, result)
}

// Get All Data Person
func GetAllPerson(c echo.Context) error {
	collection, err := ConnectionDatabase(context.Background())
	if err != nil {
		return err
	}

	var datas []Person

	// Proses get all data
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Looping Cursor
	for cursor.Next(context.Background()) {
		var data Person
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
func GetPersonById(c echo.Context) error {
	collection, err := ConnectionDatabase(context.Background())
	if err != nil {
		return err
	}

	// untuk mengubah tipe data string pada id menjadi Primitive.ObjectID -> WAJIB
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	var data Person

	// Proses get all data
	err = collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&data) // bson.M -> digunakan untuk filter data based on ...
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, data)
}

// Update Data Person
func UpdateDataPerson(c echo.Context) error {
	collection, err := ConnectionDatabase(context.Background())
	if err != nil {
		return err
	}

	// untuk mengubah tipe data string pada id menjadi Primitive.ObjectID -> WAJIB
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	p := new(Person)
	err = c.Bind(p)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Proses Update Data
	result, err := collection.UpdateOne(
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
func DeleteDataPerson(c echo.Context) error {
	collection, err := ConnectionDatabase(context.Background())
	if err != nil {
		return err
	}

	// untuk mengubah tipe data string pada id menjadi Primitive.ObjectID -> WAJIB
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Proses Delete Data
	result, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Untuk return 404
	if result.DeletedCount == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Data Not Found")
	}

	return c.JSON(http.StatusCreated, result)
}

func main() {
	e := echo.New()

	e.POST("/person", CreatePerson)
	e.GET("/all-person", GetAllPerson)
	e.GET("/person/:id", GetPersonById)
	e.PUT("/person/:id", UpdateDataPerson)
	e.DELETE("/person/:id", DeleteDataPerson)
	e.Logger.Fatal(e.Start(":1234"))
}
