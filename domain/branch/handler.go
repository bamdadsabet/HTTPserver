package branch

import (
	"database/sql"
	"http-server/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)


type HandlerInterface interface {
	CreateBranch(c *gin.Context)
	GetBranchList(c *gin.Context)
	GetBranch(c *gin.Context)
	DeleteBranch(c *gin.Context)
	UpdateBranch(c *gin.Context)
}

type HandlerStruct struct {
	Repository BranchRepositoryInterface
}

func NewHandler (db *sql.DB) HandlerInterface {
	r := NewRepository(db)
	return HandlerStruct{r}
}


func (h HandlerStruct) CreateBranch(c *gin.Context) {
	var reqBody BranchInfoStruct

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
		"massage": "Branch(id: "+ res.Id +") created",
		"data": res,
	})
}

func (h HandlerStruct) GetBranchList(c *gin.Context) {
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

func (h HandlerStruct) GetBranch(c *gin.Context) {
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

func (h HandlerStruct) DeleteBranch(c *gin.Context) {
	id := c.Param("id")

	h.Repository.DeleteById(id)
}

func (h HandlerStruct) UpdateBranch (c *gin.Context) {
	id := c.Param("id")

	var reqBody BranchInfoStruct

	if ok := helper.GetReqBody(c, &reqBody); !ok  {
		return
	}

	h.Repository.Update(id, reqBody)
}

