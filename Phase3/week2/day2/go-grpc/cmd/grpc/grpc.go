package main

import (
	"Phase3/week2/day2/go-grpc/model"
	"Phase3/week2/day2/go-grpc/pb"
	"context"
	"fmt"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

func InitDatabase() (*sqlx.DB, error) {
	var (
		user     = "root"
		password = ""
		host     = "127.0.0.1"
		port     = "3306"
		dbname   = "test"
	)

	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbname)
	db, err := sqlx.Connect("mysql", source)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Connection Pool

	return db, nil
}

type Handler struct {
	db *sqlx.DB
}

func NewHandler(db *sqlx.DB) *Handler {
	return &Handler{
		db: db,
	}
}

func (h *Handler) AddStudent(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	var s = model.Student{
		Name:  in.Name,
		Email: in.Email,
	}

	fmt.Println(s.Name)
	fmt.Println(s.Email)

	var (
		args []interface{}
	)

	args = append(args, s.Name, s.Email)
	query := `INSERT INTO student (name, email) VALUES (?,?)`
	_, err := h.db.DB.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	var response = &pb.Response{
		Message: "Data berhasil dibuat",
	}

	return response, nil
}

func main() {
	db, err := InitDatabase()
	if err != nil {
		log.Println(err)
	}

	StudentHandler := NewHandler(db)

	grpcServer := grpc.NewServer()
	pb.RegisterStudentServiceServer(grpcServer, StudentHandler)

	fmt.Println("INI GRPC")

	// start gRPC server
	listen, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Println(err)
	}

	err = grpcServer.Serve(listen)
	if err != nil {
		log.Println(err)
	}

}
