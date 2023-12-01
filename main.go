package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func sum(a int, b int) (int, bool) {
	return a + b, a > b
}

func main() {
	myVar := "Hello  world"
	println(myVar)

	sum, status := sum(1, 0)

	fmt.Printf("Sum = %d, First greather than second = %s", sum, strconv.FormatBool((status)))

	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HELLO WORLD"))
}
