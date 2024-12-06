package sms

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/Ali-Assar/SchoolAttendanceSystem/issues/db"
	"github.com/Ali-Assar/SchoolAttendanceSystem/issues/handler"
)

func ScheduleDailyAt(store db.Querier, ctx context.Context, hour, minute int) {
	for {
		now := time.Now()

		// Set next execution time based on current OS time
		next := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, now.Location())
		if now.After(next) {
			next = next.Add(24 * time.Hour)
		}

		// Check and log current system time each day
		log.Printf("Current system time: %v", time.Now())
		log.Printf("Scheduled time for next SMS: %v", next)

		sleepDuration := next.Sub(now)
		log.Printf("Sleeping for %v until next scheduled time...\n", sleepDuration)
		time.Sleep(sleepDuration)

		// Perform the scheduled task
		date := int(time.Now().Unix())
		names, phone, err := handler.GetFormattedAbsentTeachers(store, ctx, date)
		if err != nil {
			log.Println("Error fetching absent teachers:", err)
		} else if names != "" {
			message := fmt.Sprintf("مدیر گرامی، همکاران %s .امروز غیبت داشته‌اند", names)
			fmt.Println(message)
			err := sendSMS(phone, message)
			if err != nil {
				log.Println("SMS send error:", err)
			}
		}

		// Ensure it runs again in the next 24 hours regardless of time sync issues
		time.Sleep(2 * time.Minute)
	}
}

func ScheduleDelayDailyAt(store db.Querier, ctx context.Context, hour, minute int) {
	for {
		now := time.Now()

		// Set next execution time based on current OS time
		next := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, now.Location())
		if now.After(next) {
			next = next.Add(24 * time.Hour)
		}

		log.Printf("Current system time: %v", time.Now())
		log.Printf("Scheduled time for next SMS: %v", next)

		sleepDuration := next.Sub(now)
		log.Printf("Sleeping for %v until next scheduled time...\n", sleepDuration)
		time.Sleep(sleepDuration)

		// Perform the scheduled task
		date := int(time.Now().Unix())
		names, phone, err := handler.GetFormattedTeachersDelay(store, ctx, date)
		if err != nil {
			log.Println("Error fetching delayed teachers:", err)
		} else if names != "" {
			message := fmt.Sprintf("مدیر گرامی، همکاران %s امروز تاخییر داشته‌اند.", names)
			fmt.Println(message)
			err := sendSMS(phone, message)
			if err != nil {
				log.Println("SMS send error:", err)
			}
		}

		// Sleep for 24 hours before the next scheduled time
		time.Sleep(24 * time.Hour)
	}
}

func sendSMS(phone string, message string) error {
	// Replace with your username, password, and sender number
	username := "mt.09351851182"
	password := "Eram@12321"
	from := "10009611"

	encodedMessage := url.QueryEscape(message)

	// Build the URL for the SMS request
	url := fmt.Sprintf("https://media.sms24.ir/SMSInOutBox/SendSms?username=%s&password=%s&from=%s&to=%s&text=%s",
		username, password, from, phone, encodedMessage)

	fmt.Println(url)
	// Make the HTTP request
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read the response body for logging purposes
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Log the response
	log.Printf("SMS API response: %s", string(body))
	return nil
}
