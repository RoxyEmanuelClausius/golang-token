package main

import (
	"log"
	"net/http"
	"token/models"
)

func main() {
	models.ConnectDatabase()

	http.HandleFunc("/login", login)
	http.HandleFunc("/home", home)
	http.HandleFunc("/refresh", refresh)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
