package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "Phase3/week2/day2/go-grpc-redis/user-service/pb"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"gopkg.in/mgo.v2/bson"
)

const (
	port = ":50051"
)

func SetupMongoDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

type Handler struct {
	db *mongo.Client
}

func NewHandler(db *mongo.Client) *Handler {
	return &Handler{
		db: db,
	}
}

func (s *Handler) GetUser(ctx context.Context, in *pb.UserIdRequest) (*pb.UserResponse, error) {
	collection := s.db.Database("hacktiv8").Collection("users")
	var user pb.UserResponse

	objectIDString := in.Id

	objectID, err := primitive.ObjectIDFromHex(objectIDString)
	if err != nil {
		fmt.Println("Error converting string to ObjectID:", err)
		return nil, err
	}

	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (s *Handler) CreateUser(ctx context.Context, in *pb.UserRequest) (*pb.UserResponse, error) {
	id := primitive.NewObjectID()
	idString := id.Hex()

	collection := s.db.Database("hacktiv8").Collection("users")
	_, err := collection.InsertOne(ctx, bson.M{
		"_id":   id,
		"name":  in.Name,
		"email": in.Email,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{Id: idString, Name: in.Name, Email: in.Email}, nil
}

func main() {
	db := SetupMongoDB()
	userHandler := NewHandler(db)

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userHandler)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
