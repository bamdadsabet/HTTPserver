package crud

import (
	"database/sql"
	"http-server/helper"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/qustavo/dotsql"
)

type HandlerInterface interface {
	Create(c *gin.Context)
	GetList(c *gin.Context)
	Get(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
}

type HandlerStruct[T any] struct {
	Repository T
	TableName string
}

func NewCrudHandler[T, U any](db *sql.DB, table string) HandlerInterface {
	r := NewRepository[T, U](db)
	return HandlerStruct[RepositoryInterface[T, U]]{r, table}
}

func (h HandlerStruct[T]) Create(c *gin.Context) {
	var reqBody T

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

func (h HandlerStruct[T]) GetList(c *gin.Context) {
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

func (h HandlerStruct[T]) Get(c *gin.Context) {
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

func (h HandlerStruct[T]) Delete(c *gin.Context) {
	id := c.Param("id")

	h.Repository.DeleteById(id)
}

func (h HandlerStruct[T]) Update(c *gin.Context) {
	id := c.Param("id")

	var reqBody T

	if ok := helper.GetReqBody(c, &reqBody); !ok  {
		return
	}

	h.Repository.Update(id, reqBody)
}
