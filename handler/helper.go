package handler

import (
	"context"
	"fmt"
	"strings"

	"github.com/Ali-Assar/SchoolAttendanceSystem/issues/db"
)

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
		firstName := strings.ReplaceAll(teacher.FirstName, " ", "-")
		lastName := strings.ReplaceAll(teacher.LastName, " ", "-")
		fullName := fmt.Sprintf("%s-%s", firstName, lastName)

		names = append(names, fullName)
	}
	name = strings.Join(names, ",")
	phone = absentTeachers[0].PhoneNumber

	return name, phone, nil
}
