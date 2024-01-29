package api

import (
	"github.com/devdesignersid/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "Sid",
		LastName:  "S Kumar",
	}
	return c.JSON(u)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"name": "Sid"})
}
