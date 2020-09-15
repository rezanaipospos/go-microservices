package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := UserDao.GetUser(0)
	assert.Nil(t, user, "we were not expecting a user with id 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t,"not_found", err.Code)
	assert.EqualValues(t,"User 0 was not found", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	user, err := UserDao.GetUser(123)
	assert.Nil(t,err)
	assert.NotNil(t,user)
	assert.EqualValues(t,123,user.Id)
	assert.EqualValues(t,"Reza",user.FirstName)
	assert.EqualValues(t,"Naipospos",user.LastName)
	assert.EqualValues(t,"rezanaipospos@gmail.com",user.Email)
}


