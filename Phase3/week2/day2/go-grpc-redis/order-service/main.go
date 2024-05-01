package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"

	pb "Phase3/week2/day2/go-grpc-redis/order-service/pb"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"gopkg.in/mgo.v2/bson"
)

const (
	port      = ":50052"
	redisAddr = "localhost:6379"
)

func SetupMongoDB() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func SetupRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	return rdb
}

type Handler struct {
	db  *mongo.Client
	rdb *redis.Client
}

func NewHandler(db *mongo.Client, rdb *redis.Client) *Handler {
	return &Handler{
		db:  db,
		rdb: rdb,
	}
}

func (s *Handler) GetOrder(ctx context.Context, in *pb.OrderIdRequest) (*pb.OrderResponse, error) {
	// Try to get the data from Redis first
	orderJson, err := s.rdb.Get(ctx, in.Id).Result()
	if err == redis.Nil {

		// Not found in Redis, fetch from MongoDB
		collection := s.db.Database("hacktiv8").Collection("orders")
		var order pb.OrderResponse
		objectID, err := primitive.ObjectIDFromHex(in.Id)
		if err != nil {
			fmt.Println("Error converting string to ObjectID:", err)
			return nil, err
		}

		err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&order)
		if err != nil {
			return nil, err
		}

		// Store the fetched data in Redis, serialize with JSON
		orderData, _ := json.Marshal(order)
		s.rdb.Set(ctx, in.Id, orderData, 0) // Using default expiration
		return &order, nil
	} else if err != nil {
		// Handle other potential Redis errors
		return nil, err
	}

	// Found in Redis, unmarshal JSON to return
	var cachedOrder pb.OrderResponse
	json.Unmarshal([]byte(orderJson), &cachedOrder)
	return &cachedOrder, nil
}

func (s *Handler) CreateOrder(ctx context.Context, in *pb.OrderRequest) (*pb.OrderResponse, error) {
	id := primitive.NewObjectID()
	idString := id.Hex()

	collection := s.db.Database("hacktiv8").Collection("orders")
	_, err := collection.InsertOne(ctx, bson.M{
		"_id":     id,
		"user_id": in.UserId,
		"product": in.Product,
		"status":  in.Status,
	})
	if err != nil {
		return nil, err
	}

	return &pb.OrderResponse{Id: idString, UserId: in.UserId, Product: in.Product, Status: in.Status}, nil
}

func main() {
	db := SetupMongoDB()
	rdb := SetupRedis()
	orderHandler := NewHandler(db, rdb)

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, orderHandler)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
