package employee

import "http-server/app"


func RegisterRoutes() {
	h := NewHandler(app.App.DB)
	r := app.App.Router
	employee := r.Group("/employee")
	employee.GET("/", h.GetEmployeeList)
	employee.GET("/:id", h.GetEmployee)
	employee.PUT("/:id", h.UpdateEmployee)
	employee.POST("/", h.CreateEmployee)
	employee.DELETE("/:id", h.DeleteEmployee)
}
