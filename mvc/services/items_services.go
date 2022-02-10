package services

import (
	"net/http"

	"github.com/manyiu/go-microservices/mvc/domain"
	"github.com/manyiu/go-microservices/mvc/utils"
)

type itemsService struct{}

var (
	ItemsService itemsService
)

func (i *itemsService) GetItem(itemId int64) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
	}
}
