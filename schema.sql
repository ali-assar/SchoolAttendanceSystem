-- Users Table
CREATE TABLE users (
    user_id INTEGER PRIMARY KEY AUTOINCREMENT,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    phone_number TEXT,
    image_path TEXT,
    role_id INTEGER,
    is_admin INTEGER DEFAULT 0,
    FOREIGN KEY (role_id) REFERENCES roles(role_id)
);

-- Roles Table
CREATE TABLE roles (
    role_id INTEGER PRIMARY KEY AUTOINCREMENT,
    role_name TEXT NOT NULL
);

-- Attendance Table
CREATE TABLE attendance (
    attendance_id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER,
    date TEXT NOT NULL,
    entry_time TEXT,
    exit_time TEXT,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);