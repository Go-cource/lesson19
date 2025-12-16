package main

import (
	"log"
	"net/http"
)

func userHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {

	http.HandleFunc("/user/", userHandler)

	log.Println("User service started...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("8000 - err")
	}

}
