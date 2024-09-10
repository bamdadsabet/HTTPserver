package helper

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CHeckErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func GetReqBody[T any](c *gin.Context, reqBody T) bool {
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		ResErr(c, http.StatusBadRequest, err)
		return false
	}
	return true
}

func ResErr(c *gin.Context, code int, err error) {
	c.JSON(code, gin.H{
		"status":  code,
		"massage": err.Error(),
	})
}

 