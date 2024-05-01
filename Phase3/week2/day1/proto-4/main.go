package main

import (
	pb "Phase3/week2/day1/proto-4/internal/order"
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {
	order := &pb.Order{
		OrderId:     "1",
		OrderDate:   "2022-02-02",
		OrderStatus: "paid",
		SubTotal:    100000,
		Items: []*pb.ProductItem{
			{

				ProductName: "baju",
				Qty:         "2",
				Price:       "20000",
			},
			{

				ProductName: "celana",
				Qty:         "1",
				Price:       "80000",
			},
		},
	}

	data, err := proto.Marshal(order)
	if err != nil {
		log.Fatal("error while marshalling request")
	}
	fmt.Println(data)
	dataOrder := &pb.Order{}

	if err := proto.Unmarshal(data, dataOrder); err != nil {
		log.Fatal("error while Unmarshalling request")
	}
	// fmt.Println(dataOrder)
	// fmt.Println(dataOrder.GetItems())
	fmt.Println(dataOrder.GetOrderId())
}
