package github_provider

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAuthorizationHeader(t *testing.T) {
	accessToken := "token_123"
	header := getAuthorizationHeader(accessToken)
	assert.Equal(t, "token token_123", header)
}

func TestDefer(t *testing.T) {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")

	fmt.Println("normal")
}
