package handler

import (
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

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Teacher created",
		"id":      id,
	})
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
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Student created",
		"id":      id,
	})
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

func (h *Handlers) HandleGetUsersWithFalseBiometric(c *fiber.Ctx) error {

	users, err := h.Store.GetUsersWithFalseBiometric(c.Context())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(http.StatusOK).JSON(users)
}

func (h *Handlers) HandleGetUsersWithTrueBiometric(c *fiber.Ctx) error {

	users, err := h.Store.GetUsersWithTrueBiometric(c.Context())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(http.StatusOK).JSON(users)
}

// UPDATE handlers

func (h *Handlers) HandleUpdateUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	_, err = h.Store.GetUserByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON("User not found")
	}
	var updateParams db.UpdateUserDetailsParams
	if err := c.BodyParser(&updateParams); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	updateParams.UserID = id

	if err := h.Store.UpdateUserDetails(c.Context(), updateParams); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "user updated",
		"id":      id,
	})
}

func (h *Handlers) HandleUpdateUserBiometric(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	_, err = h.Store.GetUserByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON("User not found")
	}
	var updateParams db.UpdateUserBiometricParams
	if err := c.BodyParser(&updateParams); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	updateParams.UserID = id
	updateParams.IsBiometricActive = true

	if err := h.Store.UpdateUserBiometric(c.Context(), updateParams); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "user biometric updated",
		"id":      id,
	})
}

func (h *Handlers) HandleUpdateStudentAllowedTime(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	_, err = h.Store.GetStudentByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON("Student not found")
	}

	var updateParams db.UpdateStudentAllowedTimeParams
	if err := c.BodyParser(&updateParams); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	updateParams.UserID = id

	if err := h.Store.UpdateStudentAllowedTime(c.Context(), updateParams); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "student allowed time updated",
		"id":      id,
	})
}

func (h *Handlers) HandleUpdateTeacherAllowedTime(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	_, err = h.Store.GetTeacherByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON("Teacher not found")
	}

	var updateParams db.UpdateTeacherAllowedTimeParams
	if err := c.BodyParser(&updateParams); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	updateParams.UserID = id

	if err := h.Store.UpdateTeacherAllowedTime(c.Context(), updateParams); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "teacher allowed time updated",
		"id":      id,
	})
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
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "user deleted",
		"id":      id,
	})
}

func (h *Handlers) HandleGetUserByJWT(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(db.Admin)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user_name": user.UserName,
	})
}

