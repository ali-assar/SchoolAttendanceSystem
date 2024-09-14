package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Ali-Assar/SchoolAttendanceSystem/issues/db"
	"github.com/gofiber/fiber/v2"
)

// Entrance Handlers

func (h *Handlers) HandleCreateEntrance(c *fiber.Ctx) error {
	var postParams db.CreateEntranceParams
	if err := c.BodyParser(&postParams); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	id, err := h.Store.CreateEntrance(c.Context(), postParams)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(fmt.Sprintf("ID: %d", id))
}

func (h *Handlers) HandleGetEntrancesByUserID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	userID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	entrance, err := h.Store.GetEntrancesByUserID(c.Context(), userID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(entrance)
}

func (h *Handlers) HandleUpdateEntrance(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	var updateParams db.UpdateEntranceParams
	if err := c.BodyParser(&updateParams); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	updateParams.ID = id

	if err := h.Store.UpdateEntrance(c.Context(), updateParams); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(fmt.Sprintf("ID: %d updated", id))
}

func (h *Handlers) HandleDeleteEntrance(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	if err := h.Store.DeleteEntrance(c.Context(), id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(fmt.Sprintf("ID: %d deleted", id))
}

// Exit Handlers

func (h *Handlers) HandleCreateExit(c *fiber.Ctx) error {
	var postParams db.CreateExitParams
	if err := c.BodyParser(&postParams); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	id, err := h.Store.CreateExit(c.Context(), postParams)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(fmt.Sprintf("ID: %d", id))
}

func (h *Handlers) HandleGetExitsByUserID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	userID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	exit, err := h.Store.GetExitsByUserID(c.Context(), userID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(exit)
}

func (h *Handlers) HandleUpdateExit(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	var updateParams db.UpdateExitParams
	if err := c.BodyParser(&updateParams); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	updateParams.ID = id

	if err := h.Store.UpdateExit(c.Context(), updateParams); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(fmt.Sprintf("ID: %d updated", id))
}

func (h *Handlers) HandleDeleteExit(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	if err := h.Store.DeleteExit(c.Context(), id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(fmt.Sprintf("ID: %d deleted", id))
}