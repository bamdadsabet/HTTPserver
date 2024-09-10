package db

import (
	"database/sql"
	"fmt"
	"http-server/utils/helper"

	_ "github.com/lib/pq"
)


type DBConfig struct {
	Name string
	User string
	Password string
	SSLMode string
}

func ConnectDB(c DBConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s", c.User, c.Name, c.Password, c.SSLMode)
	fmt.Println(connStr)
	db, err := sql.Open("postgres", connStr)
	helper.CHeckErr(err, "sql.open failed")
	if err := db.Ping(); err != nil {
		return nil, err
	}
	
	return db, err
}


