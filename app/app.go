package app

import (
	"database/sql"
	"http-server/db"
	"log"

	"github.com/gin-gonic/gin"
)

type AppStructure struct {
	DB *sql.DB
	Router *gin.Engine
}

var App AppStructure


func (a *AppStructure) CreateDBConnection() {
	db, err := db.ConnectDB()

	if err != nil {
		log.Fatal(err.Error())
		return 
	}

	a.DB = db
}

func (a * AppStructure) InitRouter(port string, registerRoutes func()) {
	r := gin.Default()
	a.Router = r
	registerRoutes()
	a.Router.Run(port)
}