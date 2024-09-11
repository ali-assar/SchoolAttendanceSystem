package handler

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Ali-Assar/SchoolAttendanceSystem/issues/db"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthentication(userStore db.Querier) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenStr := c.Get("X-Api-Token")
		if tokenStr == "" {
			return ErrAuthorization()
		}

		claims, err := validateToken(tokenStr)
		if err != nil {
			return ErrAuthorization()
		}

		expiresStr, ok := claims["expires"].(string)
		if !ok {
			return fmt.Errorf("invalid credentials")
		}
		expires, err := time.Parse(time.RFC3339, expiresStr)
		if err != nil || time.Now().After(expires) {
			return NewError(http.StatusUnauthorized, "token expired")
		}

		// Check if the username exists in the claims and fetch the user
		adminUsername, ok := claims["user_name"].(string)
		if !ok {
			return ErrAuthorization()
		}

		// Check if user exists and is an admin
		admin, err := userStore.GetAdminByUserName(c.Context(), adminUsername)
		if err != nil {
			return ErrAuthorization()
		}

		// Set the current authenticated user to the context
		c.Locals("user", admin)

		// Continue to the next middleware/handler
		return c.Next()
	}
}

func validateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("invalid signing method", token.Header["alg"])
			return nil, ErrAuthorization()
		}

		secret := os.Getenv("JWT_SECRET")
		return []byte(secret), nil
	})

	if err != nil {
		fmt.Println("field to parse jwt token : ", err)
		return nil, ErrAuthorization()
	}
	if !token.Valid {
		fmt.Println("invalid token")
		return nil, ErrAuthorization()
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrAuthorization()
	}
	return claims, nil
}
