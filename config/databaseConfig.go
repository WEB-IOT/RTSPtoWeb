package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Conn *sql.DB

func init() {
	cfg, err := LoadConfig()
	if err != nil {
		panic(err.Error())
	}

	log.Print(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.mysqlConfig.Username,
		cfg.mysqlConfig.Password,
		cfg.mysqlConfig.URL,
		cfg.mysqlConfig.Port,
		cfg.mysqlConfig.Database))
		
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.mysqlConfig.Username,
		cfg.mysqlConfig.Password,
		cfg.mysqlConfig.URL,
		cfg.mysqlConfig.Port,
		cfg.mysqlConfig.Database))
	if err != nil {
		panic(err.Error())
	}

	Conn = db
}