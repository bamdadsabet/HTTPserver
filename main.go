package main

import (
	"http-server/app"
	"http-server/domain/employee"

)


func main() {
	app.LoadDotENV()
	app.App.CreateDBConnection()
	app.App.InitRouter(func() {
		employee.RegisterRoutes()
	})
}