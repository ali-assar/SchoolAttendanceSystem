package handler

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Ali-Assar/SchoolAttendanceSystem/issues/db"
)

// Helper function to apply ExtractUnixTime or return fetched value if zero
func ExtractUnixTimeOrFetched(newTime, fetchedTime int64) int64 {
	// if newTime == 0 {
	// 	return ExtractUnixTime(fetchedTime)
	// }
	return ExtractUnixTime(newTime)
}

func UnixToMinute(timestamp int64) int64 {
	return (timestamp % 3600) / 60
}

func ExtractUnixDate(timestamp int64) int64 {
	return (timestamp / 86400) * 86400
}

func UnixToDayOfWeek(date int64) int64 {
	return ((date / 86400) + 4) % 7
}

func FindAbsentTeachers(store db.Querier, ctx context.Context, date int) ([]db.GetAbsentTeachersByDateRow, error) {
	date = int(ExtractUnixDate(int64(date)))
	dayOfWeek := ((date / 86400) + 4) % 7

	absentUsers, err := store.GetAbsentTeachersByDate(ctx, int64(date))
	if err != nil {
		return nil, err
	}

	var absentOnDay []db.GetAbsentTeachersByDateRow
	for _, user := range absentUsers {
		switch dayOfWeek {
		case 1:
			if user.MondayEntryTime != 0 {
				absentOnDay = append(absentOnDay, user)
			}
		case 2:
			if user.TuesdayEntryTime != 0 {
				absentOnDay = append(absentOnDay, user)
			}
		case 3:
			if user.WednesdayEntryTime != 0 {
				absentOnDay = append(absentOnDay, user)
			}
		case 4:
			if user.ThursdayEntryTime != 0 {
				absentOnDay = append(absentOnDay, user)
			}
		case 5:
			if user.FridayEntryTime != 0 {
				absentOnDay = append(absentOnDay, user)
			}
		case 6:
			if user.SaturdayEntryTime != 0 {
				absentOnDay = append(absentOnDay, user)
			}
		case 0:
			if user.SundayEntryTime != 0 {
				absentOnDay = append(absentOnDay, user)
			}
		}
	}
	return absentOnDay, nil
}

func FindAbsentStudents(store db.Querier, ctx context.Context, date int) ([]db.GetAbsentStudentByDateRow, error) {
	date = int(ExtractUnixDate(int64(date)))
	dayOfWeek := ((date / 86400) + 4) % 7

	if dayOfWeek == 4 || dayOfWeek == 5 {
		return nil, nil
	}

	absentStudents, err := store.GetAbsentStudentByDate(ctx, int64(date))
	if err != nil {
		return nil, err
	}
	return absentStudents, nil
}

func GetFormattedAbsentTeachers(store db.Querier, ctx context.Context, date int) (name string, phone string, err error) {
	absentTeachers, err := FindAbsentTeachers(store, ctx, date)
	if err != nil {
		return "", "", err
	}

	var names []string
	for _, teacher := range absentTeachers {
		// firstName := strings.ReplaceAll(teacher.FirstName, " ", "-")
		// lastName := strings.ReplaceAll(teacher.LastName, " ", "-")
		// fullName := fmt.Sprintf("%s-%s", firstName, lastName)
		fullName := fmt.Sprintf("%s %s", teacher.FirstName, teacher.LastName)

		names = append(names, fullName)
	}
	name = strings.Join(names, ", ")
	if len(absentTeachers) > 0 {
		phone = absentTeachers[0].PhoneNumber
	}

	return name, phone, nil
}

func ExtractUnixTime(timestamp int64) int64 {
	return timestamp % 86400 // This strips off the date by keeping only seconds of the day
}

func GetLocalTimeOffset() int64 {
	nowLocal := time.Now()
	nowUTC := time.Now().UTC()

	hourOffset := int64(nowLocal.Hour() - nowUTC.Hour())
	minuteOffset := int64(nowLocal.Minute() - nowUTC.Minute())

	if hourOffset > 12 {
		hourOffset -= 24
	} else if hourOffset < -12 {
		hourOffset += 24
	}

	offsetInSeconds := hourOffset*3600 + minuteOffset*60
	fmt.Println(offsetInSeconds)
	return offsetInSeconds
}

func FindTeachersDelay(store db.Querier, ctx context.Context, date int) ([]db.GetFullDetailsTeacherAttendanceByDateRow, error) {
	date = int(ExtractUnixDate(int64(date)))
	dayOfWeek := ((date / 86400) + 4) % 7

	attendance, err := store.GetFullDetailsTeacherAttendanceByDate(ctx, int64(date))
	if err != nil {
		return nil, err
	}

	var delayOnDay []db.GetFullDetailsTeacherAttendanceByDateRow
	localTimeOffset := GetLocalTimeOffset()
	for _, user := range attendance {
		// Normalize the EnterTime to keep only hours, minutes, and seconds
		normalizedEnterTime := ExtractUnixTime(user.EnterTime) + localTimeOffset

		switch dayOfWeek {
		case 1:
			if normalizedEnterTime > user.MondayEntryTime+localTimeOffset && user.MondayEntryTime != 0 {
				delayOnDay = append(delayOnDay, user)
			}
		case 2:
			if normalizedEnterTime > user.TuesdayEntryTime+localTimeOffset && user.TuesdayEntryTime != 0 {
				delayOnDay = append(delayOnDay, user)
			}
		case 3:
			if normalizedEnterTime > user.WednesdayEntryTime+localTimeOffset && user.WednesdayEntryTime != 0 {
				delayOnDay = append(delayOnDay, user)
			}
		case 4:
			if normalizedEnterTime > user.ThursdayEntryTime+localTimeOffset && user.ThursdayEntryTime != 0 {
				delayOnDay = append(delayOnDay, user)
			}
		case 5:
			if normalizedEnterTime > user.FridayEntryTime+localTimeOffset && user.FridayEntryTime != 0 {
				delayOnDay = append(delayOnDay, user)
			}
		case 6:
			if normalizedEnterTime > user.SaturdayEntryTime+localTimeOffset && user.SaturdayEntryTime != 0 {
				delayOnDay = append(delayOnDay, user)
			}
		case 0:
			if normalizedEnterTime > user.SundayEntryTime+localTimeOffset && user.SundayEntryTime != 0 {
				delayOnDay = append(delayOnDay, user)
			}
		}
	}
	return delayOnDay, nil
}

func GetFormattedTeachersDelay(store db.Querier, ctx context.Context, date int) (name string, phone string, err error) {
	teachersDelay, err := FindTeachersDelay(store, ctx, date)
	if err != nil {
		return "", "", err
	}

	var names []string
	for _, teacher := range teachersDelay {
		// firstName := strings.ReplaceAll(teacher.FirstName, " ", "-")
		// lastName := strings.ReplaceAll(teacher.LastName, " ", "-")
		fullName := fmt.Sprintf("%s %s", teacher.FirstName, teacher.LastName)

		names = append(names, fullName)
	}
	name = strings.Join(names, ", ")
	if len(teachersDelay) > 0 {
		phone = teachersDelay[0].PhoneNumber
	}

	return name, phone, nil
}
