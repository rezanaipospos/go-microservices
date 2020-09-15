package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rezanaipospos/go-microservices/mvc/services"
	"github.com/rezanaipospos/go-microservices/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context)  {
	userId, err := strconv.ParseInt(c.Param("user_id"),10,64)
	if err != nil{
		apiErr := &utils.ApplicationError{
			Message: "User id must be a number",
			StatusCode: http.StatusBadRequest,
			Code: "bad request",
		}

		c.JSON(apiErr.StatusCode,apiErr)
		return
	}
	user, apiErr := services.UsersService.GetUser(userId)
	if apiErr != nil{
		c.JSON(apiErr.StatusCode, apiErr)
		return
	}
	c.JSON(http.StatusOK, user)
}