package utils

import "github.com/gin-gonic/gin"

func Response(c *gin.Context, status int, body interface{}) {
	if c.GetHeader("Content-Type") == "application/xml" {
		c.XML(status, body)
	}

	c.JSON(status, body)
}

func ResponseError(c *gin.Context, err *ApplicationError) {
	if c.GetHeader("Content-Type") == "application/xml" {
		c.XML(err.StatusCode, err)
	}

	c.JSON(err.StatusCode, err)
}
