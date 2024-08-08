-- Students Queries

-- name: CreateStudent :exec
INSERT INTO students (name, student_id, class, face_image_path, fingerprint_path, email, phone, status) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetStudentByID :one
SELECT id, name, student_id, class, face_image_path, fingerprint_path, email, phone, status, created_at 
FROM students 
WHERE id = ?;

-- name: UpdateStudent :exec
UPDATE students 
SET name = ?, class = ?, face_image_path = ?, fingerprint_path = ?, email = ?, phone = ?, status = ? 
WHERE id = ?;

-- name: DeleteStudent :exec
DELETE FROM students WHERE id = ?;

-- Teachers Queries

-- name: CreateTeacher :exec
INSERT INTO teachers (name, teacher_id, department, face_image_path, fingerprint_path, email, phone, status) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetTeacherByID :one
SELECT id, name, teacher_id, department, face_image_path, fingerprint_path, email, phone, status, created_at 
FROM teachers 
WHERE id = ?;

-- name: UpdateTeacher :exec
UPDATE teachers 
SET name = ?, department = ?, face_image_path = ?, fingerprint_path = ?, email = ?, phone = ?, status = ? 
WHERE id = ?;

-- name: DeleteTeacher :exec
DELETE FROM teachers WHERE id = ?;

-- Attendance Records Queries

-- name: RecordEntry :exec
INSERT INTO attendance_records (user_type, user_id, entry_time) 
VALUES (?, ?, ?);

-- name: RecordExit :exec
UPDATE attendance_records 
SET exit_time = ? 
WHERE id = ?;

-- name: GetAttendanceByUserID :many
SELECT id, user_type, user_id, entry_time, exit_time 
FROM attendance_records 
WHERE user_id = ? AND user_type = ? 
ORDER BY entry_time DESC;

-- Parents Queries

-- name: CreateParent :exec
INSERT INTO parents (name, email, phone, relation) 
VALUES (?, ?, ?, ?);

-- name: LinkParentToStudent :exec
INSERT INTO student_parents (student_id, parent_id) 
VALUES (?, ?);

-- name: GetParentsByStudentID :many
SELECT parents.id, parents.name, parents.email, parents.phone, parents.relation, parents.created_at 
FROM parents 
JOIN student_parents ON parents.id = student_parents.parent_id 
WHERE student_parents.student_id = ?;

-- Notifications Queries

-- name: RecordNotification :exec
INSERT INTO notifications (user_type, user_id, notification_type, message, sent_at) 
VALUES (?, ?, ?, ?, ?);

-- name: GetNotificationsByUserID :many
SELECT id, user_type, user_id, notification_type, message, sent_at 
FROM notifications 
WHERE user_id = ? AND user_type = ? 
ORDER BY sent_at DESC;

-- Leave Requests Queries

-- name: CreateLeaveRequest :exec
INSERT INTO leave_requests (user_type, user_id, start_date, end_date, reason, status) 
VALUES (?, ?, ?, ?, ?, ?);

-- name: UpdateLeaveRequestStatus :exec
UPDATE leave_requests 
SET status = ? 
WHERE id = ?;

-- name: GetLeaveRequestsByUserID :many
SELECT id, user_type, user_id, start_date, end_date, reason, status, created_at 
FROM leave_requests 
WHERE user_id = ? AND user_type = ? 
ORDER BY created_at DESC;

-- Admins Queries

-- name: CreateAdmin :exec
INSERT INTO admins (name, email, password_hash, role) 
VALUES (?, ?, ?, ?);

-- name: GetAdminByID :one
SELECT id, name, email, role, created_at 
FROM admins 
WHERE id = ?;

-- name: GetAdminByEmail :one
SELECT id, name, email, password_hash, role, created_at 
FROM admins 
WHERE email = ?;

-- name: UpdateAdminRole :exec
UPDATE admins 
SET role = ? 
WHERE id = ?;

-- name: DeleteAdmin :exec
DELETE FROM admins WHERE id = ?;
