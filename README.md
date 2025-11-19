# User CRUD API with JWT Authentication

A RESTful API built with Go for user management with JWT-based authentication and authorization.

## Project Structure

```
CRUD_NavaneethN/
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── handler/
│   │   └── user_handler.go     # HTTP request handlers
│   ├── service/
│   │   └── user_service.go     # Business logic layer
│   ├── repository/
│   │   └── user_repository.go  # Data access layer
│   ├── model/
│   │   └── user_model.go       # Data models
│   ├── middleware/
│   │   └── auth_middleware.go  # JWT authentication middleware
│   ├── jwt/
│   │   └── jwt.go              # JWT token generation/validation
│   └── router/
│       └── user_router.go      # Route definitions
├── go.mod
├── go.sum
└── README.md
```

## API Endpoints

### Public Endpoints (No Authentication Required)

#### 1. Register User

Create a new user account.

**Endpoint:** `POST /api/users`

**Request Body:**

```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

**Response (Success - 200):**

```json
{
  "_id": "507f1f77bcf86cd799439011",
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

---

#### 2. Login

Authenticate user and receive JWT token.

**Endpoint:** `POST /api/login`

**Request Body:**

```json
{
  "email": "john@example.com",
  "password": "password123"
}
```

**Response (Success - 200):**

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

---

### Protected Endpoints (Authentication Required)

All protected endpoints require a JWT token in the Authorization header:

```
Authorization: Bearer <your-jwt-token>
```

#### 3. Get User by ID

Retrieve a specific user's profile. **Users can only view their own profile.**

**Endpoint:** `GET /api/users/{id}`

**Headers:**

```
Authorization: Bearer <token>
Content-Type: application/json
```

**URL Parameters:**

- `id` - User ID (MongoDB ObjectID)

**Response (Success - 200):**

```json
{
  "_id": "507f1f77bcf86cd799439011",
  "name": "John Doe",
  "email": "john@example.com"
}
```

---

#### 4. Update User

Update user information. **Users can only update their own profile.**

**Endpoint:** `PUT /api/users/{id}`

**Headers:**

```
Authorization: Bearer <token>
Content-Type: application/json
```

**URL Parameters:**

- `id` - User ID (MongoDB ObjectID)

**Request Body:**

```json
{
  "name": "John Updated",
  "email": "john.updated@example.com",
  "password": "newpassword123"
}
```

**Response (Success - 200):**

```json
"User updated successfully"
```

---

#### 5. Delete User

Delete a user account. **Users can only delete their own account.**

**Endpoint:** `DELETE /api/users/{id}`

**Headers:**

```
Authorization: Bearer <token>
Content-Type: application/json
```

**URL Parameters:**

- `id` - User ID (MongoDB ObjectID)

**Response (Success - 200):**

```json
"User deleted successfully"
```

---

#### 6. Delete All Users

Delete all users from the database.

**Endpoint:** `DELETE /api/users/all`

**Headers:**

```
Authorization: Bearer <token>
Content-Type: application/json
```

**Response (Success - 200):**

```json
"All users deleted, count: 5"
```

---

## Authentication Flow

1. **Register** a new user account via `POST /api/users`
2. **Login** with credentials via `POST /api/login` to receive a JWT token
3. **Use the token** in the Authorization header for protected endpoints:
   ```
   Authorization: Bearer <your-jwt-token>
   ```
4. Token expires after **24 hours**

## Authorization Rules

- Users can only **view** their own profile (`GET /api/users/{id}`)
- Users can only **update** their own profile (`PUT /api/users/{id}`)
- Users can only **delete** their own account (`DELETE /api/users/{id}`)
- Attempting to access/modify another user's data returns **403 Forbidden**
