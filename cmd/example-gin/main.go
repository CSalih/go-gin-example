package main

import "github.com/csalih/go-gin-example/internal/api"

func main() {
	r := api.NewRoutes()
	r.Run(":8080")
}
