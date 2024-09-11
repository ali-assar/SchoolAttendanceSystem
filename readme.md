# School Attendance System

This system is designed to manage users and attendance records in a school. It includes functionalities for creating, updating, retrieving, and deleting users and attendance records. The system also supports authentication for securing routes, ensuring only authorized users have access to sensitive endpoints.

## Authentication

In this system, authentication ensures that only authorized individuals can add, modify, or delete user and attendance data. This protects the integrity of the system and ensures compliance with privacy standards.

In this system, the admin role (username: `admin`, password: `admin`) has access to all routes. Non-admin users might be restricted based on their role, ensuring secure access management.

## User Endpoints

### 1. Create User

**Endpoint:** `POST api/v1/user`

Creates a new user in the system.

#### Request Body:

```json
{
  "first_name": "John",
  "last_name": "Doe",
  "phone_number": "123456789",
  "image_path": null,
  "is_teacher": true,
  "is_biometric_active": false,
  "finger_id": null
}
```

#### Response:

- **Status 201 (Created):** 
```json
{
  "message": "user with id 1 created"
}
```

### 2. Get All Users

**Endpoint:** `GET api/v1/user`

Retrieves a list of all users in the system.

#### Response:

- **Status 200 (OK):**
```json
[
  {
    "user_id": 1,
    "first_name": "John",
    "last_name": "Doe",
    "phone_number": "123456789",
    "image_path": null,
    "is_teacher": true,
    "is_biometric_active": false,
    "finger_id": null
  }
]
```

### 3. Get User by ID

**Endpoint:** `GET api/v1/user/:id`

Retrieves a user by their ID.

#### Response:

- **Status 200 (OK):**
```json
{
  "user_id": 1,
  "first_name": "John",
  "last_name": "Doe",
  "phone_number": "123456789",
  "image_path": null,
  "is_teacher": true,
  "is_biometric_active": false,
  "finger_id": null
}
```

### 4. Get User by Phone Number

**Endpoint:** `GET api/v1/user/phone/:phone`

Retrieves a user by their phone number.

#### Response:

- **Status 200 (OK):**
```json
{
  "user_id": 1,
  "first_name": "John",
  "last_name": "Doe",
  "phone_number": "123456789",
  "image_path": null,
  "is_teacher": true,
  "is_biometric_active": false,
  "finger_id": null
}
```

### 5. Get User by Name

**Endpoint:** `GET api/v1/user/name/:first_name/:last_name`

Retrieves a user by their first and last name.

#### Response:

- **Status 200 (OK):**
```json
{
  "user_id": 1,
  "first_name": "John",
  "last_name": "Doe",
  "phone_number": "123456789",
  "image_path": null,
  "is_teacher": true,
  "is_biometric_active": false,
  "finger_id": null
}
```

### 6. Update User

**Endpoint:** `PUT api/v1/user/:id`

Updates a user's information.

#### Request Body:

```json
{
  "first_name": "Jane",
  "last_name": "Smith",
  "phone_number": "987654321",
  "image_path": null,
  "is_teacher": true,
  "is_biometric_active": true,
  "finger_id": null
}
```

#### Response:

- **Status 200 (OK):**
```json
{
  "message": "user updated"
}
```

### 7. Delete User

**Endpoint:** `DELETE api/v1/user/:id`

Deletes a user by their ID.

#### Response:

- **Status 200 (OK):**
```json
{
  "message": "user deleted"
}
```

---

## Attendance Endpoints

### 1. Create Attendance

**Endpoint:** `POST api/v1/attendance`

Creates a new attendance record.

#### Request Body:

```json
{
  "user_id": 1,
  "date": 1694019110000, 
  "entry_time": 1694022710000,
  "exit_time": 1694030000000
}
```

#### Response:

- **Status 201 (Created):**
```json
{
  "message": "attendance record created",
  "attendance_id": 1
}
```

### 2. Get Attendance by User ID and Date

**Endpoint:** `GET api/v1/attendance/:user_id/:date`

Retrieves an attendance record by user ID and date.

#### Response:

- **Status 200 (OK):**
```json
{
  "attendance_id": 1,
  "user_id": 1,
  "date": 1694019110000,
  "entry_time": 1694022710000,
  "exit_time": 1694030000000
}
```

### 3. Get All Users' Attendance by Date

**Endpoint:** `GET api/v1/attendances/date/:date`

Retrieves all attendance records for a specific date.

#### Response:

- **Status 200 (OK):**
```json
[
  {
    "attendance_id": 1,
    "user_id": 1,
    "first_name": "John",
    "last_name": "Doe",
    "date": 1694019110000,
    "entry_time": 1694022710000,
    "exit_time": 1694030000000
  }
]
```

### 4. Update Attendance

**Endpoint:** `PUT api/v1/attendance`

Updates an attendance record by attendance ID.

#### Request Body:

```json
{
  "attendance_id": 1,
  "entry_time": 1694022710000,
  "exit_time": 1694030000000
}
```

#### Response:

- **Status 200 (OK):**
```json
{
  "message": "attendance record updated"
}
```

### 5. Delete Attendance

**Endpoint:** `DELETE api/v1/attendance/:attendance_id`

Deletes an attendance record by attendance ID.

#### Response:

- **Status 200 (OK):**
```json
{
  "message": "attendance record deleted"
}
```

---

### Error Responses

- **Status 400 (Bad Request):** When the input data is invalid.
- **Status 404 (Not Found):** When the requested user or attendance record is not found.
- **Status 500 (Internal Server Error):** When an unexpected error occurs on the server.
