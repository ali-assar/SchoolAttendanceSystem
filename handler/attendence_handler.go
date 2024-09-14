package handler

import (
	"fmt"
	"net/http"

	"github.com/Ali-Assar/SchoolAttendanceSystem/issues/db"
	"github.com/gofiber/fiber/v2"
)

// Entrance Handlers

func (h *Handlers) HandlePostEntrance(c *fiber.Ctx) error {
	var postParams db.CreateEntranceParams
	if err := c.BodyParser(&postParams); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	postParams.Date = (postParams.EnterTime / 86400) * 86400
	id, err := h.Store.CreateEntrance(c.Context(), postParams)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(http.StatusCreated).JSON(fmt.Sprintf("ID: %d", id))
}

type updateExit struct {
	db.GetAttendanceByUserIDAndDateParams
	db.UpdateExitParams
}

func (h *Handlers) HandleUpdateExit(c *fiber.Ctx) error {
	var args updateExit
	if err := c.BodyParser(&args); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	//check if attendance exit in database
	attendance, err := h.Store.GetAttendanceByUserIDAndDate(c.Context(), args.GetAttendanceByUserIDAndDateParams)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(err.Error())
	}

	args.UpdateExitParams.AttendanceID = attendance.AttendanceID

	err = h.Store.UpdateExit(c.Context(), args.UpdateExitParams)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(http.StatusOK).JSON(fmt.Sprintf("ID: %d", attendance.AttendanceID))

}
