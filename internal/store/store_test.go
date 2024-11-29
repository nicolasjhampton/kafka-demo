package store

import (
	"testing"
	"github.com/nicolasjhampton/kafka-demo/pkg/models"
)

func TestFindUserByID(t *testing.T) {

	t.Run("returns an error if no user is found", func(t *testing.T) {
		_, err := FindUserByID(0, []models.User{})
		if err != ErrUserNotFoundInProducer {
			t.Error("no error thrown when user not found")
		}
	})

	t.Run("returns an empty User model if no user is found", func(t *testing.T) {
		usr, _ := FindUserByID(0, []models.User{})
		if usr.Name != "" || usr.ID != 0 {
			t.Error("non empty user returned when matching user not found")
		}
	})
}