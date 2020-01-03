package controllers

import (
	"encoding/json"
	"github.com/ttnny/microservices-with-go/first_mvc/services"
	"log"
	"net/http"
	"strconv"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	userIdParam := req.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		_, _ = res.Write([]byte(err.Error()))
		return
	}

	user, err := services.GetUser(uint64(userId))
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		_, _ = res.Write([]byte(err.Error()))
		log.Println(err)
	}

	jsonValue, err := json.Marshal(user)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		_, _ = res.Write([]byte(err.Error()))
		log.Println(err)
	}

	_, err = res.Write(jsonValue)
	if err != nil {
		log.Println(err)
	}
}
