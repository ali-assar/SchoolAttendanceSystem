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

	app := fiber.New()
	app.Use(logger.New())

	userStore := db.New(database)

	userHandler := handler.NewUserHandler(userStore)
	apiv1 := app.Group("/api/v1")

	apiv1.Post("user/", userHandler.HandlePostUser)
	apiv1.Get("user/:id", userHandler.HandleGetUserByID)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Printf("Server running on port %s", port)
	err = app.Listen(":" + port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
