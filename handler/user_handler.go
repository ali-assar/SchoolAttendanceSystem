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
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}
	id, err := h.Store.CreateUser(c.Context(), postParams.CreateUserParams)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	postParams.CreateTeacherParams.UserID = id
	h.Store.CreateTeacher(c.Context(), postParams.CreateTeacherParams)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err})
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
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}
	id, err := h.Store.CreateUser(c.Context(), postParams.CreateUserParams)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	postParams.CreateStudentParams.UserID = id
	h.Store.CreateStudent(c.Context(), postParams.CreateStudentParams)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err})
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
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	user, err := h.Store.GetUserByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return c.Status(http.StatusOK).JSON(user)
}

func (h *Handlers) HandleGetUserByName(c *fiber.Ctx) error {
	var args db.GetUserByNameParams
	args.FirstName = c.Params("first_name")
	args.LastName = c.Params("last_name")

	user, err := h.Store.GetUserByName(c.Context(), args)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err})
	}

	return c.Status(http.StatusOK).JSON(user)
}

func (h *Handlers) HandleGetTeacherByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	user, err := h.Store.GetTeacherByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return c.Status(http.StatusOK).JSON(user)
}

func (h *Handlers) HandleGetTeachers(c *fiber.Ctx) error {

	user, err := h.Store.GetTeachers(c.Context())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return c.Status(http.StatusOK).JSON(user)
}

func (h *Handlers) HandleGetStudentByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	user, err := h.Store.GetStudentByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return c.Status(http.StatusOK).JSON(user)
}

func (h *Handlers) HandleGetStudents(c *fiber.Ctx) error {

	user, err := h.Store.GetStudents(c.Context())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return c.Status(http.StatusOK).JSON(user)
}

func (h *Handlers) HandleGetUsersWithFalseBiometric(c *fiber.Ctx) error {

	users, err := h.Store.GetUsersWithFalseBiometric(c.Context())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return c.Status(http.StatusOK).JSON(users)
}

func (h *Handlers) HandleGetUsersWithTrueBiometric(c *fiber.Ctx) error {

	users, err := h.Store.GetUsersWithTrueBiometric(c.Context())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return c.Status(http.StatusOK).JSON(users)
}

// UPDATE handlers
func (h *Handlers) HandleUpdateUserBiometric(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	_, err = h.Store.GetUserByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}
	var updateParams db.UpdateUserBiometricParams
	if err := c.BodyParser(&updateParams); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	updateParams.UserID = id
	updateParams.IsBiometricActive = true

	if err := h.Store.UpdateUserBiometric(c.Context(), updateParams); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "user biometric updated",
		"id":      id,
	})
}

type updateStudentParams struct {
	db.UpdateUserDetailsParams
	db.UpdateStudentAllowedTimeParams
}

func (h *Handlers) HandleUpdateStudent(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	fetchedStudent, err := h.Store.GetStudentByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Teacher not found"})
	}

	fetchedUser, err := h.Store.GetUserByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	var updateParams updateStudentParams
	if err := c.BodyParser(&updateParams); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	// Fill missing user details with the fetched user data
	if updateParams.UpdateUserDetailsParams.FirstName == "" {
		updateParams.UpdateUserDetailsParams.FirstName = fetchedUser.FirstName
	}
	if updateParams.UpdateUserDetailsParams.LastName == "" {
		updateParams.UpdateUserDetailsParams.LastName = fetchedUser.LastName
	}
	if updateParams.UpdateUserDetailsParams.PhoneNumber == "" {
		updateParams.UpdateUserDetailsParams.PhoneNumber = fetchedUser.PhoneNumber
	}
	if updateParams.UpdateUserDetailsParams.ImagePath == "" {
		updateParams.UpdateUserDetailsParams.ImagePath = fetchedUser.ImagePath
	}

	// Fill missing teacher allowed time params with fetched teacher data
	if updateParams.UpdateStudentAllowedTimeParams.RequiredEntryTime == 0 {
		updateParams.UpdateStudentAllowedTimeParams.RequiredEntryTime = fetchedStudent.RequiredEntryTime
	}

	updateParams.UpdateUserDetailsParams.UserID = id
	updateParams.UpdateStudentAllowedTimeParams.UserID = id

	if err := h.Store.UpdateUserDetails(c.Context(), updateParams.UpdateUserDetailsParams); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	if err := h.Store.UpdateStudentAllowedTime(c.Context(), updateParams.UpdateStudentAllowedTimeParams); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "student updated",
		"id":      id,
	})
}

type updateTeacherParams struct {
	db.UpdateUserDetailsParams
	db.UpdateTeacherAllowedTimeParams
}

func (h *Handlers) HandleUpdateTeacher(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	fetchedTeacher, err := h.Store.GetTeacherByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Teacher not found"})
	}

	fetchedUser, err := h.Store.GetUserByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	var updateParams updateTeacherParams
	if err := c.BodyParser(&updateParams); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	// Fill missing user details with the fetched user data
	if updateParams.UpdateUserDetailsParams.FirstName == "" {
		updateParams.UpdateUserDetailsParams.FirstName = fetchedUser.FirstName
	}
	if updateParams.UpdateUserDetailsParams.LastName == "" {
		updateParams.UpdateUserDetailsParams.LastName = fetchedUser.LastName
	}
	if updateParams.UpdateUserDetailsParams.PhoneNumber == "" {
		updateParams.UpdateUserDetailsParams.PhoneNumber = fetchedUser.PhoneNumber
	}
	if updateParams.UpdateUserDetailsParams.ImagePath == "" {
		updateParams.UpdateUserDetailsParams.ImagePath = fetchedUser.ImagePath
	}

	// Fill missing teacher allowed time params with fetched teacher data
	if updateParams.UpdateTeacherAllowedTimeParams.SundayEntryTime == 0 {
		updateParams.UpdateTeacherAllowedTimeParams.SundayEntryTime = fetchedTeacher.SundayEntryTime
	}
	if updateParams.UpdateTeacherAllowedTimeParams.MondayEntryTime == 0 {
		updateParams.UpdateTeacherAllowedTimeParams.MondayEntryTime = fetchedTeacher.MondayEntryTime
	}
	if updateParams.UpdateTeacherAllowedTimeParams.TuesdayEntryTime == 0 {
		updateParams.UpdateTeacherAllowedTimeParams.TuesdayEntryTime = fetchedTeacher.TuesdayEntryTime
	}
	if updateParams.UpdateTeacherAllowedTimeParams.WednesdayEntryTime == 0 {
		updateParams.UpdateTeacherAllowedTimeParams.WednesdayEntryTime = fetchedTeacher.WednesdayEntryTime
	}
	if updateParams.UpdateTeacherAllowedTimeParams.ThursdayEntryTime == 0 {
		updateParams.UpdateTeacherAllowedTimeParams.ThursdayEntryTime = fetchedTeacher.ThursdayEntryTime
	}
	if updateParams.UpdateTeacherAllowedTimeParams.FridayEntryTime == 0 {
		updateParams.UpdateTeacherAllowedTimeParams.FridayEntryTime = fetchedTeacher.FridayEntryTime
	}
	if updateParams.UpdateTeacherAllowedTimeParams.SaturdayEntryTime == 0 {
		updateParams.UpdateTeacherAllowedTimeParams.SaturdayEntryTime = fetchedTeacher.SaturdayEntryTime
	}

	updateParams.UpdateUserDetailsParams.UserID = id
	updateParams.UpdateTeacherAllowedTimeParams.UserID = id

	if err := h.Store.UpdateUserDetails(c.Context(), updateParams.UpdateUserDetailsParams); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	if err := h.Store.UpdateTeacherAllowedTime(c.Context(), updateParams.UpdateTeacherAllowedTimeParams); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "teacher updated",
		"id":      id,
	})
}

// Delete handlers

func (h *Handlers) HandleDeleteUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	err = h.Store.DeleteUser(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err})
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
