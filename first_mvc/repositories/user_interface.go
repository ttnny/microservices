package repositories

import (
	"fmt"
	"github.com/ttnny/microservices-with-go/first_mvc/models"
	"net/http"
)

// Dummy data
var (
	users = map[uint64]*models.User{
		123: {Id: 1, FirstName: "First", LastName: "Last", Email: "email@example.com"},
	}
)

func GetUser(userId uint64) (*models.User, *models.Error) {
	user := users[userId]
	if user != nil {
		return user, nil
	}

	return nil, &models.Error{
		Message:    fmt.Sprintf("User %v is not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
