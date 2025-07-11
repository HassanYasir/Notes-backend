package fetchnoteid

import "github.com/gofiber/fiber/v2"

func IdentifyNote(c *fiber.Ctx) error {
	note := c.Get("note")

	if note == "" {
		c.Status(401).JSON(fiber.Map{
			"error": "invalid noteId",
		})
	}
	c.Locals("noteId", note)
	return c.Next()
}
