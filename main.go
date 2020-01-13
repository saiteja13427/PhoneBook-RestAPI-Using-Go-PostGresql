package main

import (
	"PhoneBook/app"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main(){

	router := mux.NewRouter()

	router.Use(app.JWTAuthentication)
	port := os.Getenv("db_port")

	err := http.ListenAndServe(port, router)

	if err!= nil {
		fmt.Println(err)
	}
}
