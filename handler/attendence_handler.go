package handler

import (
	"log"
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
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error(), "success": false})
	}
	date := ExtractUnixDate(params.Time)
	log.Printf("time = %d", params.Time)

	fetchedUser, err := h.Store.GetUserByID(c.Context(), params.UserID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "success": false})
	}

	attendance, _ := h.Store.GetAttendanceByUserIDAndDate(c.Context(), db.GetAttendanceByUserIDAndDateParams{
		UserID: params.UserID,
		Date:   date,
	})

	if attendance == nil || (attendance[len(attendance)-1].ExitTime > 0) {
		if len(attendance) != 0 {
			if params.Time-attendance[len(attendance)-1].ExitTime < 60 {
				return c.Status(http.StatusBadRequest).JSON(fiber.Map{
					"message":    "repetitive record",
					"first_name": fetchedUser.FirstName,
					"last_name":  fetchedUser.LastName,
					"success":    false,
				})
			}
		}
		attendanceID, err := h.Store.CreateEntrance(c.Context(), db.CreateEntranceParams{
			UserID:    params.UserID,
			EnterTime: params.Time,
			Date:      date,
		})

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "success": false})
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message":    "Entrance created",
			"first_name": fetchedUser.FirstName,
			"last_name":  fetchedUser.LastName,
			"id":         attendanceID,
			"success":    true,
		})
	}

	if params.Time-attendance[len(attendance)-1].EnterTime < 60 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message":    "repetitive record",
			"first_name": fetchedUser.FirstName,
			"last_name":  fetchedUser.LastName,
			"success":    false,
		})
	}

	if attendance[len(attendance)-1].ExitTime == 0 {
		err = h.Store.UpdateExit(c.Context(), db.UpdateExitParams{
			AttendanceID: attendance[len(attendance)-1].AttendanceID,
			ExitTime:     params.Time,
		})

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "success": false})
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message":    "Exit created",
			"first_name": fetchedUser.FirstName,
			"last_name":  fetchedUser.LastName,
			"id":         attendance[len(attendance)-1].AttendanceID,
			"success":    true,
		})
	}

	return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "success": false})
}

func (h *Handlers) HandleGetAttendanceByTypeAndDate(c *fiber.Ctx) error {
	attendanceType := c.Params("type")

	date, err := c.ParamsInt("date")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid date format", "success": false})
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
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid attendance type. Use 'all', 'student', or 'teacher'", "success": false})
	}

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "success": false})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": attendanceRecords, "success": true})
}

func (h *Handlers) HandleGetAttendanceByTypeAndDateRange(c *fiber.Ctx) error {
	attendanceType := c.Params("type")
	startDate, err := c.ParamsInt("startDate")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid start date format", "success": false})
	}
	endDate, err := c.ParamsInt("endDate")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid end date format", "success": false})
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
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid attendance type. Use 'all', 'student', or 'teacher'", "success": false})
	}

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "success": false})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": attendanceRecords, "success": true})
}

func (h *Handlers) HandleGetAbsentTeachersByDate(c *fiber.Ctx) error {
	date, err := c.ParamsInt("date")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid date format", "success": false})
	}

	absentTeachers, err := FindAbsentTeachers(h.Store, c.Context(), date)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "success": false})
	}

	if absentTeachers == nil {
		return c.Status(http.StatusOK).JSON(fiber.Map{"message": "no absent teacher for given date", "success": true})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": absentTeachers, "success": true})
}

func (h *Handlers) HandleGetAbsentStudentsByDate(c *fiber.Ctx) error {
	date, err := c.ParamsInt("date")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid date format", "success": false})
	}

	absentStudents, err := FindAbsentStudents(h.Store, c.Context(), date)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "success": false})
	}
	if absentStudents == nil {
		return c.Status(http.StatusOK).JSON(fiber.Map{"message": "no absent student for given date", "success": false})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": absentStudents, "success": true})
}

func (h *Handlers) HandleUpdateExitByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error(), "success": false})
	}

	exit, err := c.ParamsInt("exit_time")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid date format", "success": false})
	}

	params := db.UpdateExitParams{
		ExitTime:     int64(exit),
		AttendanceID: int64(id),
	}

	err = h.Store.UpdateExit(c.Context(), params)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "success": false})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "exit time updated",
		"id":      id,
		"success": true,
	})

}

func (h *Handlers) HandleUpdateEntranceByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error(), "success": false})
	}
	exit, err := c.ParamsInt("enter_time")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid date format", "success": false})
	}

	params := db.UpdateEntranceByIDParams{
		EnterTime:    int64(exit),
		AttendanceID: int64(id),
	}

	err = h.Store.UpdateEntranceByID(c.Context(), params)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "success": false})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "enter time updated",
		"id":      id,
		"success": true,
	})

}
