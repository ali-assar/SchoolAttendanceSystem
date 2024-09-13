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
	ErrorCreateUsersTable    = errors.New("could not create users table")
	ErrorCreateTeachersTable = errors.New("could not create teachers table")
	ErrorCreateStudentsTable = errors.New("could not create students table")
	ErrorCreateEntranceTable = errors.New("could not create entrance table")
	ErrorCreateExitTable     = errors.New("could not create exit table")
	ErrorCreateAdminTable    = errors.New("could not create admin table")
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

	if err := CreateTeachersTable(db); err != nil {
		return err
	}

	if err := CreateStudentsTable(db); err != nil {
		return err
	}

	if err := CreateEntranceTable(db); err != nil {
		return err
	}

	if err := CreateExitTable(db); err != nil {
		return err
	}

	if err := CreateAdminTable(db); err != nil {
		return err
	}

	return nil
}

func TearDown(db *sql.DB) {
	db.Exec("DROP TABLE IF EXISTS users")
	db.Exec("DROP TABLE IF EXISTS teachers")
	db.Exec("DROP TABLE IF EXISTS students")
	db.Exec("DROP TABLE IF EXISTS entrance")
	db.Exec("DROP TABLE IF EXISTS exit")
	db.Exec("DROP TABLE IF EXISTS admin")
}

func CreateUsersTable(db *sql.DB) error {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		user_id INTEGER PRIMARY KEY,
    	first_name VARCHAR(100) NOT NULL,
    	last_name VARCHAR(100) NOT NULL,
    	phone_number VARCHAR(50) NOT NULL,
    	image_path TEXT DEFAULT NULL,
    	finger_id TEXT DEFAULT NULL,
    	is_biometric_active BOOLEAN NOT NULL DEFAULT FALSE
	);
	`
	_, err := db.Exec(createUsersTable)
	if err != nil {
		return errors.Join(ErrorCreateUsersTable, err)
	}
	return nil
}

func CreateTeachersTable(db *sql.DB) error {
	createTeachersTable := `
	CREATE TABLE IF NOT EXISTS teachers (
		teacher_id INTEGER PRIMARY KEY,
		user_id INTEGER NOT NULL,
		sunday_entry_time INTEGER NOT NULL,
		monday_entry_time INTEGER NOT NULL,
		tuesday_entry_time INTEGER NOT NULL,
		wednesday_entry_time INTEGER NOT NULL,
		thursday_entry_time INTEGER NOT NULL,
		friday_entry_time INTEGER NOT NULL,
		saturday_entry_time INTEGER NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
	);
	`
	_, err := db.Exec(createTeachersTable)
	if err != nil {
		return errors.Join(ErrorCreateTeachersTable, err)
	}
	return nil
}

func CreateStudentsTable(db *sql.DB) error {
	createStudentsTable := `
	CREATE TABLE IF NOT EXISTS students (
		student_id INTEGER PRIMARY KEY,
		user_id INTEGER NOT NULL,
		required_entry_time INTEGER NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
	);
	`
	_, err := db.Exec(createStudentsTable)
	if err != nil {
		return errors.Join(ErrorCreateStudentsTable, err)
	}
	return nil
}

func CreateEntranceTable(db *sql.DB) error {
	createEntranceTable := `
	CREATE TABLE IF NOT EXISTS entrance (
		id INTEGER PRIMARY KEY,
		user_id INTEGER NOT NULL,
		entry_time INTEGER NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
	);
	`
	_, err := db.Exec(createEntranceTable)
	if err != nil {
		return errors.Join(ErrorCreateEntranceTable, err)
	}
	return nil
}

func CreateExitTable(db *sql.DB) error {
	createExitTable := `
	CREATE TABLE IF NOT EXISTS exit (
		id INTEGER PRIMARY KEY,
		user_id INTEGER NOT NULL,
		exit_time INTEGER NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
	);
	`
	_, err := db.Exec(createExitTable)
	if err != nil {
		return errors.Join(ErrorCreateExitTable, err)
	}
	return nil
}

func CreateAdminTable(db *sql.DB) error {
	createAdminTable := `
	CREATE TABLE IF NOT EXISTS admin (
		user_name VARCHAR(100) PRIMARY KEY,
		password VARCHAR(100) NOT NULL UNIQUE
	);
	`
	_, err := db.Exec(createAdminTable)
	if err != nil {
		return errors.Join(ErrorCreateAdminTable, err)
	}
	return nil
}
