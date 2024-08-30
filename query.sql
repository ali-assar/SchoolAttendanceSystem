
-- name: CreateUser :exec
INSERT INTO users (first_name, last_name, phone_number, image_path, is_teacher, is_admin) 
VALUES (?, ?, ?, ?, ?, ?);

-- name: GetUserByID :one
SELECT user_id, first_name, last_name, phone_number, image_path, is_teacher, is_admin 
FROM users 
WHERE user_id = ?;

-- name: GetAllUsers :many
SELECT user_id, first_name, last_name, phone_number, image_path, is_teacher, is_admin 
FROM users;

-- name: UpdateUser :exec
UPDATE users 
SET first_name = ?, last_name = ?, phone_number = ?, image_path = ?, is_teacher = ?, is_admin = ? 
WHERE user_id = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE user_id = ?;

-- name: CreateAttendance :exec
INSERT INTO attendance (user_id, date, entry_time, exit_time) 
VALUES (?, ?, ?, ?);

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
WHERE user_id = ? AND date = ?;

-- name: DeleteAttendance :exec
DELETE FROM attendance WHERE attendance_id = ?;
