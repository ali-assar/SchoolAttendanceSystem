
# School Attendance System

This system manages the attendance of students and teachers, incorporating user authentication and biometric services. The system is built using the Fiber web framework, JWT for authentication, and SQLite as the database.

## Authentication

- **Login**:  
  `POST /login`  
  This endpoint is used for authenticating users and generating a JWT token.

  **Request Body (JSON)**:
  ```json
  {
    "username": "admin",
    "password": "admin"
  }
  ```

## API Endpoints

### Authenticated Endpoints (`/api/v1`)

These routes require a valid JWT token.

#### **Users**

- **Add Teacher**:  
  `POST /api/v1/teacher/`  
  Adds a new teacher to the system by inserting into both the users and teachers tables.

  **Request Body (JSON)**:
  ```json
  {
    "first_name": "james",
    "last_name": "foo",
    "phone_number": "09371327163",
    "sunday_entry_time": 830,
    "monday_entry_time": 800,
    "tuesday_entry_time":800,
    "wednesday_entry_time":800,
    "thursday_entry_time":800,
    "friday_entry_time":800,
    "saturday_entry_time":800
  }
  ```


  **response Body (JSON)**:
  ```json
  {
    "id": 1,
    "message": "Teacher created"
  }
  ```

- **Add Student**:  
  `POST /api/v1/student/`  
  Adds a new student to the system by inserting into both the users and students tables.

  **Request Body (JSON)**:
  ```json
  {
    "first_name": "bar",
    "last_name": "baz",
    "phone_number": "789456123",
    "required_entry_time": 900
  }
  ```

  **response Body (JSON)**:
  ```json
  {
    "id": 1,
    "message": "Student created"
  }
  ```

- **Get User by ID**:  
  `GET /api/v1/user/:id`  
  Retrieves user details by user ID.


  **response Body (JSON)**:
  ```json
  {
    "user_id": 12,
    "first_name": "1 Student",
    "last_name": "Lastname 1",
    "phone_number": "09108969774",
    "image_path": "",
    "finger_id": "",
    "is_biometric_active": false
  }
  ```

- **Get Teacher by ID**:  
  `GET /api/v1/teacher/:id`  
  Retrieves teacher details by teacher ID.
    ```json
  {
    "user_id": 1,
    "first_name": "0 Teacher",
    "last_name": "Lastname 0",
    "sunday_entry_time": 0,
    "monday_entry_time": 800,
    "tuesday_entry_time": 0,
    "wednesday_entry_time": 800,
    "thursday_entry_time": 800,
    "friday_entry_time": 800,
    "saturday_entry_time": 800
  }
  ```
  

- **Get Student by ID**:  
  `GET /api/v1/student/:id`  
  Retrieves student details by student ID.

  **response Body (JSON)**:
    ```json
  {
    "user_id": 12,
    "first_name": "1 Student",
    "last_name": "Lastname 1",
    "required_entry_time": 0
  }
  ```
- **Get User by Name**:  
  `GET /api/v1/user/name/:first_name/:last_name`  
  Retrieves user details by first and last name.


  **response Body (JSON)**:
    ```json
    {
        "user_id": 103,
        "first_name": "bar",
        "last_name": "baz",
        "phone_number": "789456123",
        "image_path": "",
        "is_biometric_active": false
    }
  ```

- **Update User**:  
  `PUT /api/v1/user/:id`  
  Updates a user’s details by user ID.

  **Request Body (JSON)**:
  ```json
  {
    "first_name": "UpdatedFirstName",
    "last_name": "UpdatedLastName",
    "phone_number": "09123456787",
    "image_path": "/images/updated.jpg",
  }
  ```

   **response Body (JSON)**:
    ```json
  {
    "id": 1,
    "message": "user updated"
  }
  ```

- **Update Student's Allowed Entry Time**:  
  `PUT /api/v1/student/:id`  
  Updates allowed entry time for a student by student ID.

  **Request Body (JSON)**:
  ```json
  {
    "required_entry_time": 900
  }
  ```

  **response Body (JSON)**:
    ```json
  {
    "id": 1,
    "message": "student allowed updated"
  }
  ```

- **Update Teacher's Allowed Entry Time**:  
  `PUT /api/v1/teacher/:id`  
  Updates allowed entry time for a teacher by teacher ID.

  **Request Body (JSON)**:
  ```json
  {
    "sunday_entry_time": 830,
    "monday_entry_time": 800,
    "tuesday_entry_time":800,
    "wednesday_entry_time":800,
    "thursday_entry_time":800,
    "friday_entry_time":800,
    "saturday_entry_time":800
  }
  ```
  **response Body (JSON)**:
    ```json
  {
    "id": 1,
    "message": "teacher allowed updated"
  }
  ```


- **Delete User**:  
  `DELETE /api/v1/user/:id`  
  Deletes a user from the system by user ID.

**response Body (JSON)**:
  ```json
  {
    "id": 1,
    "message": "user deleted"
  }
  ```

#### **Attendance**

- **Get Attendance by Date**:  
  `GET /api/v1/attendance/:date`  
  Retrieves attendance records for a specific date.
  
- **Get Attendance Between Dates**:  
  `GET /api/v1/attendance/range/:startDate/:endDate`  
  Retrieves attendance records between two dates.
  
- **Get Absent Users by Date**:  
  `GET /api/v1/attendance/absent/:date`  
  Retrieves users who were absent on a specific date.

#### **Admin**

- **Update Admin Password**:  
  `PUT /api/v1/admin/`  
  Updates the admin's password by username.

  **Request Body (JSON)**:
  ```json
  {
    "user_name": "admin",
    "password": "newPassword456"
  }
  ```

### Biometric Endpoints (`/biometric`)

These routes do not require authentication and handle biometric-based attendance.

- **Get Users Without Biometric Data**:  
  `GET /biometric/`  
  Retrieves users whose biometric authentication is inactive.

- **Get Users with Active Biometric Data**:  
  `GET /biometric/user`  
  Retrieves users whose biometric authentication is active.

- **Update User Biometric Data**:  
  `PUT /biometric/:id`  
  Updates a user's biometric data by user ID.

  **Request Body (JSON)**:
  ```json
  {
    "image_path": "home/downloads/fadsfasdf.jpg",
    "finger_id": "hashidforfinger"
  }
  ```

- **Post Attendance (Biometric Entry/Exit)**:  
  `POST /biometric/attendance/`  
  Handles attendance entry and exit using biometric authentication.

  **Request Body (JSON)**:
  ```json
  {
    "user_id": 1,
    "time": 1726336438
  }
  ```

