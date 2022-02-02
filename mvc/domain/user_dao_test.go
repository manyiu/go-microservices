package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "We were not expecting a user with id 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "User 0 was not found", err.Message)

	// if user != nil {
	// 	t.Error(("We were not expecting a user with id 0"))
	// }

	// if err == nil {
	// 	t.Error("We were expecting an error when user id is 0")
	// }

	// if err != nil && err.StatusCode != http.StatusNotFound {
	// 	t.Error("We were expecting 404 when user is not found")
	// }
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Michael", user.FirstName)
	assert.EqualValues(t, "Lee", user.LastName)
	assert.EqualValues(t, "abcd@example.com", user.Email)
}
