package controller

import (
	"Phase3/week2/day2/NGC-5/ms-auth/model"
	pb "Phase3/week2/day2/NGC-5/ms-auth/pb/auth"
	"Phase3/week2/day2/NGC-5/ms-auth/repository"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthController struct {
	pb.UnimplementedAuthServiceServer
	repoAuth *repository.AuthRepository
}

func NewAuthController(repoAuth *repository.AuthRepository) *AuthController {
	return &AuthController{repoAuth: repoAuth}
}

func (ac *AuthController) RegisterAuth(ctx context.Context, userReq *pb.UserRequest) (*pb.UserResponse, error) {
	userPayload := model.User{
		Name:     userReq.Name,
		Email:    userReq.Email,
		Password: userReq.Password,
	}

	userRes, err := ac.repoAuth.Insert(&userPayload)
	if err != nil {
		return &pb.UserResponse{}, err
	}

	return &pb.UserResponse{
		Id:    userRes.Id.Hex(),
		Name:  userRes.Name,
		Email: userRes.Password,
	}, nil
}

func (ac *AuthController) GetUsers(ctx context.Context, empty *emptypb.Empty) (*pb.UsersResponse, error) {

	userRes, err := ac.repoAuth.FindAll()
	if err != nil {
		return &pb.UsersResponse{}, err
	}

	// ada cara yg lebih efisien?
	y := pb.UsersResponse{}
	for _, u := range *userRes {
		x := pb.UserResponse{
			Id:    u.Id.Hex(),
			Name:  u.Name,
			Email: u.Email,
		}
		y.Users = append(y.Users, &x)
	}

	return &y, nil
}
