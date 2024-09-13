package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Ali-Assar/SchoolAttendanceSystem/issues/db"
	"github.com/gofiber/fiber/v2"
)

// POST Handlers
type createUserInput struct {
	db.CreateUserParams
	db.CreateTeacherParams
}

func (h *Handlers) HandlePostTeacher(c *fiber.Ctx) error {
	var postParams createUserInput
	if err := c.BodyParser(&postParams); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	id, err := h.Store.CreateUser(c.Context(), postParams.CreateUserParams)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	postParams.CreateTeacherParams.UserID = id
	h.Store.CreateTeacher(c.Context(), postParams.CreateTeacherParams)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(http.StatusCreated).JSON(fmt.Sprintf("ID: %d", id))
}

type createStudentInput struct {
	db.CreateUserParams
	db.CreateStudentParams
}

func (h *Handlers) HandlePostStudent(c *fiber.Ctx) error {
	var postParams createStudentInput
	if err := c.BodyParser(&postParams); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	id, err := h.Store.CreateUser(c.Context(), postParams.CreateUserParams)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	postParams.CreateStudentParams.UserID = id
	h.Store.CreateStudent(c.Context(), postParams.CreateStudentParams)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(http.StatusCreated).JSON(fmt.Sprintf("ID: %d", id))
}

// Get Handlers

func (h *Handlers) HandleGetUserByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	user, err := h.Store.GetUserByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(http.StatusOK).JSON(user)
}

func (h *Handlers) HandleGetUserByPhoneNumber(c *fiber.Ctx) error {
	phone := c.Params("phone")
	user, err := h.Store.GetUserByPhoneNumber(c.Context(), phone)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(user)
}

func (h *Handlers) HandleGetTeacherByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	user, err := h.Store.GetTeacherByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(http.StatusOK).JSON(user)
}

func (h *Handlers) HandleGetStudentByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	user, err := h.Store.GetStudentByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(http.StatusOK).JSON(user)
}

// UPDATE handlers

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

	if err := h.Store.UpdateUser(c.Context(), updateParams); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(http.StatusCreated).JSON(fmt.Sprintf("ID: %d", id))
}

func (h *Handlers) HandleUpdateStudentAllowedTime(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	var updateParams db.UpdateStudentAllowedTimeParams
	if err := c.BodyParser(&updateParams); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	updateParams.StudentID = id

	if err := h.Store.UpdateStudentAllowedTime(c.Context(), updateParams); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(http.StatusCreated).JSON(fmt.Sprintf("ID: %d", id))
}

func (h *Handlers) HandleUpdateTeacherAllowedTime(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	var updateParams db.UpdateTeacherAllowedTimeParams
	if err := c.BodyParser(&updateParams); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	updateParams.TeacherID = id

	if err := h.Store.UpdateTeacherAllowedTime(c.Context(), updateParams); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(http.StatusCreated).JSON(fmt.Sprintf("ID: %d", id))
}

// Delete handlers

func (h *Handlers) HandleDeleteUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	err = h.Store.DeleteUser(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(err)
	}
	return c.Status(http.StatusCreated).JSON(fmt.Sprintf("ID: %d", id))
}

