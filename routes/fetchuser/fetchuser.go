package fetchuser

import (
	"NotesApp/Jwt"

	"github.com/gofiber/fiber/v2"
)

func Fetchuser(c *fiber.Ctx) error {
	token := c.Get("token")

	userId, err := Jwt.ValidateJWT(token)

	if err != nil {
		c.Status(401).JSON(fiber.Map{
			"error": "invalid token",
		})
	}
	c.Locals("userId", userId)
	return c.Next()
}
