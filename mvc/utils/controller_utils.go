package utils

import (
	"github.com/gin-gonic/gin"
)

func Respon(c *gin.Context, status int, body interface{})  {
	if c.GetHeader("Accept") == "application/xml"{
		c.XML(status,body)
		return
	}
	c.JSON(status,body)
}

func ResponError(c *gin.Context, err *ApplicationError)  {
	if c.GetHeader("Accept") == "application/xml"{
		c.XML(err.StatusCode, err)
		return
	}
	c.JSON(err.StatusCode,err)
}
