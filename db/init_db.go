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
	ErrorCreateUsersTable = errors.New("could not create users table")
	ErrorCreateAttendance = errors.New("could not create attendance table")
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
	if err := CreateUsersTable(db); err != nil {
		return err
	}

	if err := CreateAttendanceTable(db); err != nil {
		return err
	}

	return nil
}

func TearDown(db *sql.DB) {
	db.Exec("DROP TABLE IF EXISTS users")
	db.Exec("DROP TABLE IF EXISTS roles")
	db.Exec("DROP TABLE IF EXISTS attendance")
}

func CreateUsersTable(db *sql.DB) error {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		user_id INTEGER PRIMARY KEY,
		first_name VARCHAR(100) NOT NULL,
		last_name VARCHAR(100) NOT NULL,
		phone_number INTEGER NOT NULL,
		is_teacher BOOLEAN NOT NULL DEFAULT FALSE,
	
		image_path TEXT,
		finger_id TEXT,
		is_biometric_active BOOLEAN DEFAULT FALSE
	);
	`
	_, err := db.Exec(createUsersTable)
	if err != nil {
		return errors.Join(ErrorCreateUsersTable, err)
	}
	return nil
}

func CreateAttendanceTable(db *sql.DB) error {
	createAttendanceTable := `
	CREATE TABLE IF NOT EXISTS attendance (
		attendance_id INTEGER PRIMARY KEY,
		user_id INT NOT NULL,
		date DATE NOT NULL,
		entry_time TIME,
		exit_time TIME,
		FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
	);
	`
	_, err := db.Exec(createAttendanceTable)
	if err != nil {
		return errors.Join(ErrorCreateAttendance, err)
	}
	return nil
}
