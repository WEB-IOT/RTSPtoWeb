package main

import (
	"fmt"
	"github.com/deepch/RTSPtoWeb/config"
	"github.com/deepch/RTSPtoWeb/model"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

func test() (string, error) {
	db := config.Conn

	err := db.Ping()
	if err != nil {
		fmt.Println("Ping Failed!!", err.Error())
		return "", err
	} else {
		fmt.Println("Successful database connection")
	}

	return "OK", nil
}

func getUserByName(username string) (model.User, error) {
	db := config.Conn

	user := &model.User{}

	err := db.QueryRow("SELECT id, name, username, password, role, status FROM users WHERE username = ?", username).Scan(&user.ID, &user.Name, &user.Username, &user.Password, &user.Role, &user.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No user found with the given username: %s", username)
			return *user, err
		}
		log.Fatal(err)
		return *user, err
	}

	return *user, nil
}
