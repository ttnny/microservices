package repositories

import (
	"errors"
	"fmt"
	"github.com/ttnny/microservices-with-go/first_mvc/models"
)

// Dummy data
var (
	users = map[uint64]*models.User{
		123: {Id: 1, FirstName: "First", LastName: "Last", Email: "email@example.com"},
	}
)

func GetUser(userId uint64) (*models.User, error) {
	user := users[userId]
	if user != nil {
		return user, nil
	}

	return nil, errors.New(fmt.Sprintf("User %v is not found", userId))
}
