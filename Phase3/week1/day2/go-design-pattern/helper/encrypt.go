package helper

import (
	"database/sql"
	"fmt"
	"sync"
)

type Database struct {
	Conn *sql.DB
}

var instance *Database
var once sync.Once

func GetInstance() *Database {
	once.Do(func() {
		db, err := sql.Open("lalalala", "asasas")
		if err != nil {
			fmt.Println("koneksi gagal")
		}
		instance = &Database{
			Conn: db,
		}
	})
	return instance
}
