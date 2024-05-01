package main

import (
	"Phase3/week2/day2/go-grpc/helpers"
	"Phase3/week2/day2/go-grpc/model"
	"Phase3/week2/day2/go-grpc/pb"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
)

type jwtClaims struct {
	Email string `json:"email"`
	Exp   int64  `json:"exp"`
	jwt.StandardClaims
}

type Handler struct {
	studentGRPC pb.StudentServiceClient
	// scoreGRPC pb.ScoreServiceClient
}

func NewHandler(studentGRPC pb.StudentServiceClient) *Handler {
	return &Handler{
		studentGRPC: studentGRPC,
	}
}

func (h *Handler) CreateStudent(c echo.Context) error {
	var payload model.Student

	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	fmt.Println("MASUK")

	// validation -> ex : validate email, password, token, dll

	in := pb.Request{
		Name:  payload.Name,
		Email: payload.Email,
	}

	response, err := h.studentGRPC.AddStudent(context.TODO(), &in)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, response)
}

func (h *Handler) Login(c echo.Context) error {
	var (
		payload model.Account
		token   string
	)

	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if payload.Email == "test@email.com" && payload.Password == "testong" {
		token = helpers.GenerateToken(payload.Email)
	} else {
		return c.JSON(http.StatusBadRequest, errors.New("Invalid Email & Password"))
	}

	fmt.Println(token)

	return c.JSON(http.StatusOK, token)
}

func main() {
	connection, err := grpc.Dial(":50001", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}

	// ScoreConnection, err := grpc.Dial(":50002", grpc.WithInsecure())
	// if err != nil {
	// 	log.Println(err)
	// }

	serviceClient := pb.NewStudentServiceClient(connection)

	// scoreClient := pb.NewScoreServiceClient(ScoreConnection)

	handler := NewHandler(serviceClient)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/login", handler.Login) // public api
	e.POST("/register", handler.CreateStudent)

	// private api
	// r := e.Group("/api/v1")
	// r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningKey: []byte("go-grpc"),
	// }))
	// r.POST("/register", handler.CreateStudent)

	err = e.Start(":1234")
	if err != nil {
		panic(err)
	}
}
