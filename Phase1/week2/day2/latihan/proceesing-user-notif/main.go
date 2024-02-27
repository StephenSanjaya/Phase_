package main

import (
	"fmt"
	"time"
)

type Notificaiton struct {
	UserId  int
	Message string
}

func sendEmailAysnc(userId int, msg string) {
	time.Sleep(2 * time.Second)
	fmt.Printf("Email notification send to user %d: %s\n", userId, msg)
}

func main() {
	notification := []Notificaiton{
		{101, "Your order has been confirmed."},
		{202, "Your account has been created."},
		{303, "Your paymenr was successful."},
	}

	for _, n := range notification {
		go sendEmailAysnc(n.UserId, n.Message)
	}

	fmt.Println("Main app continues...")

	time.Sleep(3 * time.Second)

	fmt.Println("Main app finish.")
}