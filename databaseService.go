package main

import (
	"fmt"
	"github.com/deepch/RTSPtoWeb/config"
	_ "github.com/go-sql-driver/mysql"
)


func test() (string, error) {
	db := config.Conn

	err := db.Ping()
	if err != nil {
		fmt.Println("Ping Failed!! %s", err.Error())
	} else {
		fmt.Println("Successful database connection")
	}

	return "OK", nil
}