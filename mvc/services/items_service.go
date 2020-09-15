package services

import (
	"github.com/rezanaipospos/go-microservices/mvc/domain"
	"github.com/rezanaipospos/go-microservices/mvc/utils"
	"net/http"
)

type itemsService struct {}
var (
	ItemsService itemsService
)

func (s *itemsService) GetItem(itemId string )(*domain.Item, *utils.ApplicationError)  {
	return nil, &utils.ApplicationError{
		Message: "implement me",
		StatusCode: http.StatusInternalServerError,
	}
}
