package services

import (
	"net/http"
	"testing"

	"github.com/manyiu/go-microservices/mvc/domain"
	"github.com/manyiu/go-microservices/mvc/utils"
	"github.com/stretchr/testify/assert"
)

var (
	userDaoMock usersDaoMock

	getUserFunction func(int64) (*domain.User, *utils.ApplicationError)
)

func init() {
	domain.UserDao = &usersDaoMock{}
}

type usersDaoMock struct{}

func (m *usersDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userId)
}

func TestGetUserNotFoundInDatabase(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message:    "User 0 was not found",
		}
	}
	userId := int64(0)
	user, err := UsersService.GetUser(userId)

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "User 0 was not found", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			Id: 123,
		}, nil
	}

	user, err := UsersService.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
}
