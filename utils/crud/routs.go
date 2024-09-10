package crud

import (
	"database/sql"
	"http-server/app"
)


func RegisterRoutes[T, U any](group string, db *sql.DB, table string) {
	h := NewCrudHandler[T, U](db, table)
	r := app.App.Router
	employee := r.Group("/"+group)
	employee.GET("/", h.GetList)
	employee.GET("/:id", h.Get)
	employee.PUT("/:id", h.Update)
	employee.POST("/", h.Create)
	employee.DELETE("/:id", h.Delete)
}
