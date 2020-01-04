package controllers

import (
	"encoding/json"
	"github.com/ttnny/microservices-with-go/first_mvc/models"
	"github.com/ttnny/microservices-with-go/first_mvc/services"
	"log"
	"net/http"
	"strconv"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	userIdParam := req.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		apiErr := &models.Error{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		jsonValue, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.StatusCode)
		_, _ = res.Write(jsonValue)
		return
	}

	user, apiErr := services.GetUser(uint64(userId))
	if apiErr != nil {
		res.WriteHeader(apiErr.StatusCode)
		_, _ = res.Write([]byte(apiErr.Message))
		log.Println(err)
		return
	}

	jsonValue, _ := json.Marshal(user)
	_, _ = res.Write(jsonValue)
}
