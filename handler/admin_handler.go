package handler

import (
	"fmt"
	"net/http"

	"github.com/Ali-Assar/SchoolAttendanceSystem/issues/db"
	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) HandlePostAdmin(c *fiber.Ctx) error {
	var postParams db.CreateAdminParams
	if err := c.BodyParser(&postParams); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	adminUsername, err := h.Store.CreateAdmin(c.Context(), postParams)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(fmt.Sprintf("admin with username %s created", adminUsername))
}

func (h *Handlers) HandleGetAdminByUserName(c *fiber.Ctx) error {
	userName := c.Params("username")
	admin, err := h.Store.GetAdminByUserName(c.Context(), userName)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(admin)
}

func (h *Handlers) HandleDeleteAdmin(c *fiber.Ctx) error {
	userName := c.Params("username")
	err := h.Store.DeleteAdmin(c.Context(), userName)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(fmt.Sprintf("admin with username %s deleted", userName))
}
