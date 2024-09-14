package handler

import (
	"fmt"
	"net/http"

	"github.com/Ali-Assar/SchoolAttendanceSystem/issues/db"
	"github.com/gofiber/fiber/v2"
)

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
	date := ExtractUnixDate(params.Time)

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

	if UnixToMinute(params.Time)-UnixToMinute(attendance.EnterTime) < 1 {
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

func (h *Handlers) GetAttendanceByDate(c *fiber.Ctx) error {
	date, err := c.ParamsInt("date")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Invalid date format")
	}

	attendanceRecords, err := h.Store.GetAttendanceByDate(c.Context(), int64(date))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(attendanceRecords)
}

func (h *Handlers) GetAttendanceBetweenDates(c *fiber.Ctx) error {
	startDate, err := c.ParamsInt("startDate")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Invalid start date format")
	}

	endDate, err := c.ParamsInt("endDate")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Invalid end date format")
	}

	attendanceRecords, err := h.Store.GetAttendanceBetweenDates(c.Context(), db.GetAttendanceBetweenDatesParams{
		FromDate: int64(startDate),
		ToDate:   int64(endDate),
	})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(attendanceRecords)
}

// Handler to fetch absent users on a particular date (int)
func (h *Handlers) GetAbsentUsersByDate(c *fiber.Ctx) error {
	date, err := c.ParamsInt("date")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Invalid date format")
	}

	absentUsers, err := h.Store.GetAbsentUsersByDate(c.Context(), int64(date))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(absentUsers)
}
