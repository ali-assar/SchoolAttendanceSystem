package handler

import (
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type AuthHandler struct {
	Store *session.Store
}

func (h *AuthHandler) IsAuthenticated(c *fiber.Ctx) error {
	sess, err := h.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not retrieve session")
	}

	if auth, ok := sess.Get("authenticated").(bool); !ok || !auth {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized access, please login")
	}

	return c.Next()
}

// Login route for admin authentication
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var login struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&login); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid login payload")
	}

	if login.Username == "admin" && login.Password == "admin" {
		sess, err := h.Store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to create session")
		}

		sess.Set("authenticated", true)
		sess.SetExpiry(2 * time.Hour) 
		if err := sess.Save(); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to save session")
		}

		return c.SendString("Login successful")
	}

	return c.Status(fiber.StatusUnauthorized).SendString("Invalid credentials")
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	sess, err := h.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not retrieve session")
	}

	sess.Destroy()
	return c.SendString("Logged out successfully")
}
