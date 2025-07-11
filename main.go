package main

import (
	"NotesApp/database"
	"NotesApp/models"
	"NotesApp/routes/auth"
	"NotesApp/routes/crudnotes"
	fetchnoteid "NotesApp/routes/fetchnoteId"
	"NotesApp/routes/fetchuser"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Connect to SQLite
	database.Connect()

	// Auto-migrate the Todo model (creates table)
	database.DB.AutoMigrate(&models.Note{})
	database.DB.AutoMigrate(&models.User{})

	// Fiber app with middlewares
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New()) // Enable CORS

	// Routes
	app.Post("/api/login", auth.Login)
	app.Post("/api/signin", auth.Signin)
	app.Post("/api/getuser", fetchuser.Fetchuser, auth.Getuser)
	app.Post("/api/createnote", fetchuser.Fetchuser, crudnotes.CreateNote)
	app.Get("/api/getnotes", fetchuser.Fetchuser, crudnotes.GetNotes)
	app.Put("/api/updatenote", fetchuser.Fetchuser, fetchnoteid.IdentifyNote, crudnotes.UpdateNote)
	app.Delete("/api/deletenote", fetchuser.Fetchuser, fetchnoteid.IdentifyNote, crudnotes.DeleteNote)

	app.Listen(":3000")
}
