package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/Ali-Assar/SchoolAttendanceSystem/issues/db"
	"github.com/Ali-Assar/SchoolAttendanceSystem/issues/handler"
	"github.com/joho/godotenv"
)

const (
	numberOfDays  = 30
	numberOfUsers = 100
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	defer database.Close()

	db.TearDown(database)

	err = db.CreateTables(database)
	if err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}

	store := db.New(database)
	ctx := context.Background()

	err = seedDB(ctx, store)
	if err != nil {
		log.Fatalf("Failed to insert to the tables: %v", err)
	}
}

func seedDB(ctx context.Context, store db.Querier) error {
	var totalUsers = numberOfUsers
	var teacherCount = totalUsers / 10 // 10% of users will be teachers
	var studentCount = totalUsers - teacherCount

	rand.Seed(time.Now().UnixNano())

	// Seed teachers
	for i := 1; i < teacherCount+1; i++ {
		// Create teacher user
		userParams := db.CreateUserParams{
			FirstName:   fmt.Sprintf("%d teacher ", i),
			LastName:    fmt.Sprintf("teacher lastname %d", i),
			PhoneNumber: fmt.Sprintf("09%d", rand.Intn(1000000000)),
		}

		userID, err := store.CreateUser(ctx, userParams)
		if err != nil {
			return err
		}

		// Insert teacher-specific data
		teacherParams := db.CreateTeacherParams{
			UserID:             userID,
			SundayEntryTime:    randomEntryTime(),
			SaturdayEntryTime:  randomEntryTime(),
			MondayEntryTime:    randomEntryTime(),
			TuesdayEntryTime:   randomEntryTime(),
			WednesdayEntryTime: randomEntryTime(),
			ThursdayEntryTime:  randomEntryTime(),
			FridayEntryTime:    randomEntryTime(),
		}

		_, err = store.CreateTeacher(ctx, teacherParams)
		if err != nil {
			return err
		}

		// Seed entrance and exit times for 30 days (same as before)
		err = seedAttendance(ctx, store, userID)
		if err != nil {
			return err
		}
	}

	// Seed students
	for i := 1; i < studentCount+1; i++ {
		// Create student user
		userParams := db.CreateUserParams{
			FirstName:   fmt.Sprintf("%d student ", i),
			LastName:    fmt.Sprintf("student lastname %d", i),
			PhoneNumber: fmt.Sprintf("09%d", rand.Intn(1000000000)),
		}

		userID, err := store.CreateUser(ctx, userParams)
		if err != nil {
			return err
		}

		// Insert student-specific data
		studentParams := db.CreateStudentParams{
			UserID:            userID,
			RequiredEntryTime: 28800,
		}

		_, err = store.CreateStudent(ctx, studentParams)
		if err != nil {
			return err
		}

		// Seed entrance and exit times for 30 days (same as before)
		err = seedAttendance(ctx, store, userID)
		if err != nil {
			return err
		}
	}

	return nil
}

func seedAttendance(ctx context.Context, store db.Querier, userID int64) error {
	for day := 0; day < numberOfDays; day++ {
		// Random date for each day in the range
		currentDate := time.Now().AddDate(0, 0, -day)

		date := handler.ExtractUnixDate(currentDate.Unix())

		// Define time ranges for the day (entry: 7 AM to 9 AM, exit: 2 PM to 4 PM)
		entryStart := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 7, 0, 0, 0, currentDate.Location()).Unix()
		entryEnd := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 9, 0, 0, 0, currentDate.Location()).Unix()
		exitStart := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 14, 0, 0, 0, currentDate.Location()).Unix()
		exitEnd := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 16, 0, 0, 0, currentDate.Location()).Unix()

		// Randomly decide if the user is present (70% chance to be present)
		if rand.Float64() > 0.3 {
			// Generate random entry and exit times
			entryTime := rand.Int63n(entryEnd-entryStart) + entryStart
			exitTime := rand.Int63n(exitEnd-exitStart) + exitStart

			// No existing record, create entrance record
			id, err := store.CreateEntrance(ctx, db.CreateEntranceParams{
				UserID:    userID,
				EnterTime: entryTime,
				Date:      date,
			})
			if err != nil {
				return err
			}

			err = store.UpdateExit(ctx, db.UpdateExitParams{
				AttendanceID: id,
				ExitTime:     exitTime,
			})
			if err != nil {
				return err
			}

		} else {
			fmt.Printf("User %d is absent on day %d\n", userID, day+1)
		}
	}
	return nil
}

func randomEntryTime() int64 {
	// Introduce a chance for the entry time to be 0 (e.g., 30% chance)
	if rand.Float64() < 0.3 { // 30% chance to return 0
		return 0
	}

	// Randomize the hour and minute between 7 AM and 9 AM as an example (or use other ranges)
	hour := rand.Intn(3) + 7 // 7 AM to 9 AM
	minute := rand.Intn(60)  // Any minute in the hour

	// Convert the hours and minutes to total seconds
	totalSeconds := int64(hour*3600 + minute*60)

	// Return the Unix time as the seconds from the start of the day
	return totalSeconds
}
