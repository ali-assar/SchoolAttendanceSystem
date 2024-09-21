package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/Ali-Assar/SchoolAttendanceSystem/issues/db"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func CreateDefaultAdmin(store db.Querier) {
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

func (h *Handlers) HandleGetAdminByUserName(c *fiber.Ctx) error {
	userName := c.Params("username")
	admin, err := h.Store.GetAdminByUserName(c.Context(), userName)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"user_name": admin.UserName,
		"password":  admin.Password,
	})
}

func (h *Handlers) HandleUpdateAdmin(c *fiber.Ctx) error {
	var params db.UpdateAdminParams

	if err := c.BodyParser(&params); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	params.Password = string(hashedPassword)
	err = h.Store.UpdateAdmin(c.Context(), params)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update admin",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Admin password updated successfully",
	})
}
