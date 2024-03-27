/*
Reference : https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/#validating-phone-numbers-emails-country-codes

*/

package main

import (
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	dsn := "host=localhost user=panjitamzil password=password dbname=ecommerce port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// var datas []Product

	// var data = Product{
	// 	Code:  "A001",
	// 	Price: 5000,
	// }

	// datas = append(datas, data)

	// data = Product{
	// 	Code:  "A002",
	// 	Price: 10000,
	// }

	// datas = append(datas, data)

	// insert data
	// db.Create(&datas)

	var p Product
	var products []Product
	fmt.Println("Before", p)
	fmt.Println("Before", products)

	// db.Where("id = ?", "1").First(&p) // Mengambil data paling atas
	// db.Last(&p) // Mengambil data paling bawah

	err = db.Find(&products).Error
	if err != nil {
		fmt.Println(gorm.ErrRecordNotFound)
		fmt.Println(errors.New("Not Found"))
	}

	fmt.Println("After", p)
	fmt.Println("After", products)
	fmt.Println("Done")

	// Create
	// db.Create(&Product{Code: "D42", Price: 100})

	// // Read
	// var product Product
	// db.First(&product, 1)                 // find product with integer primary key
	// db.First(&product, "code = ?", "D42") // find product with code D42

	// // Update - update product's price to 200
	// db.Model(&product).Update("Price", 200)
	// // Update - update multiple fields
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - delete product
	// db.Delete(&product, 1)
}
