package handler

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) HandleGetCurrentTime(c *fiber.Ctx) error {
	nowUTC := time.Now().UTC()
	nowLocal := time.Now()

	_, offset := nowLocal.Zone()

	nowLocalUnix := nowLocal.Unix() + int64(offset)
	nowUTCUnix := nowUTC.Unix()

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"utc_time":        nowUTC,
		"local_time":      nowLocal,
		"utc_time_unix":   nowUTCUnix,
		"local_time_unix": nowLocalUnix,
		"success":         true,
	})
}
