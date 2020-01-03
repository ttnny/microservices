package apps

import (
	"github.com/ttnny/microservices-with-go/first_mvc/controllers"
	"log"
	"net/http"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
	}
}
