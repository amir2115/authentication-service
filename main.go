package main

import (
	"AuthenticationService/config"
	"AuthenticationService/routes"
)

func main() {
	config.ConnectDatabase()
	r := routes.SetupRouter()

	r.Run("0.0.0.0:8000")
}
