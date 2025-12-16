package main

import (
	"log"
	"net/http"
)

func greeterHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/greeter/", greeterHandler)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("8081 - error")
	}
}
