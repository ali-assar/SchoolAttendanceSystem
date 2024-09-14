
-- name: CreateUser :one
INSERT INTO users (first_name, last_name, phone_number, image_path, finger_id, is_biometric_active)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING user_id;

-- name: GetUserByID :one
SELECT user_id, first_name, last_name, phone_number, image_path, finger_id, is_biometric_active
FROM users
WHERE user_id = ?;

-- name: GetUserByName :many
SELECT user_id, first_name, last_name, phone_number, image_path, is_biometric_active
FROM users
WHERE first_name = ? AND last_name = ?;

-- name: UpdateUser :exec
UPDATE users
SET first_name = ?, last_name = ?, phone_number = ?, image_path = ?
WHERE user_id = ?;


-- name: UpdateUserBiometric :exec
UPDATE users
SET image_path = ?, finger_id = ?, is_biometric_active = ?
WHERE user_id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE user_id = ?;

-- Queries for Teachers
-- name: CreateTeacher :one
INSERT INTO teachers (user_id, sunday_entry_time, monday_entry_time, tuesday_entry_time, wednesday_entry_time, thursday_entry_time, friday_entry_time, saturday_entry_time)
VALUES (?, ?, ?, ?, ?, ?, ?, ?)
RETURNING user_id;

-- name: GetTeacherByID :one
SELECT t.user_id, u.first_name, u.last_name, t.sunday_entry_time, t.monday_entry_time, t.tuesday_entry_time, t.wednesday_entry_time, t.thursday_entry_time, t.friday_entry_time, t.saturday_entry_time
FROM teachers t
JOIN users u ON t.user_id = u.user_id
WHERE t.user_id = ?;

-- name: UpdateTeacherAllowedTime :exec
UPDATE teachers
SET sunday_entry_time = ?, monday_entry_time = ?, tuesday_entry_time = ?, wednesday_entry_time = ?, thursday_entry_time = ?, friday_entry_time = ?, saturday_entry_time = ?
WHERE user_id = ?;

-- Queries for Students
-- name: CreateStudent :one
INSERT INTO students (user_id, required_entry_time)
VALUES (?, ?)
RETURNING user_id;

-- name: GetStudentByID :one
SELECT s.user_id, u.first_name, u.last_name, s.required_entry_time
FROM students s
JOIN users u ON s.user_id = u.user_id
WHERE s.user_id = ?;

-- name: UpdateStudentAllowedTime :exec
UPDATE students
SET required_entry_time = ?
WHERE user_id = ?;

-- Queries for Entrance
-- name: CreateEntrance :one
INSERT INTO entrance (user_id, entry_time)
VALUES (?, ?)
RETURNING id;

-- name: GetEntrancesByUserID :many
SELECT id, user_id, entry_time
FROM entrance
WHERE user_id = ?;

-- name: UpdateEntrance :exec
UPDATE entrance
SET entry_time = ?
WHERE id = ?;

-- name: DeleteEntrance :exec
DELETE FROM entrance
WHERE id = ?;

-- Queries for Exit
-- name: CreateExit :one
INSERT INTO exit (user_id, exit_time)
VALUES (?, ?)
RETURNING id;

-- name: GetExitsByUserID :many
SELECT id, user_id, exit_time
FROM exit
WHERE user_id = ?;

-- name: UpdateExit :exec
UPDATE exit
SET exit_time = ?
WHERE id = ?;

-- name: DeleteExit :exec
DELETE FROM exit
WHERE id = ?;

-- Queries for Admin
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
