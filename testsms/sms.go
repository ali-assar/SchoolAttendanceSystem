package sms

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/Ali-Assar/SchoolAttendanceSystem/issues/db"
// 	"github.com/Ali-Assar/SchoolAttendanceSystem/issues/handler"
// )

// // ScheduleDailyTask schedules a task to run every day at the specified hour and minute
// func ScheduleDailyTask(store db.Querier, hour, minute int) {
// 	now := time.Now()
// 	nextRun := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, now.Location())
// 	if nextRun.Before(now) {
// 		nextRun = nextRun.Add(24 * time.Hour)
// 	}

// 	durationUntilNextRun := nextRun.Sub(now)

// 	go func() {
// 		time.Sleep(durationUntilNextRun)
// 		runTask(store)
// 		ticker := time.NewTicker(24 * time.Hour)
// 		for range ticker.C {
// 			runTask(store)
// 		}
// 	}()
// }

// // runTask is the function that fetches absent users and sends SMS notifications
// func runTask(store db.Querier) {
// 	err := fetchAndSendAbsentUsers(store)
// 	if err != nil {
// 		log.Printf("Error fetching and sending SMS: %v", err)
// 	}
// }

// // fetchAndSendAbsentUsers fetches absent users for the current date and sends SMS notifications
// func fetchAndSendAbsentUsers(store db.Querier) error {
// 	ctx := context.Background()

// 	// Get the current date (Unix timestamp at midnight)
// 	currentDate := handler.ExtractUnixDate(time.Now().Truncate(24 * time.Hour).Unix())

// 	// Fetch absent users
// 	absentTeachers, err := store.GetAbsentTeachersByDate(ctx, currentDate)
// 	if err != nil {
// 		return fmt.Errorf("failed to fetch absent users: %w", err)
// 	}
	
	

// 	// Loop through the absent users and send SMS notifications
// 	// for _, user := range absentTeachers {
// 	// 	// err := sendSMS(user)
// 	// 	// if err != nil {
// 	// 	// 	log.Printf("Failed to send SMS to user %s %s: %v", user.FirstName, user.LastName, err)
// 	// 	// }
// 	// }

// 	return nil
// }

// // sendSMS sends an SMS notification to the user (you need to integrate with your SMS service)
// // func sendSMS(user db.GetAbsentUsersByDateRow) error {
// // 	// Placeholder for SMS sending logic
// // 	message := fmt.Sprintf("Dear %s %s, you were absent today.", user.FirstName, user.LastName)

// // 	// Call your SMS service here to send the message
// // 	fmt.Println("Sending SMS to user:", user.FirstName, user.LastName, "Message:", message)

// // 	return nil
// // }
