package db

import (
	"composeapp/ent"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *ent.Client {
	client, err := ent.Open("mysql", "root:nopass@tcp(localhost:3306)/zerodb?parseTime=True")
	if err != nil {
		fmt.Println("failed opening connection to mysql: %v", err)
	}

	return client
}
