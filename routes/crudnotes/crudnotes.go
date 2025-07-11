package crudnotes

import (
	"NotesApp/database"
	"NotesApp/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateNote(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)
	noteData := new(models.RecievedNote)

	if err := c.BodyParser(&noteData); err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "invalid data",
		})
	}

	user_id, err := strconv.ParseUint(userId, 10, 64) // base 10, 64-bit unsigned

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "internal server",
		})
	}

	note := models.Note{
		UserID:      uint(user_id),
		Title:       noteData.Title,
		Description: noteData.Description,
		Tag:         noteData.Tag,
	}

	if err := database.DB.Create(&note).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create note"})
	}

	return c.JSON(note)
}

func GetNotes(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)
	user_id, err := strconv.ParseUint(userId, 10, 64) // base 10, 64-bit unsigned

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "internal server",
		})
	}

	var user models.User
	result := database.DB.Preload("Notes").First(&user, user_id)

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "No data found",
		})
	}

	notes := models.SendingNotes{Notes: user.Notes}

	return c.JSON(notes)
}

func UpdateNote(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)
	user_id, err := strconv.ParseUint(userId, 10, 64) // base 10, 64-bit unsigned

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "internal server",
		})
	}
	noteId := c.Locals("noteId").(string)
	note_id, err := strconv.ParseUint(noteId, 10, 64) // base 10, 64-bit unsigned

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "internal server",
		})
	}

	var updates models.RecievedNote
	if err := c.BodyParser(&updates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid note data",
		})
	}

	if updates.Title == "" ||
		len(updates.Title) < 3 {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid title length",
		})
	}
	if updates.Tag == "" ||
		len(updates.Tag) < 3 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid tag length",
		})
	}

	result := database.DB.Model(&models.Note{}).
		Where("id = ? AND user_id = ?", note_id, user_id).
		Updates(updates)

	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to update note",
		})
	}
	if result.RowsAffected == 0 {
		// No rows were updated - note either doesn't exist or doesn't belong to user
		return c.Status(404).JSON(fiber.Map{
			"error": "note not found or not owned by user",
		})

	}

	return c.JSON(updates)
}
func DeleteNote(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)
	user_id, err := strconv.ParseUint(userId, 10, 64) // base 10, 64-bit unsigned

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "internal server",
		})
	}
	noteId := c.Locals("noteId").(string)
	note_id, err := strconv.ParseUint(noteId, 10, 64) // base 10, 64-bit unsigned

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "internal server",
		})
	}

	result := database.DB.
		Where("id = ? AND user_id = ?", note_id, user_id).
		Delete(&models.Note{})

	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to update note",
		})
	}
	if result.RowsAffected == 0 {
		// No rows were updated - note either doesn't exist or doesn't belong to user
		return c.Status(404).JSON(fiber.Map{
			"error": "note not found or not owned by user",
		})

	}

	return c.JSON("deleted succesfully")
}
