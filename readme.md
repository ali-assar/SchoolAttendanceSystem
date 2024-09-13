# School Attendance System

This system is designed to manage users and attendance records in a school. It includes functionalities for creating, updating, retrieving, and deleting users and attendance records. The system also supports authentication for securing routes, ensuring only authorized users have access to sensitive endpoints.

Here is the updated section for **Authentication** in your Markdown file to match the current setup, which includes JWT-based authentication:

---

## Authentication

The system uses **JWT (JSON Web Token)** authentication to secure the routes, ensuring that only authorized users can access or modify user and attendance data. The admin role (username: `admin`, password: `admin`) has full access to all routes. After logging in, a JWT token is issued, which must be included in the `Authorization` header for all subsequent requests to protected routes.

### JWT Token Authentication Flow:

1. **Login Endpoint**: Users must authenticate by providing valid credentials. A JWT token is generated upon successful login and must be used for any further requests to the system.
2. **Authorization**: JWT tokens should be included in the `Authorization` header of requests to protected routes. The format is `Bearer <token>`.

---

### Login Endpoint

#### 1. Login
- **Method**: `POST /login`
- **Description**: Authenticates a user and provides a JWT token for subsequent requests.
  
##### Request Body:

```json
{
  "username": "admin",
  "password": "admin"
}
```

##### Response:

- **Status 200 (OK)**: JWT token is provided.
  
```json
{
  "token": "your_jwt_token"
}
```

- **Status 401 (Unauthorized)**: When the username or password is incorrect.


### Securing Endpoints with JWT

Protected routes are secured by the `JWTAuthentication` middleware. You must include a valid JWT token in the `Authorization` header to access these routes.


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
  "is_teacher": true
}
```

#### Response:

- **Status 201 (Created):** 
```json
{
  "message": "id: 1"
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
        "image_path": "",
        "is_teacher": true,
        "is_biometric_active": false,
        "finger_id": ""
  },
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
    "image_path": "",
    "is_teacher": true,
    "is_biometric_active": false,
    "finger_id": ""
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
    "image_path": "",
    "is_teacher": true,
    "is_biometric_active": false,
    "finger_id": ""
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
    "image_path": "",
    "is_teacher": true,
    "is_biometric_active": false,
    "finger_id": ""
}
```

### 6. Update User biometrics

**Endpoint:** `PUT api/v1/user/:id`

Updates a user's biometrics information.

#### Request Body:

```json
{
{
  "image_path": "/home/pic",
  "is_biometric_active": true,
  "finger_id":  "asdfjdshf4456"
}
}
```

#### Response:

- **Status 200 (OK):**
```json
{
  "message": "ID: 1"
}
```

### 7. Update User image

**Endpoint:** `PUT api/v1/user/image/:id`

Updates a user's image.

#### Request Body:

```json
{
{
  "image_path": "/home/pic"
}
}
```

#### Response:

- **Status 200 (OK):**
```json
{
  "message": "ID: 1"
}
```

### 8. Delete User

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

### 1. Create Entrance

**Endpoint:** `POST api/v1/entrance`

Creates a new entrance record.

#### Request Body:

```json
{
  "user_id": 1,
  "entry_time": 1726115400
}
```

#### Response:

- **Status 201 (Created):**
```json
{
  "id": 1
}
```

### 2. Update Entrance

**Endpoint:** `PUT api/v1/entrance/:id`

Updates an entrance record.

#### Request Body:

```json
{
  "entry_time": 1726115400
}
```

#### Response:

- **Status 201 (Created):**
```json
{
  "id": 1
}
```

### 3. DELETE Entrance

**Endpoint:** `POST api/v1/entrance/:id`

Deletes an entrance record.

#### Response:

- **Status 201 (Created):**
```json
{
  "deleted"
}
```

### 4. Create Exit

**Endpoint:** `POST api/v1/exit`

Creates a new exit record.

#### Request Body:

```json
{
  "user_id": 1,
  "exit_time": 1726115400
}
```

#### Response:

- **Status 201 (Created):**
```json
{
  "id": 1
}
```

### 5. Update Exit

**Endpoint:** `PUT api/v1/exit/:id`

Updates an exit record.

#### Request Body:

```json
{
  "exit_time": 1726115400
}
```

#### Response:

- **Status 201 (Created):**
```json
{
  "id": 1
}
```

### 6. DELETE Exit

**Endpoint:** `POST api/v1/exit/:id`

Deletes an exit record.

#### Response:

- **Status 201 (Created):**
```json
{
  "deleted"
}
```
## Attendance Endpoints (Time Range and Absent Users)

### 1. Get Attendance by Time Range

**Endpoint:** `GET api/v1/attendance/`

This endpoint retrieves all users' attendance records (entry and exit times) within a specified time range.

#### Request Body (Example for JSON input):

```json
{
  "start_time": 1726115400,
  "end_time": 1726175400
}
```

#### Response:

- **Status 200 (OK):**
  
```json
[
  {
    "user_id": 1,
    "first_name": "John",
    "last_name": "Doe",
    "phone_number": "123456789",
    "entry_time": 1726115400,
    "exit_time": 1726165400
  },
  {
    "user_id": 2,
    "first_name": "Jane",
    "last_name": "Smith",
    "phone_number": "987654321",
    "entry_time": 1726116400,
    "exit_time": 1726170400
  }
]
```

---

### 2. Get Attendance by User ID (Time Range)

**Endpoint:** `GET api/v1/attendance/user`

This endpoint retrieves a specific user's attendance records (entry and exit times) within a specified time range, with the user's ID.

#### Request Body (Example for JSON input):

```json
{
  "user_id": 1,
  "start_time": 1726115400,
  "end_time": 1726175400
}
```

#### Response:

- **Status 200 (OK):**

```json
{
  "user_id": 1,
  "first_name": "John",
  "last_name": "Doe",
  "phone_number": "123456789",
  "entry_time": 1726115400,
  "exit_time": 1726165400
}
```

---

### 3. Get Absent Users

**Endpoint:** `GET api/v1/absents/`

This endpoint retrieves users who are absent (i.e., those who have no attendance records within a specified time range).

#### Request Body (Example for JSON input):

```json
{
  "start_time": 1726115400,
  "end_time": 1726175400
}
```

#### Response:

- **Status 200 (OK):**
  
```json
[
  {
    "user_id": 3,
    "first_name": "Emily",
    "last_name": "Johnson",
    "phone_number": "555555555",
    "entry_time": {
            "Int64": 0,
            "Valid": false
        }
  },
  {
    "user_id": 4,
    "first_name": "Michael",
    "last_name": "Brown",
    "phone_number": "444444444",
    "entry_time": {
            "Int64": 0,
            "Valid": false
        }
  }
]
```
