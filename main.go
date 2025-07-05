package main

import (
	"HotelReservation/api"
	"HotelReservation/db"
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"os/signal"
)

const dbURI = "mongodb://localhost:27017"

var config = fiber.Config{
	// Override default error handler
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		// Status code defaults to 500
		code := fiber.StatusInternalServerError

		// Retrieve the custom status code if it's a *fiber.Error
		var e *fiber.Error
		if errors.As(err, &e) {
			code = e.Code
		}

		// Send custom error page
		err = ctx.Status(code).SendFile(fmt.Sprintf("./%d.html", code))
		if err != nil {
			// In case the SendFile fails
			return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		// Return from handler
		return nil
	},
}

func main() {
	client, err := mongo.Connect(context.TODO(),
		options.Client(), options.Client().ApplyURI(dbURI))

	if err != nil {
		log.Fatal(err)
	}

	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))

	listenAddr := flag.String("listen", ":3000", "Listen address")
	flag.Parse()

	app := fiber.New(config)
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Hotel Reservation Server!")
	})

	apiV1 := app.Group("/api/v1")
	apiV1.Get("/user", userHandler.HandleGetUsers)
	apiV1.Get("/user/:id", userHandler.HandleGetUserByID)

	//gracefully shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	go app.Listen(*listenAddr)

	<-sig
	return
}
