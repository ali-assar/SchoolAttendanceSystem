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

	store := db.New(database)
	handlers := handler.NewHandlers(store)

	handler.CreateDefaultAdmin(store)

	app := fiber.New()

	app.Use(logger.New())

	authMiddleware := handler.JWTAuthentication(store)

	app.Post("/login", handlers.HandleAuthenticate)

	apiv1 := app.Group("/api/v1", authMiddleware)

	// User routes
	apiv1.Post("user/", handlers.HandlePostUser)
	apiv1.Get("user/:id", handlers.HandleGetUserByID)
	apiv1.Get("user/", handlers.HandleGetAllUsers)
	apiv1.Put("user/:id", handlers.HandleUpdateUser)
	apiv1.Get("user/phone/:phone", handlers.HandleGetUserByPhoneNumber)
	apiv1.Get("user/name/:first_name/:last_name", handlers.HandleGetUserByName)
	apiv1.Delete("user/:id", handlers.HandleDeleteUserByID)

	// Attendance routes

	apiv1.Post("entrance/", handlers.HandlePostEntrance)
	apiv1.Post("exit/", handlers.HandlePostExit)
	apiv1.Put("entrance/:id", handlers.HandleUpdateEntrance)
	apiv1.Put("exit/:id", handlers.HandleUpdateExit)
	apiv1.Delete("entrance/:id", handlers.HandleDeleteEntrance)
	apiv1.Delete("exit/:id", handlers.HandleDeleteExit)

	apiv1.Get("attendance/", handlers.HandleGetTimeRange)
	apiv1.Get("attendance/user", handlers.HandleGetTimeRangeByUserID)
	apiv1.Get("absents/", handlers.HandleGetAbsentUsers)

	// apiv1.Get("attendance/:id", handlers.HandleGetUserWithEntranceExitTimeRangeByID)
	// apiv1.Get("attendance/time-range", handlers.HandleGetUsersEntranceExitForSpecificTime)
	// apiv1.Get("attendance/absent", handlers.HandleGetAbsentUsersForDay)

	// admin routes
	apiv1.Put("admin/:username/password", handlers.HandleUpdateAdmin)

	// Start server
	ip := os.Getenv("IP")
	if ip == "" {
		ip = "127.0.0.1:3000"
	}

	log.Fatal(app.Listen(ip))
}
