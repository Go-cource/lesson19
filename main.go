package main

import (
	"net/http"
	_ "net/http/pprof"
)

const (
	arrd    = ":8080"
	maxSize = 10000000
)

func foo() {
	for {
		var arr = make([]int, 0, maxSize)
		for i := 0; i < maxSize; i++ {
			arr = append(arr, i)
		}
	}
}
func main() {
	go foo()
	http.ListenAndServe(arrd, nil)
}
