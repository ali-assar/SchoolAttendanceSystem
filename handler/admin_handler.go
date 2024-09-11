package handler

import (
	"net/http"

	"github.com/Ali-Assar/SchoolAttendanceSystem/issues/db"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handlers) HandleGetAdminByUserName(c *fiber.Ctx) error {
	userName := c.Params("username")
	admin, err := h.Store.GetAdminByUserName(c.Context(), userName)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(admin)
}

func (h *Handlers) HandleUpdateAdmin(c *fiber.Ctx) error {
	var params db.UpdateAdminParams

	// Parse the request body into UpdateAdminParams
	if err := c.BodyParser(&params); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	// Hash the new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	// Update the password in the database
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
