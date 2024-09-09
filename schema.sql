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

CREATE TABLE IF NOT EXISTS attendance (
    attendance_id INTEGER PRIMARY KEY,
    user_id INT NOT NULL,
    date DATE NOT NULL,
    entry_time TIME,
    exit_time TIME,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);
