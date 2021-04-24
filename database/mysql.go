package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log"
)

var DB *sql.DB

func Setup() {
	var err error
	dsn := fmt.Sprint("root:root@tcp(localhost:3306)/business-card?charset=utf8mb4&parseTime=True&loc=Local&readTimeout=5s")
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Panic("set up failed")
	}
}


