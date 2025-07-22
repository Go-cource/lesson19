package main

const (
	arrd    = "localhost:8080"
	maxSize = 10000000
)

func foo() {
	for {
		var arr []int
		for i := 0; i < maxSize; i++ {
			arr = append(arr, i)
		}
	}
}
func main() {
	go foo()
}
