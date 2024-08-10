package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

var (
	ErrorCreateStudentsTable     = errors.New("could not create students table")
	ErrorCreateTeachersTable     = errors.New("could not create teachers table")
	ErrorCreateAttendanceRecords = errors.New("could not create attendance records table")
	ErrorCreateParentsTable      = errors.New("could not create parents table")
	ErrorCreateStudentParents    = errors.New("could not create student parents table")
	ErrorCreateNotifications     = errors.New("could not create notifications table")
	ErrorCreateLeaveRequests     = errors.New("could not create leave requests table")
	ErrorCreateAdminsTable       = errors.New("could not create admins table")
)

type DBParameter struct {
	DBPath string
}

func NewDBParameter(param DBParameter) *DBParameter {
	return &DBParameter{
		DBPath: param.DBPath,
	}
}

func InitDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file")
	}

	params := DBParameter{
		DBPath: os.Getenv("DB_PATH"),
	}

	connection := fmt.Sprintf("%s?_foreign_keys=on", params.DBPath)

	db, err := sql.Open("sqlite3", connection)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(10)

	return db, err
}

func CreateTables(db *sql.DB) error {
	if err := CreateStudentsTable(db); err != nil {
		return err
	}

	if err := CreateTeachersTable(db); err != nil {
		return err
	}

	if err := CreateAttendanceRecordsTable(db); err != nil {
		return err
	}

	if err := CreateParentsTable(db); err != nil {
		return err
	}

	if err := CreateStudentParentsTable(db); err != nil {
		return err
	}

	if err := CreateNotificationsTable(db); err != nil {
		return err
	}

	if err := CreateLeaveRequestsTable(db); err != nil {
		return err
	}

	if err := CreateAdminsTable(db); err != nil {
		return err
	}

	return nil
}

func TearDown(db *sql.DB) {
	db.Exec("DROP TABLE IF EXISTS students")
	db.Exec("DROP TABLE IF EXISTS teachers")
	db.Exec("DROP TABLE IF EXISTS attendance_records")
	db.Exec("DROP TABLE IF EXISTS parents")
	db.Exec("DROP TABLE IF EXISTS student_parents")
	db.Exec("DROP TABLE IF EXISTS notifications")
	db.Exec("DROP TABLE IF EXISTS leave_requests")
	db.Exec("DROP TABLE IF EXISTS admins")
}

func CreateStudentsTable(db *sql.DB) error {
	createStudentsTable := `
	CREATE TABLE IF NOT EXISTS students (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		student_id TEXT UNIQUE NOT NULL,
		class TEXT NOT NULL,
		face_image_path TEXT,
		fingerprint_path TEXT,
		email TEXT,
		phone TEXT,
		status TEXT DEFAULT 'active',
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := db.Exec(createStudentsTable)
	if err != nil {
		return errors.Join(ErrorCreateStudentsTable, err)
	}
	return nil
}

func CreateTeachersTable(db *sql.DB) error {
	createTeachersTable := `
	CREATE TABLE IF NOT EXISTS teachers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		teacher_id TEXT UNIQUE NOT NULL,
		department TEXT NOT NULL,
		face_image_path TEXT,
		fingerprint_path TEXT,
		email TEXT,
		phone TEXT,
		status TEXT DEFAULT 'active',
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := db.Exec(createTeachersTable)
	if err != nil {
		return errors.Join(ErrorCreateTeachersTable, err)
	}
	return nil
}

func CreateAttendanceRecordsTable(db *sql.DB) error {
	createAttendanceRecordsTable := `
	CREATE TABLE IF NOT EXISTS attendance_records (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_type TEXT NOT NULL,
		user_id INTEGER NOT NULL,
		entry_time DATETIME NOT NULL,
		exit_time DATETIME,
		FOREIGN KEY (user_id) REFERENCES students(id) ON DELETE CASCADE,
		FOREIGN KEY (user_id) REFERENCES teachers(id) ON DELETE CASCADE
	);
	`
	_, err := db.Exec(createAttendanceRecordsTable)
	if err != nil {
		return errors.Join(ErrorCreateAttendanceRecords, err)
	}
	return nil
}

func CreateParentsTable(db *sql.DB) error {
	createParentsTable := `
	CREATE TABLE IF NOT EXISTS parents (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT,
		phone TEXT,
		relation TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := db.Exec(createParentsTable)
	if err != nil {
		return errors.Join(ErrorCreateParentsTable, err)
	}
	return nil
}

func CreateStudentParentsTable(db *sql.DB) error {
	createStudentParentsTable := `
	CREATE TABLE IF NOT EXISTS student_parents (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		student_id INTEGER NOT NULL,
		parent_id INTEGER NOT NULL,
		FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE CASCADE,
		FOREIGN KEY (parent_id) REFERENCES parents(id) ON DELETE CASCADE
	);
	`
	_, err := db.Exec(createStudentParentsTable)
	if err != nil {
		return errors.Join(ErrorCreateStudentParents, err)
	}
	return nil
}

func CreateNotificationsTable(db *sql.DB) error {
	createNotificationsTable := `
	CREATE TABLE IF NOT EXISTS notifications (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_type TEXT NOT NULL,
		user_id INTEGER NOT NULL,
		notification_type TEXT NOT NULL,
		message TEXT NOT NULL,
		sent_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES students(id) ON DELETE CASCADE,
		FOREIGN KEY (user_id) REFERENCES teachers(id) ON DELETE CASCADE,
		FOREIGN KEY (user_id) REFERENCES parents(id) ON DELETE CASCADE
	);
	`
	_, err := db.Exec(createNotificationsTable)
	if err != nil {
		return errors.Join(ErrorCreateNotifications, err)
	}
	return nil
}

func CreateLeaveRequestsTable(db *sql.DB) error {
	createLeaveRequestsTable := `
	CREATE TABLE IF NOT EXISTS leave_requests (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_type TEXT NOT NULL,
		user_id INTEGER NOT NULL,
		start_date DATE NOT NULL,
		end_date DATE NOT NULL,
		reason TEXT,
		status TEXT DEFAULT 'pending',
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES students(id) ON DELETE CASCADE,
		FOREIGN KEY (user_id) REFERENCES teachers(id) ON DELETE CASCADE
	);
	`
	_, err := db.Exec(createLeaveRequestsTable)
	if err != nil {
		return errors.Join(ErrorCreateLeaveRequests, err)
	}
	return nil
}

func CreateAdminsTable(db *sql.DB) error {
	createAdminsTable := `
	CREATE TABLE IF NOT EXISTS admins (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		password_hash TEXT NOT NULL,
		role TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := db.Exec(createAdminsTable)
	if err != nil {
		return errors.Join(ErrorCreateAdminsTable, err)
	}
	return nil
}
