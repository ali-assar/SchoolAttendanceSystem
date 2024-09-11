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

	sessionStore := session.New()

	authHandler := &handler.AuthHandler{Store: sessionStore}
	handlers := handler.NewHandlers(store)

	app.Post("/login", authHandler.Login)
	app.Post("/logout", authHandler.Logout)

	apiv1 := app.Group("/api/v1", authHandler.IsAuthenticated)

	apiv1.Post("user/", handlers.HandlePostUser)
	apiv1.Get("user/:id", handlers.HandleGetUserByID)
	apiv1.Get("user/", handlers.HandleGetAllUsers)
	apiv1.Put("user/:id", handlers.HandleUpdateUser)
	apiv1.Get("user/phone/:phone", handlers.HandleGetUserByPhoneNumber)
	apiv1.Get("user/name/:first_name/:last_name", handlers.HandleGetUserByName)
	apiv1.Delete("user/:id", handlers.HandleDeleteUserByID)

	apiv1.Post("attendance/", handlers.HandlePostAttendance)
	apiv1.Get("attendance/:user_id/:date", handlers.HandleGetAttendanceByUserIDAndDate)
	apiv1.Get("attendances/:date", handlers.HandleGetAllUsersAttendanceByDate)
	apiv1.Put("attendance/", handlers.HandleUpdateAttendanceByID)
	apiv1.Delete("attendance/:attendance_id", handlers.HandleDeleteAttendanceByID)

	ip := os.Getenv("IP")
	if ip == "" {
		ip = "127.0.0.1:3000"
	}

	log.Fatal(app.Listen(ip))
}
