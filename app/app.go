package app

import (
	"database/sql"
	"fmt"
	"http-server/db"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)




type AppStructure struct {
	DB     *sql.DB
	Router *gin.Engine
}

var App AppStructure

func LoadDotENV() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
		return
	}
}

func (a *AppStructure) CreateDBConnection() {

	config := db.DBConfig{
		User:     os.Getenv("DB_USER"),
		Name:     os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	}

	fmt.Println(config)

	db, err := db.ConnectDB(config)

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	a.DB = db
}

func (a *AppStructure) InitRouter(registerRoutes func()) {
	r := gin.Default()
	a.Router = r
	registerRoutes()
	port := ":"+os.Getenv("PORT")
	a.Router.Run(port)
}
