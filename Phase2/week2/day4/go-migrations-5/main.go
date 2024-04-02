package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	User_Id       int `gorm:"primaryKey:`
	Nama          string
	DepositAmount int
	Status        string
}

func main() {
	dsn := "host= localhost user=postgres password=12345 dbname=h8-test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	dbtrx := db.Transaction(func(tx *gorm.DB) error {

		patrick := &User{}
		result := tx.First(&patrick, 1)
		fmt.Println(result.Error, "<---4")
		if result.Error != nil {
			return err
		}
		patrick.DepositAmount = patrick.DepositAmount - 5000
		resultSavePatrick := tx.Model(&User{}).Where("user_id", 1).Updates(patrick)
		fmt.Println(resultSavePatrick.Error, "<---1")
		if resultSavePatrick.Error != nil {
			return err
		}

		plankton := &User{}
		result2 := tx.First(&plankton, 2)
		if result2.Error != nil {
			return err
		}
		plankton.DepositAmount = plankton.DepositAmount + 5000
		resultSavePlankton := tx.Model(&User{}).Where("user_id", 2).Updates(plankton)
		if resultSavePlankton.Error != nil {
			return err
		}

		return nil
	})

	if dbtrx != nil {
		log.Fatal(err)
	}

	// ....

}
