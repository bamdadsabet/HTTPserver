package employee

import (
	"database/sql"
	"http-server/utils/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
	CreateEmployee(c *gin.Context)
	GetEmployeeList(c *gin.Context)
	GetEmployee(c *gin.Context)
	DeleteEmployee(c *gin.Context)
	UpdateEmployee(c *gin.Context)
}
type HandlerStruct struct {
	Repository EmployeeRepositoryInterface
}

func NewHandler (db *sql.DB) HandlerInterface {
	r := NewRepository(db)
	return HandlerStruct{r}
}

func (h HandlerStruct) CreateEmployee(c *gin.Context) {
	var reqBody EmployeeInfoStruct

	if ok := helper.GetReqBody(c, &reqBody); !ok  {
		return
	}

	res, RepositoryErr := h.Repository.Save(reqBody);

	if RepositoryErr != nil {
		helper.ResErr(c, http.StatusInternalServerError, RepositoryErr)
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"status":  http.StatusBadRequest,
		"massage": "Employee(id: "+ res.Id +") created",
		"data": res,
	})
}

func (h HandlerStruct) GetEmployeeList(c *gin.Context) {
	res, err := h.Repository.FindAll()

	if err != nil {
		helper.ResErr(c, http.StatusBadRequest, err)
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusBadRequest,
		"data": res,
	})
}

func (h HandlerStruct) GetEmployee(c *gin.Context) {
	id := c.Param("id")

	res, err := h.Repository.FindById(id)

	if err != nil {
		helper.ResErr(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusBadRequest,
		"data": res,
	})
}

func (h HandlerStruct) DeleteEmployee(c *gin.Context) {
	id := c.Param("id")

	h.Repository.DeleteById(id)
}

func (h HandlerStruct) UpdateEmployee (c *gin.Context) {
	id := c.Param("id")

	var reqBody EmployeeInfoStruct

	if ok := helper.GetReqBody(c, &reqBody); !ok  {
		return
	}

	h.Repository.Update(id, reqBody)
}

