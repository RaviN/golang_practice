package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	port := 3000
	_, err := startWebServer(port, 2)
	fmt.Println(err)
}

/*
Functions
*/
func startWebServer(port int, numberOfRetries int) (int, error) {
	fmt.Println("Starting server...")
	// do important things
	fmt.Println("Server started on port", port)
	return port, nil
}
