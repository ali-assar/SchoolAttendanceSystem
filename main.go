package main

import (
	"context"
	"log"
	"os"

	"github.com/Ali-Assar/SchoolAttendanceSystem/issues/db"
	"github.com/Ali-Assar/SchoolAttendanceSystem/issues/handler"
	sms "github.com/Ali-Assar/SchoolAttendanceSystem/issues/testsms"

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

	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",                                                        // Specify allowed origins
		AllowMethods: "GET,POST,PUT,DELETE,OPTION",                               // Specify allowed methods
		AllowHeaders: "Origin, Content-Type, Accept, Authorization, X-API-TOKEN", // Specify allowed headers
	}))
	authMiddleware := handler.JWTAuthentication(store)

	app.Post("/login", handlers.HandleAuthenticate)

	// Routes requiring authentication
	apiv1 := app.Group("/api/v1", authMiddleware)
	apiv1.Get("info/", handlers.HandleGetUserByJWT)

	// Teacher and Student routes
	apiv1.Post("teacher/", handlers.HandlePostTeacher)
	apiv1.Post("student/", handlers.HandlePostStudent)
	apiv1.Get("user/:id", handlers.HandleGetUserByID)
	apiv1.Get("teacher/:id", handlers.HandleGetTeacherByID)
	apiv1.Get("student/:id", handlers.HandleGetStudentByID)
	apiv1.Get("teacher/", handlers.HandleGetTeachers)
	apiv1.Get("student/", handlers.HandleGetStudents)
	apiv1.Get("user/name/:first_name/:last_name", handlers.HandleGetUserByName)
	apiv1.Put("student/:id", handlers.HandleUpdateStudent)
	apiv1.Put("teacher/:id", handlers.HandleUpdateTeacher)
	apiv1.Delete("user/:id", handlers.HandleDeleteUser)

	// Combined Attendance routes (all, teacher, student)
	apiv1.Get("attendance/:type/:date", handlers.HandleGetAttendanceByTypeAndDate)
	apiv1.Get("attendance/range/:type/:startDate/:endDate", handlers.HandleGetAttendanceByTypeAndDateRange)

	//update Attendance routes
	apiv1.Put("attendance/enter/:id/:enter_time", handlers.HandleUpdateEntranceByID)
	apiv1.Put("attendance/exit/:id/:exit_time", handlers.HandleUpdateExitByID)

	// Absent users and teachers
	apiv1.Get("attendance/absent/teacher/:date", handlers.HandleGetAbsentTeachersByDate)
	apiv1.Get("attendance/absent/student/:date", handlers.HandleGetAbsentStudentsByDate)

	// Admin routes
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

	go sms.ScheduleDailyAt(store, context.Background(), 10, 0)
	go sms.ScheduleDelayDailyAt(store, context.Background(), 10, 0)

	log.Fatal(app.Listen(ip))
}
