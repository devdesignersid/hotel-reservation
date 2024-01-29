package main

import (
	"flag"

	api "github.com/devdesignersid/hotel-reservation/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	port := flag.String("port", ":8000", "port to connect to")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("/api/v1")
	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)

	app.Listen(*port)
}
