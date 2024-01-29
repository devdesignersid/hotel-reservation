package main

import (
	"context"
	"flag"

	api "github.com/devdesignersid/hotel-reservation/api/handlers"
	"github.com/devdesignersid/hotel-reservation/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbUri          = "mongodb://localhost:27017"
	dbName         = "hotel-reservation"
	userCollection = "users"
)

func main() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri))
	if err != nil {
		panic(err)
	}

	port := flag.String("port", ":8000", "port to connect to")
	flag.Parse()

	mongoUserStore := db.NewMongoUserStore(client)
	userHandler := api.NewUserHandler(mongoUserStore)

	var config = fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.JSON(map[string]string{"error": err.Error()})
		},
	}

	app := fiber.New(config)
	apiv1 := app.Group("/api/v1")
	apiv1.Get("/user/:id", userHandler.HandleGetUser)

	app.Listen(*port)
}
