package dao

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/w1ndyz/gocamp/database"
	"github.com/w1ndyz/gocamp/ch02/xerr"
)

type User struct {
	ID int `json:"id"`
	Username string `json:"username"`
}

func InsertRow() error {
	_, err := database.DB.Exec("insert into users(username) values(?)", "张三")
	return err
}

func QueryRow() error {
	var user User
	err := database.DB.QueryRowContext(context.Background(), "select username from users where username = ?", "李思").Scan(&user)
	if err == sql.ErrNoRows {
		return fmt.Errorf("%w", xerr.ErrNoRows)
	}
	return nil
}
