package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/manyiu/go-microservices/mvc/services"
	"github.com/manyiu/go-microservices/mvc/utils"
)

func GetUser(c *gin.Context) {
	userIdString := c.Param("user_id")
	// log.Printf("About to process user_id %v", userId)

	userId, err := strconv.ParseInt(userIdString, 10, 64)

	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		utils.ResponseError(c, apiErr)

		return
	}

	user, apiErr := services.UsersService.GetUser(userId)

	if apiErr != nil {
		utils.ResponseError(c, apiErr)

		return
	}

	utils.Response(c, http.StatusOK, user)
}
