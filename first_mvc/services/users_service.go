package services

import (
	"github.com/ttnny/microservices-with-go/first_mvc/models"
	"github.com/ttnny/microservices-with-go/first_mvc/repositories"
)

func GetUser(userId uint64) (*models.User, error) {
	return repositories.GetUser(userId)
}
