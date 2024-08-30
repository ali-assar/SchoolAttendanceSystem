package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/Ali-Assar/SchoolAttendanceSystem/issues/db"
	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
	Store db.Querier
}

func NewHandlers(store db.Querier) *Handlers {
	return &Handlers{
		Store: store,
	}
}

// --- User Handlers ---

func (h *Handlers) HandlePostUser(c *fiber.Ctx) error {
	var postParams db.CreateUserParams
	if err := c.BodyParser(&postParams); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	if err := h.Store.CreateUser(c.Context(), postParams); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err)
	}

	return c.Status(http.StatusCreated).JSON("user created")
}

func (h *Handlers) HandleGetUserByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	fetchedUser, err := h.Store.GetUserByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(err)
	}

	return c.Status(http.StatusOK).JSON(fetchedUser)
}

/*
	func (h *Handlers) HandleGetAllUsers(c *fiber.Ctx) error {
		users, err := h.Store.GetAllUsers(c.Context())
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(err)
		}
		return c.Status(http.StatusOK).JSON(users)
	}
*/
func (h *Handlers) HandleUpdateUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	var updateParams db.UpdateUserParams
	if err := c.BodyParser(&updateParams); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	updateParams.UserID = id

	if err := h.Store.UpdateUser(c.Context(), updateParams); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err)
	}

	return c.Status(http.StatusOK).JSON("user updated")
}

func (h *Handlers) HandleDeleteUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	if err := h.Store.DeleteUser(c.Context(), id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err)
	}

	return c.Status(http.StatusOK).JSON("user deleted")
}

// --- Attendance Handlers ---

func (h *Handlers) HandlePostAttendance(c *fiber.Ctx) error {
	var postParams db.CreateAttendanceParams
	if err := c.BodyParser(&postParams); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	if err := h.Store.CreateAttendance(c.Context(), postParams); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err)
	}

	return c.Status(http.StatusCreated).JSON("attendance record created")
}

func (h *Handlers) HandleGetAttendanceByUserIDAndDate(c *fiber.Ctx) error {
	idStr := c.Params("user_id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	date := c.Params("date")

	params := db.GetAttendanceByUserIDAndDateParams{
		UserID: sql.NullInt64{Int64: id, Valid: true},
		Date:   date,
	}

	attendance, err := h.Store.GetAttendanceByUserIDAndDate(c.Context(), params)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(err)
	}

	return c.Status(http.StatusOK).JSON(attendance)
}

func (h *Handlers) HandleGetAllUsersAttendanceByDate(c *fiber.Ctx) error {
	date := c.Params("date")

	attendanceRecords, err := h.Store.GetAllUsersAttendanceByDate(c.Context(), date)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err)
	}

	return c.Status(http.StatusOK).JSON(attendanceRecords)
}

func (h *Handlers) HandleUpdateAttendance(c *fiber.Ctx) error {
	var updateParams db.UpdateAttendanceParams
	if err := c.BodyParser(&updateParams); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	if err := h.Store.UpdateAttendance(c.Context(), updateParams); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err)
	}

	return c.Status(http.StatusOK).JSON("attendance record updated")
}

func (h *Handlers) HandleDeleteAttendance(c *fiber.Ctx) error {
	idStr := c.Params("attendance_id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	if err := h.Store.DeleteAttendance(c.Context(), id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err)
	}

	return c.Status(http.StatusOK).JSON("attendance record deleted")
}
