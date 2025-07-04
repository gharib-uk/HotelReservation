package api

import "github.com/gofiber/fiber/v2"

func HandleGetUsers(c *fiber.Ctx) error {
	return c.SendString("TODO")
}

func HandleGetUserByID(c *fiber.Ctx) error {
	return c.SendString("TODO")
}
