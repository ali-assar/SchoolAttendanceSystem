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
	date := ExtractUnixDate(params.Time)

	attendance, err := h.Store.GetAttendanceByUserIDAndDate(c.Context(), db.GetAttendanceByUserIDAndDateParams{
		UserID: params.UserID,
		Date:   date,
	})

	if err != nil {
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

func (h *Handlers) GetAbsentTeachersByDate(c *fiber.Ctx) error {
	date, err := c.ParamsInt("date")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Invalid date format")
	}

	absentTeachers, err := FindAbsentTeachers(h.Store, c.Context(), date)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	if absentTeachers == nil {
		return c.Status(http.StatusOK).JSON(fiber.Map{"message": "no absent teacher for given date"})
	}

	return c.Status(http.StatusOK).JSON(absentTeachers)
}

func (h *Handlers) GetAbsentStudentsByDate(c *fiber.Ctx) error {
	date, err := c.ParamsInt("date")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Invalid date format")
	}

	absentStudents, err := FindAbsentStudents(h.Store, c.Context(), date)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	if absentStudents == nil {
		return c.Status(http.StatusOK).JSON(fiber.Map{"message": "no absent student for given date"})
	}
	return c.Status(http.StatusOK).JSON(absentStudents)
}

/*
func (h *Handlers) GetAbsentUsersByDate(c *fiber.Ctx) error {
	attendanceType := c.Params("type")
	date, err := c.ParamsInt("date")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON("Invalid start date format")
	}
	var absentRecords interface{}

	switch attendanceType {
	case "teacher":
		absentRecords, err := FindAbsentTeachers(h.Store, c.Context(), date)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(err.Error())
		}
		if absentRecords == nil {
			return c.Status(http.StatusOK).JSON(fiber.Map{"message": "no absent teachers for given date"})
		}
	case "student":
		absentRecords, err := FindAbsentStudents(h.Store, c.Context(), date)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(err.Error())
		}
		if absentRecords == nil {
			return c.Status(http.StatusOK).JSON(fiber.Map{"message": "no absent student for given date"})
		}
	default:
		return c.Status(http.StatusBadRequest).JSON("Invalid attendance type. Use 'student', or 'teacher'.")
	}
	return c.Status(http.StatusOK).JSON(absentRecords)

}
*/
