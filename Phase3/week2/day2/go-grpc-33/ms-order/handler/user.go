package handler

import (
	pb "Phase3/week2/day2/go-grpc-33/ms-order/internal/user"
	"Phase3/week2/day2/go-grpc-33/ms-order/models"
	"Phase3/week2/day2/go-grpc-33/ms-order/repo"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	repoUser repo.UserRepo
}

func NewHandler(repoUser repo.UserRepo) UserHandler {
	return UserHandler{
		repoUser: repoUser,
	}
}

func (uh UserHandler) CreateUser(ctx context.Context, data *pb.User) (*emptypb.Empty, error) {
	var empty *emptypb.Empty
	newId := primitive.NewObjectID()

	// set up data
	userData := models.User{
		Id:      newId,
		Name:    data.Name,
		Age:     data.Age,
		Address: data.Address,
	}

	// call insert function in repo
	err := uh.repoUser.Insert(userData)
	if err != nil {
		return empty, err
	}

	return new(emptypb.Empty), nil
}

/*
Notes :
jika menggunakan `*emptypb.Empty` sebagai return, maka hasil return akhir harus ditulis :
`new(emptypb.Empty)`

*/
