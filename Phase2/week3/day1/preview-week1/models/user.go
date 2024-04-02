package models

import "time"

type User struct {
	UserID    int       `json:"user_id,omitempty" gorm:"primaryKey;autoIncrement;not null"`
	Username  string    `json:"username" gorm:"not null;unique"`
	Email     string    `json:"email" gorm:"not null;unique"`
	Password  string    `json:"password" gorm:"not null"`
	Age       int       `json:"age" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Photo struct {
	PhotoID   int       `json:"photo_id,omitempty" gorm:"primaryKey;autoIncrement;not null"`
	Title     string    `json:"title" gorm:"not null"`
	Caption   string    `json:"caption" gorm:"not null"`
	Age       int       `json:"age" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CREATE TABLE photos (
//     photo_id SERIAL PRIMARY KEY,
//     title VARCHAR(255) NOT NULL,
//     caption VARCHAR(255) NOT NULL,
//     photo_url VARCHAR(255) NOT NULL,
//     user_id INT NOT NULL,
//     created_at DATE,
//     updated_at DATE,
//     FOREIGN KEY (user_id) REFERENCES users(user_id)
// );
