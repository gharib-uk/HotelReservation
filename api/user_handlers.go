package api

import (
	"HotelReservation/db"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	return c.SendString("TODO")
}

func (h *UserHandler) HandleGetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	ctx := c.Context()
	user, err := h.userStore.GetUserByID(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(user)
}
