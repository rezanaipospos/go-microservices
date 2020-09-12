package services

import (
	"github.com/rezanaipospos/go-microservices/mvc/domain"
	"github.com/rezanaipospos/go-microservices/mvc/utils"
)

func GetUser(userId int64)(*domain.User,*utils.ApplicationError){
	return domain.GetUser(userId)
}