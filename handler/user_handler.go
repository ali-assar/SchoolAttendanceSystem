package handler

import (
	"net/http"
	"strconv"

	"github.com/Ali-Assar/SchoolAttendanceSystem/issues/db"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userStore db.Querier
}

func NewUserHandler(userStore db.Querier) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {
	var postParams db.CreateUserParams
	if err := c.BodyParser(&postParams); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	if err := h.userStore.CreateUser(c.Context(), postParams); err != nil {
		c.Status(http.StatusInternalServerError).JSON(err)
	}

	return c.Status(http.StatusCreated).JSON("user created")
}

func (h *UserHandler) HandleGetUserByID(c *fiber.Ctx) error {

	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(err)
	}

	fetchedUser, err := h.userStore.GetUserByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(err)
	}

	return c.Status(http.StatusOK).JSON(fetchedUser)
}
