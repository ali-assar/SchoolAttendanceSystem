package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

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
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error(), "success": false})
	}

	// Apply date stripping to CreatedAt if needed
	postParams.CreateUserParams.CreatedAt = ExtractUnixDate(time.Now().Unix())

	// Store user and retrieve ID
	id, err := h.Store.CreateUser(c.Context(), postParams.CreateUserParams)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "success": false})
	}
	postParams.CreateTeacherParams.UserID = id

	// Extract Unix time consistently for each weekday entry time
	dayEntryTimes := []*int64{
		&postParams.CreateTeacherParams.MondayEntryTime,
		&postParams.CreateTeacherParams.TuesdayEntryTime,
		&postParams.CreateTeacherParams.WednesdayEntryTime,
		&postParams.CreateTeacherParams.ThursdayEntryTime,
		&postParams.CreateTeacherParams.SaturdayEntryTime,
		&postParams.CreateTeacherParams.SundayEntryTime,
	}
	for _, entryTime := range dayEntryTimes {
		if entryTime != nil {
			*entryTime = ExtractUnixTime(*entryTime)
			fmt.Printf("Extracted time for entry: %d\n", *entryTime) // Logging for debugging
		}
	}

	// Store teacher
	_, err = h.Store.CreateTeacher(c.Context(), postParams.CreateTeacherParams)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "success": false})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Teacher created",
		"id":      id,
		"success": true,
	})
}

type createStudentInput struct {
	db.CreateUserParams
	db.CreateStudentParams
}

func (h *Handlers) HandlePostStudent(c *fiber.Ctx) error {
	var postParams createStudentInput
	if err := c.BodyParser(&postParams); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error(), "success": false})
	}
	postParams.CreateUserParams.CreatedAt = ExtractUnixDate(time.Now().Unix())
	id, err := h.Store.CreateUser(c.Context(), postParams.CreateUserParams)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "success": false})
	}
	postParams.CreateStudentParams.UserID = id
	_, err = h.Store.CreateStudent(c.Context(), postParams.CreateStudentParams)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "success": false})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Student created",
		"id":      id,
		"success": true,
	})
}

// Get Handlers

func (h *Handlers) HandleGetUserByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error(), "success": false})
	}

	user, err := h.Store.GetUserByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "success": false})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": user, "success": true})
}

func (h *Handlers) HandleGetUserByName(c *fiber.Ctx) error {
	var args db.GetUserByNameParams
	args.FirstName = c.Params("first_name")
	args.LastName = c.Params("last_name")

	user, err := h.Store.GetUserByName(c.Context(), args)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": err.Error(), "success": false})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": user, "success": true})
}

func (h *Handlers) HandleGetTeacherByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": err.Error(), "success": false})
	}

	user, err := h.Store.GetTeacherByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": err.Error(), "success": false})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": user, "success": true})
}

func (h *Handlers) HandleGetTeachers(c *fiber.Ctx) error {

	user, err := h.Store.GetTeachers(c.Context())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error(), "success": false})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": user, "success": true})
}

func (h *Handlers) HandleGetStudentByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error(), "success": false})
	}

	user, err := h.Store.GetStudentByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error(), "success": false})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": user, "success": true})
}

func (h *Handlers) HandleGetStudents(c *fiber.Ctx) error {

	user, err := h.Store.GetStudents(c.Context())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error(), "success": false})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": user, "success": true})
}

func (h *Handlers) HandleGetUsersWithFalseBiometric(c *fiber.Ctx) error {

	users, err := h.Store.GetUsersWithFalseBiometric(c.Context())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error(), "success": false})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": users, "success": true})
}

func (h *Handlers) HandleGetUsersWithTrueBiometric(c *fiber.Ctx) error {

	users, err := h.Store.GetUsersWithTrueBiometric(c.Context())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error(), "success": false})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": users, "success": true})
}

// UPDATE handlers
func (h *Handlers) HandleUpdateUserBiometric(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error(), "success": false})
	}

	_, err = h.Store.GetUserByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "User not found", "success": false})
	}
	var updateParams db.UpdateUserBiometricParams
	if err := c.BodyParser(&updateParams); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error(), "success": false})
	}

	updateParams.UserID = id
	updateParams.IsBiometricActive = true

	if err := h.Store.UpdateUserBiometric(c.Context(), updateParams); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error(), "success": false})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "user biometric updated",
		"id":      id,
		"success": true,
	})
}

// UPDATE handlers
func (h *Handlers) HandleUpdateUserBiometricToFalse(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error(), "success": false})
	}

	_, err = h.Store.GetUserByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "User not found", "success": false})
	}
	var updateParams db.UpdateUserBiometricParams

	updateParams.UserID = id
	updateParams.IsBiometricActive = false

	if err := h.Store.UpdateUserBiometric(c.Context(), updateParams); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error(), "success": false})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "user biometric updated",
		"id":      id,
		"success": true,
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
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error(), "success": false})
	}

	fetchedStudent, err := h.Store.GetStudentByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Teacher not found", "success": false})
	}

	fetchedUser, err := h.Store.GetUserByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "User not found", "success": false})
	}

	var updateParams updateStudentParams
	if err := c.BodyParser(&updateParams); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error(), "success": false})
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
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error(), "success": false})
	}

	if err := h.Store.UpdateStudentAllowedTime(c.Context(), updateParams.UpdateStudentAllowedTimeParams); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error(), "success": false})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "student updated",
		"id":      id,
		"success": true,
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
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error(), "success": false})
	}

	fetchedTeacher, err := h.Store.GetTeacherByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Teacher not found", "success": false})
	}

	fetchedUser, err := h.Store.GetUserByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "User not found", "success": false})
	}

	var updateParams updateTeacherParams
	if err := c.BodyParser(&updateParams); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error(), "success": false})
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

	// Apply ExtractUnixTime to all entry times to ensure consistency
	updateParams.UpdateTeacherAllowedTimeParams.SundayEntryTime = ExtractUnixTimeOrFetched(
		updateParams.UpdateTeacherAllowedTimeParams.SundayEntryTime, fetchedTeacher.SundayEntryTime)
	updateParams.UpdateTeacherAllowedTimeParams.MondayEntryTime = ExtractUnixTimeOrFetched(
		updateParams.UpdateTeacherAllowedTimeParams.MondayEntryTime, fetchedTeacher.MondayEntryTime)
	updateParams.UpdateTeacherAllowedTimeParams.TuesdayEntryTime = ExtractUnixTimeOrFetched(
		updateParams.UpdateTeacherAllowedTimeParams.TuesdayEntryTime, fetchedTeacher.TuesdayEntryTime)
	updateParams.UpdateTeacherAllowedTimeParams.WednesdayEntryTime = ExtractUnixTimeOrFetched(
		updateParams.UpdateTeacherAllowedTimeParams.WednesdayEntryTime, fetchedTeacher.WednesdayEntryTime)
	updateParams.UpdateTeacherAllowedTimeParams.ThursdayEntryTime = ExtractUnixTimeOrFetched(
		updateParams.UpdateTeacherAllowedTimeParams.ThursdayEntryTime, fetchedTeacher.ThursdayEntryTime)
	updateParams.UpdateTeacherAllowedTimeParams.FridayEntryTime = ExtractUnixTimeOrFetched(
		updateParams.UpdateTeacherAllowedTimeParams.FridayEntryTime, fetchedTeacher.FridayEntryTime)
	updateParams.UpdateTeacherAllowedTimeParams.SaturdayEntryTime = ExtractUnixTimeOrFetched(
		updateParams.UpdateTeacherAllowedTimeParams.SaturdayEntryTime, fetchedTeacher.SaturdayEntryTime)

	updateParams.UpdateUserDetailsParams.UserID = id
	updateParams.UpdateTeacherAllowedTimeParams.UserID = id

	if err := h.Store.UpdateUserDetails(c.Context(), updateParams.UpdateUserDetailsParams); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error(), "success": false})
	}

	if err := h.Store.UpdateTeacherAllowedTime(c.Context(), updateParams.UpdateTeacherAllowedTimeParams); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error(), "success": false})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "teacher updated",
		"id":      id,
		"success": true,
	})
}

// Delete handlers

func (h *Handlers) HandleDeleteUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error(), "success": false})
	}

	err = h.Store.DeleteUser(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error(), "success": false})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "user deleted",
		"id":      id,
		"success": true,
	})
}

func (h *Handlers) HandleGetUserByJWT(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(db.Admin)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized", "success": false,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user_name": user.UserName,
		"success":   true,
	})
}
