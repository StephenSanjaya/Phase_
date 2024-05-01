package cmd

import (
	"Phase3/week2/day2/go-grpc-33/ms-order/handler"
	pb "Phase3/week2/day2/go-grpc-33/ms-order/internal/user"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func InitGrpc(UserHandler handler.UserHandler) {

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, UserHandler)

	// start gRPC server
	listen, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Println(err)
	}

	log.Println("Server listening on :50051")
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("INI GRPC")
}
