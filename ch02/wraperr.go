package main

import (
	"errors"
	"log"

	"github.com/w1ndyz/gocamp/ch02/dao"
	"github.com/w1ndyz/gocamp/ch02/xerr"
	"github.com/w1ndyz/gocamp/database"
	_ "gorm.io/driver/mysql"
)

func init() {
	database.Setup()
}

func main() {
	err := dao.InsertRow()
	if err != nil {
		log.Println(err.Error())
	}
	err = dao.QueryRow()
	if errors.Is(err, xerr.ErrNoRows) {
		log.Println(err.Error())
	}
}
