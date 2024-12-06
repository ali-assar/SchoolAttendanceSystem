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

func ScheduleDailyNotifications(store db.Querier, ctx context.Context, hour, minute int) {
	for {
		now := time.Now()

		// Set next execution time based on current OS time
		next := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, 0, 0, now.Location())
		if now.After(next) {
			next = next.Add(24 * time.Hour)
		}

		log.Printf("Current system time: %v", time.Now())
		log.Printf("Scheduled time for next notifications: %v", next)

		sleepDuration := next.Sub(now)
		log.Printf("Sleeping for %v until next scheduled time...\n", sleepDuration)
		time.Sleep(sleepDuration)

		date := int(time.Now().Unix())

		delayedNames, delayedPhone, err := handler.GetFormattedTeachersDelay(store, ctx, date)
		if err != nil {
			log.Println("Error fetching delayed teachers:", err)
		} else if delayedNames != "" {
			delayMessage := fmt.Sprintf("مدیر گرامی، همکاران %s امروز تاخییر داشته‌اند.", delayedNames)
			fmt.Println(delayMessage) // Log the delay message
			err := sendSMS(delayedPhone, delayMessage)
			if err != nil {
				log.Println("SMS send error (delays):", err)
			}
		}

		time.Sleep(1 * time.Minute)

		absentNames, absentPhone, err := handler.GetFormattedAbsentTeachers(store, ctx, date)
		if err != nil {
			log.Println("Error fetching absent teachers:", err)
		} else if absentNames != "" {
			absentMessage := fmt.Sprintf("مدیر گرامی، همکاران %s امروز غایب بوده‌اند.", absentNames)
			fmt.Println(absentMessage) // Log the absentee message
			err := sendSMS(absentPhone, absentMessage)
			if err != nil {
				log.Println("SMS send error (absents):", err)
			}
		}

		// Sleep for 24 hours before the next scheduled time
		// time.Sleep(2 * time.Minute)
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
