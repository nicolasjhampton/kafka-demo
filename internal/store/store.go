package store

import (
	"errors"
	"github.com/nicolasjhampton/kafka-demo/pkg/models"
)

var ErrUserNotFoundInProducer = errors.New("user not found")

func FindUserByID(id int, users []models.User) (models.User, error) {

	return models.User{ Name: "Bud"}, nil
}