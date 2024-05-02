package cmd

import (
	"Phase3/week2/day2/NGC-5/ms-auth/controller"
	pb "Phase3/week2/day2/NGC-5/ms-auth/pb/auth"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func InitGrpc(AuthCtrler *controller.AuthController) {

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, AuthCtrler)

	// start gRPC server
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Println(err)
	}

	log.Printf("server listening at %v", listen.Addr())
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("INI GRPC")
}
