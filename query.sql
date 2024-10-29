
-- name: CreateUser :one
INSERT INTO users (first_name, last_name, phone_number, image_path, finger_id, is_biometric_active, created_at)
VALUES (?, ?, ?, ?, ?, ?, ?)
RETURNING user_id;

-- name: GetUserByID :one
SELECT user_id, first_name, last_name, phone_number, image_path, finger_id, is_biometric_active, created_at
FROM users
WHERE user_id = ?;

-- name: GetUserByName :many
SELECT user_id, first_name, last_name, phone_number, image_path, is_biometric_active, created_at
FROM users
WHERE first_name = ? AND last_name = ?;

-- name: UpdateUserDetails :exec
UPDATE users
SET first_name = ?, last_name = ?, phone_number = ?, image_path = ?
WHERE user_id = ?;

-- name: GetUsersWithFalseBiometric :many
SELECT user_id, is_biometric_active, first_name, last_name
FROM users
WHERE is_biometric_active = false;

-- name: GetUsersWithTrueBiometric :many
SELECT user_id, is_biometric_active, image_path, finger_id
FROM users
WHERE is_biometric_active = true;

-- name: UpdateUserBiometric :exec
UPDATE users
SET image_path = ?, finger_id = ?, is_biometric_active = ?
WHERE user_id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE user_id = ?;

-- name: CreateTeacher :one
INSERT INTO teachers (user_id, sunday_entry_time, monday_entry_time, tuesday_entry_time, wednesday_entry_time, thursday_entry_time, friday_entry_time, saturday_entry_time)
VALUES (?, ?, ?, ?, ?, ?, ?, ?)
RETURNING user_id;

-- name: GetTeacherByID :one
SELECT t.user_id, u.first_name, u.last_name, u.phone_number ,u.created_at, t.sunday_entry_time, t.monday_entry_time, t.tuesday_entry_time, t.wednesday_entry_time, t.thursday_entry_time, t.friday_entry_time, t.saturday_entry_time
FROM teachers t
JOIN users u ON t.user_id = u.user_id
WHERE t.user_id = ?;

-- name: GetTeachers :many
SELECT t.user_id, u.first_name, u.last_name, u.phone_number, u.created_at, t.sunday_entry_time, t.monday_entry_time, t.tuesday_entry_time, t.wednesday_entry_time, t.thursday_entry_time, t.friday_entry_time, t.saturday_entry_time
FROM teachers t
JOIN users u ON t.user_id = u.user_id;

-- name: UpdateTeacherAllowedTime :exec
UPDATE teachers
SET sunday_entry_time = ?, monday_entry_time = ?, tuesday_entry_time = ?, wednesday_entry_time = ?, thursday_entry_time = ?, friday_entry_time = ?, saturday_entry_time = ?
WHERE user_id = ?;

-- name: CreateStudent :one
INSERT INTO students (user_id, required_entry_time)
VALUES (?, ?)
RETURNING user_id;

-- name: GetStudentByID :one
SELECT s.user_id, u.first_name, u.last_name, u.phone_number,u.created_at, s.required_entry_time
FROM students s
JOIN users u ON s.user_id = u.user_id
WHERE s.user_id = ?;

-- name: GetStudents :many
SELECT s.user_id, u.first_name, u.last_name, u.phone_number,u.created_at, s.required_entry_time
FROM students s
JOIN users u ON s.user_id = u.user_id;

-- name: UpdateStudentAllowedTime :exec
UPDATE students
SET required_entry_time = ?
WHERE user_id = ?;

-- name: CreateAdmin :one
INSERT INTO admin (user_name, password)
VALUES (?, ?)
RETURNING user_name;

-- name: GetAdminByUserName :one
SELECT user_name, password
FROM admin
WHERE user_name = ?;

-- name: UpdateAdmin :exec
UPDATE admin
SET password = ?
WHERE user_name = ?;

-- name: DeleteAdmin :exec
DELETE FROM admin
WHERE user_name = ?;


-- name: CreateEntrance :one
INSERT INTO attendance (user_id, date, enter_time, exit_time)
VALUES (?, ?, ?, 0)
RETURNING attendance_id;

-- name: UpdateExit :exec
UPDATE attendance
SET exit_time = ?
WHERE attendance_id = ?;

-- name: UpdateEntranceByID :exec
UPDATE attendance
SET enter_time = ?
WHERE attendance_id = ?;

-- name: DeleteAttendance :exec
DELETE FROM attendance
WHERE attendance_id = ?;

-- name: GetAttendanceByUserIDAndDate :many
SELECT attendance_id, user_id, date, enter_time, exit_time
FROM attendance
WHERE user_id = ? AND date = ?;

-- name: GetAttendanceByUserID :many
SELECT attendance_id, user_id, date, enter_time, exit_time
FROM attendance
WHERE user_id = ?;

-- name: GetAttendanceByDate :many
SELECT a.attendance_id, a.user_id, u.first_name, u.last_name, a.date, a.enter_time, a.exit_time
FROM attendance a
JOIN users u ON a.user_id = u.user_id 
WHERE date = ?;

-- name: GetAttendanceBetweenDates :many
SELECT a.attendance_id, a.user_id, u.first_name, u.last_name, a.date, a.enter_time, a.exit_time
FROM attendance a
JOIN users u ON a.user_id = u.user_id 
WHERE date BETWEEN ? AND ?;

-- name: GetAbsentUsersByDate :many
SELECT u.user_id, u.first_name, u.last_name
FROM users u
LEFT JOIN attendance a ON u.user_id = a.user_id AND a.date = ?
WHERE a.user_id IS NULL;

-- name: GetAbsentTeachersByDate :many
SELECT u.user_id, u.first_name, u.last_name, u.phone_number,t.sunday_entry_time, t.monday_entry_time, t.tuesday_entry_time, t.wednesday_entry_time, t.thursday_entry_time, t.friday_entry_time, t.saturday_entry_time
FROM users u
JOIN teachers t ON u.user_id = t.user_id
LEFT JOIN attendance a ON u.user_id = a.user_id AND a.date = ?
WHERE a.user_id IS NULL;

-- name: GetAbsentStudentByDate :many
SELECT u.user_id, u.first_name, u.last_name, u.phone_number, s.required_entry_time
FROM users u
JOIN students s ON u.user_id = s.user_id
LEFT JOIN attendance a ON u.user_id = a.user_id AND a.date = ?
WHERE a.user_id IS NULL;



-- name: GetTeacherAttendanceByDate :many
SELECT 
    a.attendance_id, 
    a.user_id, 
    u.first_name, 
    u.last_name, 
    a.date, 
    a.enter_time, 
    a.exit_time,
    u.phone_number
FROM attendance a
JOIN users u ON a.user_id = u.user_id
JOIN teachers t ON u.user_id = t.user_id
WHERE a.date = ?;

-- name: GetStudentAttendanceByDate :many
SELECT 
    a.attendance_id, 
    a.user_id, 
    u.first_name, 
    u.last_name, 
    a.date, 
    a.enter_time, 
    a.exit_time,
    u.phone_number

FROM attendance a
JOIN users u ON a.user_id = u.user_id
JOIN students s ON u.user_id = s.user_id
WHERE a.date = ?;

-- name: GetTeacherAttendanceBetweenDates :many
SELECT 
    a.attendance_id, 
    a.user_id, 
    u.first_name, 
    u.last_name, 
    a.date, 
    a.enter_time, 
    a.exit_time,
    u.phone_number

FROM attendance a
JOIN users u ON a.user_id = u.user_id
JOIN teachers t ON u.user_id = t.user_id
WHERE a.date BETWEEN ? AND ?;

-- name: GetStudentAttendanceBetweenDates :many
SELECT 
    a.attendance_id, 
    a.user_id, 
    u.first_name, 
    u.last_name, 
    a.date, 
    a.enter_time, 
    a.exit_time,
    u.phone_number

FROM attendance a
JOIN users u ON a.user_id = u.user_id
JOIN students s ON u.user_id = s.user_id
WHERE a.date BETWEEN ? AND ?;


-- name: GetFullDetailsTeacherAttendanceByDate :many
SELECT 
    a.attendance_id, 
    a.user_id, 
    u.first_name, 
    u.last_name, 
    u.phone_number,
    t.sunday_entry_time, 
    t.monday_entry_time, 
    t.tuesday_entry_time, 
    t.wednesday_entry_time, 
    t.thursday_entry_time, 
    t.friday_entry_time, 
    t.saturday_entry_time, 
    a.enter_time,     
    a.date
FROM attendance a
JOIN users u ON a.user_id = u.user_id
JOIN teachers t ON u.user_id = t.user_id
WHERE a.date = ?;


