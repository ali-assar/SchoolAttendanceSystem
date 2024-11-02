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
			if names != "" {
				message := fmt.Sprintf("مدیر گرامی، همکاران %s .امروز غیبت داشته‌اند", names)
				fmt.Println(message)
				err := sendSMS(phone, message)
				if err != nil {
					log.Println(err)
				}
			}

		}

		// After calling, sleep for 24 hours to run again the next day at 10 AM
		time.Sleep(24 * time.Hour)
	}
}

func ScheduleDelayDailyAt(store db.Querier, ctx context.Context, hour, minute int) {
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
		names, phone, err := handler.GetFormattedTeachersDelay(store, ctx, date)
		if err != nil {
			log.Println("Error fetching absent teachers:", err)
		} else {
			log.Println("teachers delay for the day:", names)
			log.Println("teachers delay phone:", phone)
			if names != "" {
				message := fmt.Sprintf("مدیر گرامی، همکاران %s .امروز تاخییر داشته‌اند", names)
				fmt.Println(message)
				err := sendSMS(phone, message)
				if err != nil {
					log.Println(err)
				}
			}

		}

		// After calling, sleep for 24 hours to run again the next day at 10 AM
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
