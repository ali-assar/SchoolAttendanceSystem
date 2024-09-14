-- User Table 
CREATE TABLE IF NOT EXISTS users (
    user_id INTEGER PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    phone_number VARCHAR(50) NOT NULL,
    image_path TEXT NOT NULL DEFAULT NULL,
    finger_id TEXT NOT NULL DEFAULT NULL,
    is_biometric_active BOOLEAN NOT NULL DEFAULT FALSE
);

-- Teacher Table 
CREATE TABLE IF NOT EXISTS teachers (
    user_id INTEGER PRIMARY KEY,
    sunday_entry_time INTEGER NOT NULL,
    monday_entry_time INTEGER NOT NULL,
    tuesday_entry_time INTEGER NOT NULL,
    wednesday_entry_time INTEGER NOT NULL,
    thursday_entry_time INTEGER NOT NULL,
    friday_entry_time INTEGER NOT NULL,
    saturday_entry_time INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

-- Student Table 
CREATE TABLE IF NOT EXISTS students (
    user_id INTEGER PRIMARY KEY,
    required_entry_time INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS attendance (
    attendance_id INTEGER PRIMARY KEY,
    user_id INTEGER NOT NULL,
    date INTEGER NOT NULL DEFAULT NULL,
    enter_time INTEGER NOT NULL DEFAULT NULL,
    exit_time INTEGER NOT NULL DEFAULT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

-- Admin Table
CREATE TABLE IF NOT EXISTS admin (
    user_name VARCHAR(100) PRIMARY KEY,
    password VARCHAR(100) NOT NULL UNIQUE
);
