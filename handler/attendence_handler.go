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

func (h *Handlers) GetAttendanceByTypeAndDate(c *fiber.Ctx) error {
	attendanceType := c.Params("type")
	
	date, err := c.ParamsInt("date")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Invalid date format")
	}

	date = int(ExtractUnixDate(int64(date)))

	var attendanceRecords interface{}

	switch attendanceType {
	case "all":
		attendanceRecords, err = h.Store.GetAttendanceByDate(c.Context(), int64(date))
	case "student":
		attendanceRecords, err = h.Store.GetStudentAttendanceByDate(c.Context(), int64(date))
	case "teacher":
		attendanceRecords, err = h.Store.GetTeacherAttendanceByDate(c.Context(), int64(date))
	default:
		return c.Status(http.StatusBadRequest).JSON("Invalid attendance type. Use 'all', 'student', or 'teacher'.")
	}

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(attendanceRecords)
}



func (h *Handlers) GetAttendanceByTypeAndDateRange(c *fiber.Ctx) error {
	attendanceType := c.Params("type")
	startDate, err := c.ParamsInt("startDate")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Invalid start date format")
	}
	endDate, err := c.ParamsInt("endDate")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Invalid end date format")
	}
	var attendanceRecords interface{}

	switch attendanceType {
	case "all":
		attendanceRecords, err = h.Store.GetAttendanceBetweenDates(c.Context(), db.GetAttendanceBetweenDatesParams{
			FromDate: int64(startDate),
			ToDate:   int64(endDate),
		})
	case "student":
		attendanceRecords, err = h.Store.GetStudentAttendanceBetweenDates(c.Context(), db.GetStudentAttendanceBetweenDatesParams{
			FromDate: int64(startDate),
			ToDate:   int64(endDate),
		})
	case "teacher":
		attendanceRecords, err = h.Store.GetTeacherAttendanceBetweenDates(c.Context(), db.GetTeacherAttendanceBetweenDatesParams{
			FromDate: int64(startDate),
			ToDate:   int64(endDate),
		})
	default:
		return c.Status(http.StatusBadRequest).JSON("Invalid attendance type. Use 'all', 'student', or 'teacher'.")
	}

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(attendanceRecords)
}





func (h *Handlers) GetAbsentUsersByDate(c *fiber.Ctx) error {
	date, err := c.ParamsInt("date")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Invalid date format")
	}

	date = int(ExtractUnixDate(int64(date)))

	absentUsers, err := h.Store.GetAbsentUsersByDate(c.Context(), int64(date))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(absentUsers)
}



func (h *Handlers) GetAbsentTeachersByDate(c *fiber.Ctx) error {
	date, err := c.ParamsInt("date")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Invalid date format")
	}
	date = int(ExtractUnixDate(int64(date)))

	dayOfWeek := ((date / 86400) + 4) % 7

	absentUsers, err := h.Store.GetAbsentTeachersByDate(c.Context(), int64(date))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	var absentOnDay []db.GetAbsentTeachersByDateRow
	for _, user := range absentUsers {
		switch dayOfWeek {
		case 1:
			if user.MondayEntryTime != 0 {
				absentOnDay = append(absentOnDay, user)
			}
		case 2:
			if user.TuesdayEntryTime != 0 {
				absentOnDay = append(absentOnDay, user)
			}
		case 3:
			if user.WednesdayEntryTime != 0 {
				absentOnDay = append(absentOnDay, user)
			}
		case 4:
			if user.ThursdayEntryTime != 0 {
				absentOnDay = append(absentOnDay, user)
			}
		case 5:
			if user.FridayEntryTime != 0 {
				absentOnDay = append(absentOnDay, user)
			}
		case 6:
			if user.SaturdayEntryTime != 0 {
				absentOnDay = append(absentOnDay, user)
			}
		case 0:
			if user.SundayEntryTime != 0 {
				absentOnDay = append(absentOnDay, user)
			}
		}
	}

	return c.Status(http.StatusOK).JSON(absentOnDay)
}


