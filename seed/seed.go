package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"math/rand"
// 	"time"

// 	"github.com/Ali-Assar/SchoolAttendanceSystem/issues/db"
// 	"github.com/joho/godotenv"
// )

// func main() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	database, err := db.InitDB()
// 	if err != nil {
// 		log.Fatalf("Failed to initialize the database: %v", err)
// 	}
// 	defer database.Close()

// 	db.TearDown(database)

// 	err = db.CreateTables(database)
// 	if err != nil {
// 		log.Fatalf("Failed to create tables: %v", err)
// 	}

// 	store := db.New(database)
// 	ctx := context.Background()

// 	err = seedDB(ctx, store)
// 	if err != nil {
// 		log.Fatalf("Failed to insert to the tables: %v", err)
// 	}
// }

// func seedDB(ctx context.Context, store db.Querier) error {
// 	var totalUsers = 100
// 	var teacherCount = totalUsers / 10 // 10% of users will be teachers
// 	var studentCount = totalUsers - teacherCount

// 	rand.Seed(time.Now().UnixNano())

// 	// Seed teachers
// 	for i := 0; i < teacherCount; i++ {
// 		// Create teacher user
// 		userParams := db.CreateUserParams{
// 			FirstName:   fmt.Sprintf("%d Teacher", i),
// 			LastName:    fmt.Sprintf("Lastname %d", i),
// 			PhoneNumber: fmt.Sprintf("09%d", rand.Intn(1000000000)),
// 		}

// 		userID, err := store.CreateUser(ctx, userParams)
// 		if err != nil {
// 			return err
// 		}

// 		// Insert teacher-specific data
// 		teacherParams := db.CreateTeacherParams{
// 			UserID: userID,
// 			// Add other fields as per your schema
// 		}

// 		_,err = store.CreateTeacher(ctx, teacherParams)
// 		if err != nil {
// 			return err
// 		}

// 		// Seed entrance and exit times for 30 days (same as before)
// 		err = seedAttendance(ctx, store, userID)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	// Seed students
// 	for i := 0; i < studentCount; i++ {
// 		// Create student user
// 		userParams := db.CreateUserParams{
// 			FirstName:   fmt.Sprintf("%d Student", i),
// 			LastName:    fmt.Sprintf("Lastname %d", i),
// 			PhoneNumber: fmt.Sprintf("09%d", rand.Intn(1000000000)),
// 		}

// 		userID, err := store.CreateUser(ctx, userParams)
// 		if err != nil {
// 			return err
// 		}

// 		// Insert student-specific data
// 		studentParams := db.CreateStudentParams{
// 			UserID: userID,
// 			// Add other fields as per your schema
// 		}

// 		_,err = store.CreateStudent(ctx, studentParams)
// 		if err != nil {
// 			return err
// 		}

// 		// Seed entrance and exit times for 30 days (same as before)
// 		err = seedAttendance(ctx, store, userID)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// func seedAttendance(ctx context.Context, store db.Querier, userID int64) error {
// 	for day := 0; day < 30; day++ {
// 		// Random date for each day in the range
// 		currentDate := time.Now().AddDate(0, 0, -day)

// 		// Define time ranges for the day (entry: 7 AM to 9 AM, exit: 2 PM to 4 PM)
// 		entryStart := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 7, 0, 0, 0, currentDate.Location()).Unix()
// 		entryEnd := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 9, 0, 0, 0, currentDate.Location()).Unix()
// 		exitStart := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 14, 0, 0, 0, currentDate.Location()).Unix()
// 		exitEnd := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 16, 0, 0, 0, currentDate.Location()).Unix()

// 		// Randomly decide if the user is present (70% chance to be present)
// 		if rand.Float64() > 0.3 {
// 			// Generate random entry and exit times
// 			entryTime := rand.Int63n(entryEnd-entryStart) + entryStart
// 			exitTime := rand.Int63n(exitEnd-exitStart) + exitStart

// 			// Insert entrance time
// 			entranceParams := db.CreateEntranceParams{
// 				UserID:    userID,
// 				EntryTime: entryTime,
// 			}
// 			_, err := store.CreateEntrance(ctx, entranceParams)
// 			if err != nil {
// 				return err
// 			}

// 			// Insert exit time
// 			exitParams := db.CreateExitParams{
// 				UserID:   userID,
// 				ExitTime: exitTime,
// 			}
// 			_, err = store.CreateExit(ctx, exitParams)
// 			if err != nil {
// 				return err
// 			}
// 		} else {
// 			fmt.Printf("User %d is absent on day %d\n", userID, day+1)
// 		}
// 	}
// 	return nil
// }
