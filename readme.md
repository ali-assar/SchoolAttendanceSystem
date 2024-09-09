# School Attendance System API

This project is an API for managing users and attendance in a school system, built using Go and Fiber.

## Getting Started

Follow these instructions to set up and test the API locally.

### Prerequisites

1. **Go**: Make sure Go is installed on your machine. You can download it [here](https://golang.org/dl/).
2. **SQLite**: The database used is SQLite. Ensure you have SQLite installed. Download it [here](https://www.sqlite.org/download.html).
3. **Postman or Insomnia**: For testing the API, you can use Postman or any other API client tool.

### Installing

1. **Clone the repository**:

    ```bash
    git clone https://github.com/Ali-Assar/SchoolAttendanceSystem.git
    cd SchoolAttendanceSystem
    ```

2. **Install dependencies**:

    You will need to install Go dependencies, which are managed with `go mod`.

    ```bash
    go mod download
    ```

3. **Create `.env` file**:

    Create a `.env` file in the root of the project and include the following configuration:

    ```env
    PORT=3000
    DB_PATH=./database.db
    ```

   You can change the `PORT` if necessary, and the database file path `DB_PATH` should point to the location of your SQLite database.

4. **Run the application**:

    Run the following command to start the server:

    ```bash
    go run main.go
    ```

   This will start the server on `http://localhost:3000` (or the port specified in the `.env` file).

5. **Database Initialization**:

    The database will be automatically initialized on the first run. Tables for users and attendance will be created if they do not exist.

### API Documentation

The following endpoints are available:

#### User Routes

- **POST `/api/v1/user/`**: Create a new user
- **GET `/api/v1/user/:id`**: Get user by ID
- **GET `/api/v1/user/`**: Get all users
- **PUT `/api/v1/user/:id`**: Update user by ID
- **GET `/api/v1/user/phone/:phone`**: Get user by phone number
- **GET `/api/v1/user/name/:first_name/:last_name`**: Get user by name
- **DELETE `/api/v1/user/:id`**: Delete user by ID

#### Attendance Routes

- **POST `/api/v1/attendance/`**: Create a new attendance record
- **GET `/api/v1/attendance/:user_id/:date`**: Get attendance by user ID and date
- **GET `/api/v1/attendances/:date`**: Get all attendance records for a specific date
- **PUT `/api/v1/attendance/`**: Update attendance by ID
- **DELETE `/api/v1/attendance/:attendance_id`**: Delete attendance by ID

### Testing the API

To test the API, you can use Postman, curl, or any other API testing tool.

#### Example Requests

##### Create a User (POST `/api/v1/user/`)

```json
POST http://localhost:3000/api/v1/user/
Content-Type: application/json

{
  "first_name": "John",
  "last_name": "Doe",
  "phone": "1234567890"
}
```

##### Get User by ID (GET `/api/v1/user/:id`)

```json
GET http://localhost:3000/api/v1/user/1
```

##### Create Attendance (POST `/api/v1/attendance/`)

```json
POST http://localhost:3000/api/v1/attendance/
Content-Type: application/json

{
  "user_id": 1,
  "entry_time": 1693516800,  // Unix timestamp
  "exit_time": 1693520400    // Unix timestamp
}
```

##### Get Attendance by User ID and Date (GET `/api/v1/attendance/:user_id/:date`)

```json
GET http://localhost:3000/api/v1/attendance/1/1693516800
```

#### Error Handling

- **400 Bad Request**: Invalid data provided in the request body or URL parameters.
- **404 Not Found**: Requested resource does not exist.
- **500 Internal Server Error**: An unexpected error occurred on the server.

### Running Tests

You can run unit tests using the following command:

```bash
go test ./...
```

Ensure that the test dependencies like `testify` are installed.
