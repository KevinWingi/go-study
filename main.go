package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

type Person struct {
	Name string
	Id   string
	Age  int
}

func sum(a int, b int) (int, bool) {
	return a + b, a > b
}

func main() {
	// print()
	// server()
	structcase()
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HELLO WORLD"))
}

func servercase() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

func printcase() {
	myVar := "Hello  world"
	println(myVar)

	sum, status := sum(1, 0)

	fmt.Printf("Sum = %d, First greather than second = %s", sum, strconv.FormatBool((status)))
}

func structcase() {
	person := Person{
		Name: "Kevin Wingi",
		Id:   uuid.New().String(),
		Age:  28,
	}

	fmt.Printf("Name: %s, Id: %s, Age: %d", person.Name, person.Id, person.Age)
}
