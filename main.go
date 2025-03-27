package main

import (
	"context"
	"fmt"
	"rest-client/db"
	"rest-client/routes"
)

func main() {
	client := db.NewClient()
	_ = client.Prisma.Connect()

	defer func() {
		_ = client.Prisma.Disconnect()
	}()

	ctx := context.Background()

	routes.Setup(client, ctx)

	fmt.Println("The server is online @ http://localhost:7554")
	fmt.Println("OpenAPI is online @ http://localhost:7554/docs")
}
