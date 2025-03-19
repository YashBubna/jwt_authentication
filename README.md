# JWT Authentication

A simple JWT-based authentication system implemented using Go.

## Features
- User registration with hashed passwords
- User login with JWT token generation
- Middleware for protected routes
- Token validation and user authentication

## Tech Stack
- **Golang** – Backend API
- **Gin** – HTTP framework for Go
- **JWT-Go** – JWT authentication library
- **MongoDB** – Database for storing user data

## Installation

### Prerequisites
- Go installed on your machine
- MongoDB instance running

### Steps
1. Clone the repository:
   ```sh
   git clone https://github.com/YashBubna/jwt_authentication.git
   cd jwt_authentication
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Set up environment variables in a `.env` file:
   ```env
   MONGO_URI=mongodb://localhost:27017
   JWT_SECRET=your_secret_key
   ```
4. Run the application:
   ```sh
   go run main.go
   ```

## API Endpoints

### 1. Register a User
**POST** `/register`
```json
{
  "username": "testuser",
  "password": "securepassword"
}
```

### 2. Login
**POST** `/login`
```json
{
  "username": "testuser",
  "password": "securepassword"
}
```
Response:
```json
{
  "token": "your_jwt_token"
}
```

### 3. Protected Route (Requires JWT)
**GET** `/protected`
- Include `Authorization: Bearer <token>` in the headers.

## License
This project is licensed under the MIT License.
