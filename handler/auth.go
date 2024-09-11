package handler

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Ali-Assar/SchoolAttendanceSystem/issues/db"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthResponse struct {
	User  db.Admin `json:"user"`
	Token string   `json:"token"`
}

func (h *Handlers) HandleAuthenticate(c *fiber.Ctx) error {
	var params db.Admin

	if err := c.BodyParser(&params); err != nil {
		return err
	}

	admin, err := h.Store.GetAdminByUserName(c.Context(), params.UserName)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid username or password"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(params.Password))
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid username or password"})
	}

	resp := AuthResponse{
		User:  admin,
		Token: CreateTokenFromUser(admin),
	}
	return c.JSON(resp)
}

func CreateTokenFromUser(user db.Admin) string {
	expires := time.Now().Add(time.Hour * 4).Format(time.RFC3339)
	claims := jwt.MapClaims{
		"user_name": user.UserName, // Add the username here
		"expires":   expires,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println("failed to sign token:", err)
	}
	return tokenStr
}
