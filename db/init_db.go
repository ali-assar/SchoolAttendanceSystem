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
	ErrorCreateEntrance   = errors.New("could not create entrance table")
	ErrorCreateExit       = errors.New("could not create exit table")
	ErrorCreateAdmin      = errors.New("could not create admin table")
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
	db.Exec("DROP TABLE IF EXISTS admin")
	db.Exec("DROP TABLE IF EXISTS exit")
	db.Exec("DROP TABLE IF EXISTS entrance")

}

func CreateUsersTable(db *sql.DB) error {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		user_id INTEGER PRIMARY KEY,
		first_name VARCHAR(100) NOT NULL,
		last_name VARCHAR(100) NOT NULL,
		phone_number VARCHAR(50) NOT NULL,
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
		return errors.Join(ErrorCreateEntrance, err)
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
		return errors.Join(ErrorCreateExit, err)
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
		return errors.Join(ErrorCreateAdmin, err)
	}
	return nil
}
