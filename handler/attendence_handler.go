package handler

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/Ali-Assar/SchoolAttendanceSystem/issues/db"
	"github.com/gofiber/fiber/v2"
)

// HandlePostAttendance handles creating a new attendance record
func (h *Handlers) HandlePostAttendance(c *fiber.Ctx) error {
	var postParams struct {
		UserID    int64  `json:"user_id"`
		Date      *int64 `json:"date"`       // Use *int64 to allow null values
		EntryTime *int64 `json:"entry_time"` // Use *int64 to allow null values
		ExitTime  *int64 `json:"exit_time"`  // Use *int64 to allow null values
	}

	if err := c.BodyParser(&postParams); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse request"})
	}

	// Set the date to current Unix timestamp if not provided
	if postParams.Date == nil {
		now := time.Now().UnixNano() // Current time in nanoseconds
		postParams.Date = &now
	}

	// Convert pointers to sql.NullInt64
	var entryTime sql.NullInt64
	if postParams.EntryTime != nil {
		entryTime = sql.NullInt64{Int64: *postParams.EntryTime, Valid: true}
	} else {
		entryTime = sql.NullInt64{Valid: false} // Represents null
	}

	var exitTime sql.NullInt64
	if postParams.ExitTime != nil {
		exitTime = sql.NullInt64{Int64: *postParams.ExitTime, Valid: true}
	} else {
		exitTime = sql.NullInt64{Valid: false} // Represents null
	}

	params := db.CreateAttendanceParams{
		UserID:    postParams.UserID,
		Date:      *postParams.Date,
		EntryTime: entryTime,
		ExitTime:  exitTime,
	}

	attendanceID, err := h.Store.CreateAttendance(c.Context(), params)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create attendance"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "attendance record created", "attendance_id": attendanceID})
}

// HandleGetAttendanceByUserIDAndDate retrieves attendance by user ID and date
func (h *Handlers) HandleGetAttendanceByUserIDAndDate(c *fiber.Ctx) error {
	idStr := c.Params("user_id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID format."})
	}

	dateStr := c.Params("date")
	date, err := strconv.ParseInt(dateStr, 10, 64) // Expecting date as an integer (Unix time)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid date format. Use integer timestamp."})
	}

	params := db.GetAttendanceByUserIDAndDateParams{
		UserID: id,
		Date:   date,
	}

	attendance, err := h.Store.GetAttendanceByUserIDAndDate(c.Context(), params)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(attendance)
}

// HandleGetAllUsersAttendanceByDate retrieves all users' attendance for a specific date
func (h *Handlers) HandleGetAllUsersAttendanceByDate(c *fiber.Ctx) error {
	dateStr := c.Params("date")
	date, err := strconv.ParseInt(dateStr, 10, 64) // Expecting date as an integer (Unix time)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid date format. Use integer timestamp."})
	}

	attendanceRecords, err := h.Store.GetAllUsersAttendanceByDate(c.Context(), date)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(attendanceRecords)
}

// HandleUpdateAttendanceByID updates an existing attendance record
func (h *Handlers) HandleUpdateAttendanceByID(c *fiber.Ctx) error {
	var updateParams struct {
		AttendanceID int64  `json:"attendance_id"`
		EntryTime    *int64 `json:"entry_time"` // Use *int64 to allow null values
		ExitTime     *int64 `json:"exit_time"`  // Use *int64 to allow null values
	}

	if err := c.BodyParser(&updateParams); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse request"})
	}

	// Convert pointers to sql.NullInt64
	var entryTime sql.NullInt64
	if updateParams.EntryTime != nil {
		entryTime = sql.NullInt64{Int64: *updateParams.EntryTime, Valid: true}
	} else {
		entryTime = sql.NullInt64{Valid: false} // Represents null
	}

	var exitTime sql.NullInt64
	if updateParams.ExitTime != nil {
		exitTime = sql.NullInt64{Int64: *updateParams.ExitTime, Valid: true}
	} else {
		exitTime = sql.NullInt64{Valid: false} // Represents null
	}

	params := db.UpdateAttendanceParams{
		AttendanceID: updateParams.AttendanceID,
		EntryTime:    entryTime,
		ExitTime:     exitTime,
	}

	if err := h.Store.UpdateAttendance(c.Context(), params); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "attendance record updated"})
}

// HandleDeleteAttendanceByID deletes an attendance record by its ID
func (h *Handlers) HandleDeleteAttendanceByID(c *fiber.Ctx) error {
	idStr := c.Params("attendance_id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid attendance ID format."})
	}

	if err := h.Store.DeleteAttendance(c.Context(), id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "attendance record deleted"})
}
