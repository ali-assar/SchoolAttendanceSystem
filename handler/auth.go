package handler

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Ali-Assar/SchoolAttendanceSystem/issues/db"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthResponse struct {
	User  string `json:"user"`
	Token string `json:"token"`
}

func (h *Handlers) HandleAuthenticate(c *fiber.Ctx) error {
	var params db.Admin

	if err := c.BodyParser(&params); err != nil {
		return err
	}

	admin, err := h.Store.GetAdminByUserName(c.Context(), params.UserName)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid username or password", "success": false})
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(params.Password))
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid username or password", "success": false})
	}

	resp := AuthResponse{
		User:  admin.UserName,
		Token: CreateTokenFromUser(admin),
	}
	return c.JSON(fiber.Map{"message": resp, "success": true})
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
		log.Println("failed to sign token:", err)
	}
	return tokenStr
}
