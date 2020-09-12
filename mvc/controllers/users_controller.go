package controllers

import (
	"encoding/json"
	"github.com/rezanaipospos/go-microservices/mvc/services"
	"github.com/rezanaipospos/go-microservices/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request)  {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"),10,64)
	if err != nil{
		apiErr := &utils.ApplicationError{
			Message: "User id must be a number",
			StatusCode: http.StatusBadRequest,
			Code: "bad request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte(jsonValue))
		//just return bad request
		return
	}
	user, apiErr := services.GetUser(userId)
	if apiErr != nil{
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write([]byte(jsonValue))
		//handle the err and return to client
		return
	}
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}