package branch

import "http-server/app"


func RegisterRoutes() {
	h := NewHandler(app.App.DB)
	r := app.App.Router
	employee := r.Group("/branch")
	employee.GET("/", h.GetBranchList)
	employee.GET("/:id", h.GetBranch)
	employee.PUT("/:id", h.UpdateBranch)
	employee.POST("/", h.CreateBranch)
	employee.DELETE("/:id", h.DeleteBranch)
}
