package sms

import (
	"context"
	"log"
	"time"

	"github.com/Ali-Assar/SchoolAttendanceSystem/issues/db"
	"github.com/Ali-Assar/SchoolAttendanceSystem/issues/handler"
)

func ScheduleDailyAt(store db.Querier, ctx context.Context, hour, minute int) {
	for {
		now := time.Now()

		// Calculate next 10 AM (or next day's 10 AM if it's past 10 AM today)
		next := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, now.Location())
		if now.After(next) {
			next = next.Add(24 * time.Hour)
		}

		sleepDuration := next.Sub(now)
		log.Printf("Waiting for %v until next %d:%d AM...\n", sleepDuration, hour, minute)
		time.Sleep(sleepDuration)

		date := int(time.Now().Unix())
		names, phone, err := handler.GetFormattedAbsentTeachers(store, ctx, date)
		if err != nil {
			log.Println("Error fetching absent teachers:", err)
		} else {
			log.Println("Absent teachers for the day:", names)
			log.Println("Absent teachers phone:", phone)

		}

		// After calling, sleep for 24 hours to run again the next day at 10 AM
		time.Sleep(24 * time.Hour)
	}
}
