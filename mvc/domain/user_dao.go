package domain

import (
	"fmt"
	"net/http"

	"github.com/manyiu/go-microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {
			Id:        123,
			FirstName: "Michael",
			LastName:  "Lee",
			Email:     "abcd@example.com",
		},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {

	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("User %v was not fount", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_fount",
	}
}
