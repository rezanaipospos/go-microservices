package domain

import (
	"fmt"
	"github.com/rezanaipospos/go-microservices/mvc/utils"
	"log"
	"net/http"
)

var (
	users = map[int64]*User{
		123 :{Id:123, FirstName: "Reza", LastName: "Naipospos", Email:"rezanaipospos@gmail.com"},
	}
	UserDao usersServiceInterface
)

func init()  {
	UserDao = &userDao{}
}

type usersServiceInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct {
}

func (u *userDao)GetUser(userId int64) (*User , *utils.ApplicationError){
	log.Printf("Accessing database")
	if user := users[userId]; user != nil{
		//jika berhasil maka return User (data user), dan set error nil
		return user, nil
	}
	//jika error maka return User nil dan set pesan error.
	return  nil, &utils.ApplicationError{
		Message		: fmt.Sprintf("user %v does not exists", userId),
		StatusCode	: http.StatusNotFound,
		Code		: "not_found",
	}

}
