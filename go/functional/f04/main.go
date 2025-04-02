package main

import "fmt"

func A() string {
	return "hello"
}

func B(a func() string) string {
	return a() + " world"
}

func main() {
	result := B(A)
	fmt.Println(result)
}
