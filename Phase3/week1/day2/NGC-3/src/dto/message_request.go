package dto

type Message struct {
	Message      string `json:"message"`
	UserReceiver User   `json:"user_receiver"`
	UserSender   User   `json:"user_sender"`
}

type User struct {
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
}
