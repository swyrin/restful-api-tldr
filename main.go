package main

import (
	"fmt"
	"rest-client/routes"
)

func main() {
	routes.Setup()

	fmt.Println("The server is online @ http://localhost:7554")
	fmt.Println("OpenAPI is online @ http://localhost:7554/docs")
}
