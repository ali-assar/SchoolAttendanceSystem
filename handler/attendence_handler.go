package handler

import (
	"fmt"
	"net/http"

	"github.com/Ali-Assar/SchoolAttendanceSystem/issues/db"
	"github.com/gofiber/fiber/v2"
)

/*

// Entrance Handlers
func (h *Handlers) HandlePostEntrance(c *fiber.Ctx) error {
	var postParams db.CreateEntranceParams
	if err := c.BodyParser(&postParams); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	postParams.Date = (postParams.EnterTime / 86400) * 86400

	_, err := h.Store.GetAttendanceByUserIDAndDate(c.Context(), db.GetAttendanceByUserIDAndDateParams{
		UserID: postParams.UserID,
		Date:   postParams.Date,
	})

	if err == nil {
		return c.Status(http.StatusConflict).JSON(fmt.Sprintf("Attendance for user %d on date %d already exists", postParams.UserID, postParams.Date))
	}

	fmt.Println("1")

	attendanceID, err := h.Store.CreateEntrance(c.Context(), postParams)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	fmt.Println("2")

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Entrance created",
		"id":      attendanceID,
	})
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

	args.Date = (args.ExitTime / 86400) * 86400

	attendance, err := h.Store.GetAttendanceByUserIDAndDate(c.Context(), args.GetAttendanceByUserIDAndDateParams)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(err.Error())
	}

	if attendance.ExitTime != 0 {
		return c.Status(http.StatusConflict).JSON(fmt.Sprintf("exit time for user %d on date %d already exists", args.UserID, args.Date))
	}
	args.UpdateExitParams.AttendanceID = attendance.AttendanceID

	err = h.Store.UpdateExit(c.Context(), args.UpdateExitParams)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Exit created",
		"id":      attendance.AttendanceID,
	})
}
*/

type AttendanceParams struct {
	UserID int64 `json:"user_id"`
	Time   int64 `json:"time"`
}

func (h *Handlers) HandleAttendance(c *fiber.Ctx) error {
	var params AttendanceParams
	if err := c.BodyParser(&params); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	// Calculate the date based on the provided time (assuming time is in Unix format)
	date := (params.Time / 86400) * 86400

	// Try to retrieve an attendance record for the user and date
	attendance, err := h.Store.GetAttendanceByUserIDAndDate(c.Context(), db.GetAttendanceByUserIDAndDateParams{
		UserID: params.UserID,
		Date:   date,
	})

	if err != nil {
		// No existing record, so we create a new attendance record for entrance
		attendanceID, err := h.Store.CreateEntrance(c.Context(), db.CreateEntranceParams{
			UserID:    params.UserID,
			EnterTime: params.Time,
			Date:      date,
		})

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(err.Error())
		}

		return c.Status(http.StatusCreated).JSON(fiber.Map{
			"message": "Entrance created",
			"id":      attendanceID,
		})
	}

	if unixToMinute(params.Time)-unixToMinute(attendance.EnterTime) < 1 {
		return c.Status(http.StatusBadRequest).JSON("repetitive record")
	}

	// If the record exists but exit_time is already set, return a conflict
	if attendance.ExitTime != 0 {
		return c.Status(http.StatusConflict).JSON(fmt.Sprintf("Exit time for user %d on date %d already exists", params.UserID, date))
	}

	// Otherwise, we update the record with the exit_time
	err = h.Store.UpdateExit(c.Context(), db.UpdateExitParams{
		AttendanceID: attendance.AttendanceID,
		ExitTime:     params.Time,
	})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Exit created",
		"id":      attendance.AttendanceID,
	})
}
