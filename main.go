package main

import (
	"HotelReservation/api"
	"flag"
	"github.com/gofiber/fiber/v2"
	"os"
	"os/signal"
)

func main() {
	listenAddr := flag.String("listen", ":3000", "Listen address")
	//listenAddress := os.Getenv("LISTEN_ADDR")
	// listenAddr = &listenAddress

	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Hotel Reservation Server!")
	})

	apiV1 := app.Group("/api/v1")
	apiV1.Get("/user", api.HandleGetUsers)
	apiV1.Get("/user/:id", api.HandleGetUserByID)

	//gracefully shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	go app.Listen(*listenAddr)

	<-sig
	return
}
