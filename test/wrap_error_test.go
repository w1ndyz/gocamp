package test

import (
	"errors"
	"fmt"
	"github.com/w1ndyz/gocamp/database"
	"github.com/w1ndyz/gocamp/ch02/dao"
	"github.com/w1ndyz/gocamp/ch02/xerr"
	"testing"
)

func init() {
	database.Setup()
}


func TestWrapError(t *testing.T)  {
	err := dao.InsertRow()
	if err != nil {
		fmt.Println(err.Error())
	}
	err = dao.QueryRow()
	if errors.Is(err, xerr.ErrNoRows) {
		fmt.Println(err.Error())
	}
}
