package main

import (
	"os"
	"prakerja3/configs"
	"prakerja3/routes"
)

func main() {
	configs.ConnectDatabase()
	e := routes.InitRoute()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Start(":" + port)
}
