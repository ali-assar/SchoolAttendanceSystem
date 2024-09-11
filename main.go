package main

import (
	"log"
	"os"

	"github.com/Ali-Assar/SchoolAttendanceSystem/issues/db"
	"github.com/Ali-Assar/SchoolAttendanceSystem/issues/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/joho/godotenv"
)

var store = session.New() // Session middleware for storing login information

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	defer database.Close()

	err = db.CreateTables(database)
	if err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}

	app := fiber.New()

	app.Use(logger.New())

	store := db.New(database)

	handlers := handler.NewHandlers(store)

	// Login route
	app.Post("/login", handleLogin)

	// Protect all routes with authentication middleware
	apiv1 := app.Group("/api/v1", authRequired)

	// User routes
	apiv1.Post("user/", handlers.HandlePostUser)
	apiv1.Get("user/:id", handlers.HandleGetUserByID)
	apiv1.Get("user/", handlers.HandleGetAllUsers)
	apiv1.Put("user/:id", handlers.HandleUpdateUser)
	apiv1.Get("user/phone/:phone", handlers.HandleGetUserByPhoneNumber)
	apiv1.Get("user/name/:first_name/:last_name", handlers.HandleGetUserByName)
	apiv1.Delete("user/:id", handlers.HandleDeleteUserByID)

	// Attendance routes
	apiv1.Post("attendance/", handlers.HandlePostAttendance)
	apiv1.Get("attendance/:user_id/:date", handlers.HandleGetAttendanceByUserIDAndDate)
	apiv1.Get("attendances/:date", handlers.HandleGetAllUsersAttendanceByDate)
	apiv1.Put("attendance/", handlers.HandleUpdateAttendanceByID)
	apiv1.Delete("attendance/:attendance_id", handlers.HandleDeleteAttendanceByID)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))
}

// handleLogin handles login for the admin
func handleLogin(c *fiber.Ctx) error {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Check if the username and password are correct
	if loginData.Username == "admin" && loginData.Password == "admin" {
		// Create a session
		sess, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create session"})
		}

		// Set session attribute
		sess.Set("authenticated", true)

		// Save session
		if err := sess.Save(); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save session"})
		}

		return c.JSON(fiber.Map{"message": "Login successful"})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
}

// authRequired middleware checks if the user is authenticated
func authRequired(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// Check if the user is authenticated
	if auth, ok := sess.Get("authenticated").(bool); !ok || !auth {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	return c.Next()
}
