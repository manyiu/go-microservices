package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/manyiu/go-microservices/mvc/services"
	"github.com/manyiu/go-microservices/mvc/utils"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userIdString := req.URL.Query().Get("user_id")
	// log.Printf("About to process user_id %v", userId)

	userId, err := strconv.ParseInt(userIdString, 10, 64)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		jsonValue, _ := json.Marshal(apiErr)

		resp.WriteHeader(apiErr.StatusCode)
		resp.Write([]byte(jsonValue))
		return
	}

	user, apiErr := services.UsersService.GetUser(userId)

	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
