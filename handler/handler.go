package handler

import (
	"fmt"
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
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	// Call to create the user with new fields
	userID, err := h.Store.CreateUser(c.Context(), postParams)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(fmt.Sprintf("user with id %d created", userID))
}

func (h *Handlers) HandleGetAllUsers(c *fiber.Ctx) error {
	// Fetch user by ID including new fields
	fetchedUsers, err := h.Store.GetAllUsers(c.Context())
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(fetchedUsers)
}

func (h *Handlers) HandleGetUserByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	// Fetch user by ID including new fields
	fetchedUser, err := h.Store.GetUserByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(err)
	}

	return c.Status(http.StatusOK).JSON(fetchedUser)
}

// Get user by phone number
func (h *Handlers) HandleGetUserByPhoneNumber(c *fiber.Ctx) error {
	phoneStr := c.Params("phone")
	phone, err := strconv.ParseInt(phoneStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	user, err := h.Store.GetUserByPhoneNumber(c.Context(), phone)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(user)
}

// Get user by first name and last name
func (h *Handlers) HandleGetUserByName(c *fiber.Ctx) error {
	var args db.GetUserByNameParams
	args.FirstName = c.Params("first_name")
	args.LastName = c.Params("last_name")

	user, err := h.Store.GetUserByName(c.Context(), args)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(user)
}

func (h *Handlers) HandleUpdateUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	var updateParams db.UpdateUserParams
	if err := c.BodyParser(&updateParams); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	updateParams.UserID = id

	// Update user with new fields
	if err := h.Store.UpdateUser(c.Context(), updateParams); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON("user updated")
}

func (h *Handlers) HandleDeleteUserByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	// Fetch user by ID including new fields
	err = h.Store.DeleteUser(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(err)
	}
	return c.Status(http.StatusOK).JSON("user deleted")
}

/*
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
		UserID: id,
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
*/
