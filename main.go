package main

import (
	"oat431/shtlk-fiber/config"
	"oat431/shtlk-fiber/routes"
)

func main() {
	config.LoadEnvConfig()
	config.StartDatabase()
	routes.StartingApplication()
}
