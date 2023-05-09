package main

import (
	"prakerja3/configs"
	"prakerja3/routes"
)

func main() {
	configs.ConnectDatabase()
	e := routes.InitRoute()
	e.Start(":8000")
}
