package config

import( 
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)



var DB *sql.DB

func ConnectDB() {
	var err error
	DB, err = sql.Open("mysql", "root:@tcp(localhost:3306)/test-login")
	if err != nil {
		panic(err.Error())
	}
}
