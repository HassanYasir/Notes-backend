package auth

import (
	"NotesApp/Jwt"
	"NotesApp/database"
	"NotesApp/models"
	"NotesApp/validation"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequestData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type SigninRequestData struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseData struct {
	Data string `json:"data"`
}

func Login(c *fiber.Ctx) error {

	var data LoginRequestData
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if data.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "please enter conrrect details",
		})
	}

	if data.Email == "" || !validation.IsEmail(data.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "please enter conrrect details",
		})
	}

	var user models.User

	result := database.DB.Find(&user, "email = ?", data.Email)

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "unable to find record please enter correct details",
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))

	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "please enter conrrect details",
		})
	}

	// now person is authorized

	id := strconv.FormatUint(uint64(user.ID), 10)

	token, err := Jwt.GenerateJWT(id, 24)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "some internal error please try again later"})
	}

	return c.JSON(ResponseData{Data: token})

}

func Signin(c *fiber.Ctx) error {
	var data SigninRequestData
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if data.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "please enter conrrect details",
		})
	}

	if data.Email == "" && !validation.IsEmail(data.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "please enter conrrect details",
		})
	}
	if !validation.IsPassLength(data.Password, 8) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "please set atleast 8 characters long password",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error please try again",
		})
	}

	var user models.User

	user.Name = data.Name
	user.Email = data.Email
	user.Password = string(hashedPassword)

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create todo"})
	}

	id := strconv.FormatUint(uint64(user.ID), 10)

	token, err := Jwt.GenerateJWT(id, 24)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "some internal error please try again later"})
	}

	return c.JSON(ResponseData{Data: token})
}

func Getuser(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)
	var user models.User
	if err := database.DB.Select("id", "name", "email").First(&user, "id = ?", userId).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("error: %v", err))
	}

	return c.JSON(user)
}
