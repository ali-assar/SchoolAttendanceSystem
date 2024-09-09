package main

import (
	"log"
	"os"

	"github.com/Ali-Assar/SchoolAttendanceSystem/issues/db"
	"github.com/Ali-Assar/SchoolAttendanceSystem/issues/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the database connection
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	defer database.Close()

	// Create tables if they do not exist
	err = db.CreateTables(database)
	if err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}

	// Create a new Fiber app
	app := fiber.New()

	// Use logger middleware for logging HTTP requests
	app.Use(logger.New())

	// Initialize the database store (repository)
	store := db.New(database)

	// Initialize the handlers with the store
	handlers := handler.NewHandlers(store)

	// Define API routes under /api/v1 prefix
	apiv1 := app.Group("/api/v1")

	// User routes
	apiv1.Post("user/", handlers.HandlePostUser)
	apiv1.Get("user/:id", handlers.HandleGetUserByID)
	apiv1.Get("user/", handlers.HandleGetAllUsers)
	apiv1.Put("user/:id", handlers.HandleUpdateUser)
	apiv1.Get("user/phone/:phone", handlers.HandleGetUserByPhoneNumber)
	apiv1.Get("user/name/:first_name/:last_name", handlers.HandleGetUserByName)
	apiv1.Delete("user/:id", handlers.HandleDeleteUserByID)

	// Attendance routes
	apiv1.Post("attendance/", handlers.HandlePostAttendance) // Create a new attendance record
	apiv1.Get("attendance/:user_id/:date", handlers.HandleGetAttendanceByUserIDAndDate) // Get attendance by user ID and date
	apiv1.Get("attendances/:date", handlers.HandleGetAllUsersAttendanceByDate)          // Get all users' attendance by date
	apiv1.Put("attendance/", handlers.HandleUpdateAttendanceByID)                       // Update an attendance record
	apiv1.Delete("attendance/:attendance_id", handlers.HandleDeleteAttendanceByID)      // Delete an attendance record by ID

	// Start the server on the specified port
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Default port if not specified
	}

	log.Fatal(app.Listen(":" + port))
}
