package main

import (
	"context"
	"log"
	"os"

	"github.com/Ali-Assar/SchoolAttendanceSystem/issues/db"
	"github.com/Ali-Assar/SchoolAttendanceSystem/issues/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
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

	createDefaultAdmin(store)

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
	apiv1.Post("attendance/", handlers.HandlePostAttendance)
	apiv1.Get("attendance/:user_id/:date", handlers.HandleGetAttendanceByUserIDAndDate)
	apiv1.Get("attendances/:date", handlers.HandleGetAllUsersAttendanceByDate)
	apiv1.Put("attendance/", handlers.HandleUpdateAttendanceByID)
	apiv1.Delete("attendance/:attendance_id", handlers.HandleDeleteAttendanceByID)

	// admin routes
	apiv1.Put("admin/:username/password", handlers.HandleUpdateAdmin)

	// Start server
	ip := os.Getenv("IP")
	if ip == "" {
		ip = "127.0.0.1:3000"
	}

	log.Fatal(app.Listen(ip))
}

func createDefaultAdmin(store db.Querier) {
	adminUsername := "admin"
	defaultPassword := "admin"

	_, err := store.GetAdminByUserName(context.Background(), adminUsername)
	if err != nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(defaultPassword), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Failed to hash password: %v", err)
		}

		adminParams := db.CreateAdminParams{
			UserName: adminUsername,
			Password: string(hashedPassword),
		}

		_, err = store.CreateAdmin(context.Background(), adminParams)
		if err != nil {
			log.Fatalf("Failed to create default admin: %v", err)
		}

		log.Println("Default admin created with username 'admin' and password 'admin'. Please change the password after login.")
	} else {
		log.Println("Admin already exists, skipping default admin creation.")
	}
}
