package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Ali-Assar/SchoolAttendanceSystem/issues/db"
	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) HandlePostEntrance(c *fiber.Ctx) error {
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

func (h *Handlers) HandlePostExit(c *fiber.Ctx) error {
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
	return c.Status(http.StatusOK).JSON(fmt.Sprintf("ID: %d", id))
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
	return c.Status(http.StatusOK).JSON(fmt.Sprintf("ID: %d", id))
}

func (h *Handlers) HandleDeleteEntrance(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	err = h.Store.DeleteEntrance(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(err)
	}
	return c.Status(http.StatusOK).JSON("deleted")
}

func (h *Handlers) HandleDeleteExit(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	err = h.Store.DeleteExit(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(err)
	}
	return c.Status(http.StatusOK).JSON("deleted")
}

func (h *Handlers) HandleGetTimeRange(c *fiber.Ctx) error {

	var getParams db.GetTimeRangeParams
	if err := c.BodyParser(&getParams); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	rows, err := h.Store.GetTimeRange(c.Context(), getParams)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(rows)
}

func (h *Handlers) HandleGetTimeRangeByUserID(c *fiber.Ctx) error {

	var getParams db.GetTimeRangeByUserIDParams
	if err := c.BodyParser(&getParams); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	rows, err := h.Store.GetTimeRangeByUserID(c.Context(), getParams)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(rows)
}

func (h *Handlers) HandleGetAbsentUsers(c *fiber.Ctx) error {
	// Define a struct to hold start_time and end_time from the request body
	type TimeRangeParams struct {
		StartTime int64 `json:"start_time"`
		EndTime   int64 `json:"end_time"`
	}

	// Parse the request body to get start_time and end_time
	var params TimeRangeParams
	if err := c.BodyParser(&params); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	// Call the SQLC-generated function with the time range
	rows, err := h.Store.GetAbsentUsers(c.Context(), params.StartTime, params.EndTime)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Return the result
	return c.Status(http.StatusOK).JSON(rows)
}
