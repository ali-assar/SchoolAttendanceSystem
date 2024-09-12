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

-- name: CreateEntrance :one
INSERT INTO entrance (user_id, entry_time) 
VALUES (?, ?)
RETURNING id;

-- name: UpdateEntrance :exec
UPDATE entrance 
SET entry_time = ?
WHERE id = ?;

-- name: DeleteEntrance :exec
DELETE FROM entrance 
WHERE id = ?;

-- name: CreateExit :one
INSERT INTO exit (user_id, exit_time) 
VALUES (?, ?)
RETURNING id;

-- name: UpdateExit :exec
UPDATE exit 
SET exit_time = ?
WHERE id = ?;

-- name: DeleteExit :exec
DELETE FROM exit 
WHERE id = ?;

-- name: GetTimeRange :many
SELECT 
    u.first_name,
    u.last_name,
    u.phone_number,
    e.entry_time,
    ex.exit_time
FROM 
    users u
JOIN 
    entrance e ON u.user_id = e.user_id
JOIN 
    exit ex ON u.user_id = ex.user_id
WHERE 
    e.entry_time >= ?
    AND ex.exit_time <= ?
    AND e.entry_time <= ex.exit_time;

-- name: GetTimeRangeByUserID :many
SELECT 
    u.user_id,           
    u.first_name,
    u.last_name,
    u.phone_number,
    e.entry_time,
    ex.exit_time
FROM 
    users u
JOIN 
    entrance e ON u.user_id = e.user_id
JOIN 
    exit ex ON u.user_id = ex.user_id
WHERE 
    u.user_id = ?           
    AND e.entry_time >= ?   
    AND ex.exit_time <= ?   
    AND e.entry_time <= ex.exit_time; 

