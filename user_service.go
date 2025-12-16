package main

import (
	"log"
	"net/http"
)

func main() {
	
	err := http.ListenAndServe("8000", nil)
	if err != nil {
		log.Fatal("8000 - error")
	}
}
