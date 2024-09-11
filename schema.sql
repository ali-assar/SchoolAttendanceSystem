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

CREATE TABLE IF NOT EXISTS attendance (
    attendance_id INTEGER PRIMARY KEY,
    user_id INT NOT NULL,
    date INT NOT NULL,        -- Store date as integer in YYYYMMDD format
    entry_time INT,           -- Store time as integer in seconds (or minutes) since midnight
    exit_time INT,            -- Same for exit time
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS admin (
    user_name VARCHAR(100) PRIMARY KEY,
    password VARCHAR(100) NOT NULL UNIQUE
);

