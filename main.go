package main

import (
	"log"
	"os"

	"github.com/Ali-Assar/SchoolAttendanceSystem/issues/db"
	"github.com/Ali-Assar/SchoolAttendanceSystem/issues/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",                            // Specify allowed origins
		AllowMethods: "GET,POST,PUT,DELETE",          // Specify allowed methods
		AllowHeaders: "Origin, Content-Type, Accept", // Specify allowed headers
	}))
	app.Use(logger.New())

	authMiddleware := handler.JWTAuthentication(store)

	app.Post("/login", handlers.HandleAuthenticate)

	// Routes requiring authentication
	apiv1 := app.Group("/api/v1", authMiddleware)

	apiv1.Post("teacher/", handlers.HandlePostTeacher)
	apiv1.Post("student/", handlers.HandlePostStudent)
	apiv1.Get("user/:id", handlers.HandleGetUserByID)
	apiv1.Get("teacher/:id", handlers.HandleGetTeacherByID)
	apiv1.Get("student/:id", handlers.HandleGetStudentByID)
	apiv1.Get("user/name/:first_name/:last_name", handlers.HandleGetUserByName)
	apiv1.Put("user/:id", handlers.HandleUpdateUser)
	apiv1.Put("student/:id", handlers.HandleUpdateStudentAllowedTime)
	apiv1.Put("teacher/:id", handlers.HandleUpdateTeacherAllowedTime)
	apiv1.Delete("user/:id", handlers.HandleDeleteUser)

	apiv1.Get("attendance/:date", handlers.GetAttendanceByDate)                           // Get attendance for a particular date
	apiv1.Get("attendance/range/:startDate/:endDate", handlers.GetAttendanceBetweenDates) // Get attendance between date A and B
	apiv1.Get("attendance/absent/:date", handlers.GetAbsentUsersByDate)
	apiv1.Get("attendance/absent/teacher/:date", handlers.GetAbsentTeachersByDate)

	apiv1.Put("admin/", handlers.HandleUpdateAdmin)

	// Routes without authentication (biometric service)
	biometric := app.Group("/biometric")

	biometric.Get("/", handlers.HandleGetUsersWithFalseBiometric)
	biometric.Get("/user", handlers.HandleGetUsersWithTrueBiometric)
	biometric.Put("/:id", handlers.HandleUpdateUserBiometric)
	biometric.Post("/attendance/", handlers.HandleAttendance)

	// Start server
	ip := os.Getenv("IP")
	if ip == "" {
		ip = "127.0.0.1:3000"
	}

	log.Fatal(app.Listen(ip))
}
