package main

import "fmt"

// go vet error.go
func main() {
	movie_year := 1999
	movie_title := "The Matrix"
	fmt.Printf("In %s, %s was released.\n", movie_year, movie_title)
}
