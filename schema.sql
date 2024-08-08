-- Students Table
CREATE TABLE students (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    student_id TEXT UNIQUE NOT NULL,
    class TEXT NOT NULL,
    face_image_path TEXT,
    fingerprint_path TEXT,
    email TEXT,
    phone TEXT,
    status TEXT DEFAULT 'active', -- 'active', 'inactive', 'graduated'
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Teachers Table
CREATE TABLE teachers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    teacher_id TEXT UNIQUE NOT NULL,
    department TEXT NOT NULL,
    face_image_path TEXT,
    fingerprint_path TEXT,
    email TEXT,
    phone TEXT,
    status TEXT DEFAULT 'active', -- 'active', 'inactive', 'retired'
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Attendance Records Table
CREATE TABLE attendance_records (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_type TEXT NOT NULL, -- 'student' or 'teacher'
    user_id INTEGER NOT NULL, -- student_id or teacher_id
    entry_time DATETIME NOT NULL,
    exit_time DATETIME,
    FOREIGN KEY (user_id) REFERENCES students(id),
    FOREIGN KEY (user_id) REFERENCES teachers(id)
);

-- Parents Table
CREATE TABLE parents (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT,
    phone TEXT,
    relation TEXT, -- relationship to the student
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Student Parents Table
CREATE TABLE student_parents (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    student_id INTEGER NOT NULL,
    parent_id INTEGER NOT NULL,
    FOREIGN KEY (student_id) REFERENCES students(id),
    FOREIGN KEY (parent_id) REFERENCES parents(id)
);

-- Notifications Table
CREATE TABLE notifications (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_type TEXT NOT NULL, -- 'student', 'teacher', 'parent'
    user_id INTEGER NOT NULL, -- student_id, teacher_id, or parent_id
    notification_type TEXT NOT NULL, -- 'email', 'sms', 'app'
    message TEXT NOT NULL,
    sent_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES students(id),
    FOREIGN KEY (user_id) REFERENCES teachers(id),
    FOREIGN KEY (user_id) REFERENCES parents(id)
);

-- Leave Requests Table
CREATE TABLE leave_requests (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_type TEXT NOT NULL, -- 'student' or 'teacher'
    user_id INTEGER NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    reason TEXT,
    status TEXT DEFAULT 'pending', -- 'pending', 'approved', 'denied'
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES students(id),
    FOREIGN KEY (user_id) REFERENCES teachers(id)
);

-- Admins Table
CREATE TABLE admins (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    role TEXT NOT NULL, -- 'super_admin', 'admin', 'moderator', 'teacher'
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
