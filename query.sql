-- name: CreateUser :one
INSERT INTO users (first_name, last_name, phone_number, image_path, is_teacher, is_biometric_active, finger_id) 
VALUES (?, ?, ?, ?, ?, ?, ?)
RETURNING user_id;

-- name: GetUserByID :one
SELECT user_id, first_name, last_name, phone_number, image_path, is_teacher, is_biometric_active, finger_id
FROM users 
WHERE user_id = ?;

-- name: GetUserByPhoneNumber :one
SELECT user_id, first_name, last_name, phone_number, image_path, is_teacher, is_biometric_active, finger_id
FROM users
WHERE phone_number = ?;

-- name: GetUserByName :one
SELECT user_id, first_name, last_name, phone_number, image_path, is_teacher, is_biometric_active, finger_id
FROM users
WHERE first_name = ? AND last_name = ?;


-- name: GetAllUsers :many
SELECT user_id, first_name, last_name, phone_number, image_path, is_teacher, is_biometric_active, finger_id 
FROM users;

-- name: UpdateUser :exec
UPDATE users 
SET first_name = ?, last_name = ?, phone_number = ?, image_path = ?, is_teacher = ?, is_biometric_active = ?, finger_id = ? 
WHERE user_id = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE user_id = ?;

-- name: CreateAttendance :one
INSERT INTO attendance (user_id, date, entry_time, exit_time) 
VALUES (?, ?, ?, ?)
RETURNING attendance_id;

-- name: GetAttendanceByUserIDAndDate :one
SELECT attendance_id, user_id, date, entry_time, exit_time 
FROM attendance 
WHERE user_id = ? AND date = ?;

-- name: GetAllUsersAttendanceByDate :many
SELECT 
    attendance.attendance_id, 
    attendance.user_id, 
    users.first_name, 
    users.last_name, 
    attendance.date, 
    attendance.entry_time, 
    attendance.exit_time
FROM 
    attendance
INNER JOIN 
    users ON attendance.user_id = users.user_id
WHERE 
    attendance.date = ?;

-- name: UpdateAttendance :exec
UPDATE attendance 
SET entry_time = ?, exit_time = ? 
WHERE user_id = ? AND attendance_id = ? AND date = ?;

-- name: DeleteAttendance :exec
DELETE FROM attendance WHERE attendance_id = ?;


-- name: CreateAdmin :one
INSERT INTO admin (user_name, password) 
VALUES (?, ?)
RETURNING user_name;

-- name: DeleteAdmin :exec
DELETE FROM admin 
WHERE user_name = ?;

-- name: GetAdminByUserName :one
SELECT user_name, password
FROM admin
WHERE user_name = ?;

-- name: UpdateAdmin :exec
UPDATE admin
SET password = ?
WHERE user_name = ?;

-- name: GetAbsentUsersUntil9AM :many
SELECT 
    users.user_id, 
    users.first_name, 
    users.last_name, 
    users.phone_number
FROM 
    users
LEFT JOIN 
    attendance ON users.user_id = attendance.user_id AND attendance.date = ? 
WHERE 
    attendance.entry_time IS NULL OR attendance.entry_time > 32400 -- 9 AM in seconds (9 * 60 * 60)
