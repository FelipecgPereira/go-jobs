# Go Jobs API

A small Go REST API for user signup, login, customer/project management and B2B payment summaries using Gin and SQLite.

## Architecture

```
Client
  ├─ POST /signup       -> create user
  ├─ POST /auth         -> login and receive JWT
  ├─ /customer/*        -> customer management (protected)
  ├─ /project/*         -> project management (protected)
  └─ /b2b/*             -> b2b operations (protected)
```

Protected request flow:

Client -> Authorization: Bearer <token> -> Gin middleware -> route handler -> SQLite DB

## Features

- User signup and JWT authentication
- Customer CRUD for authenticated users
- Project CRUD for authenticated users
- B2B creation/update and payment summary by status/date range

## Run the app

```bash
cd /media/felipe/Lab2/www/projects/go-jobs/api
go run cmd/main.go
```

The server listens on port `3000`.

## Authentication

- Login with `POST /auth` and receive a JWT token.
- Send the token in the `Authorization` header for all `/customer`, `/project` and `/b2b` routes.

Example header:

```http
Authorization: Bearer <token>
```

## Endpoints

### Signup

`POST /signup`

Request body:

```json
{
  "name": "Alice",
  "email": "alice@example.com",
  "password": "secret"
}
```

Response:

```json
{
  "message": "User created successfully"
}
```

### Auth

`POST /auth`

Request body:

```json
{
  "email": "alice@example.com",
  "password": "secret"
}
```

Response:

```json
{
  "message": "Login successful",
  "token": "..."
}
```

## Customer routes

### Create customer

`POST /customer`

Headers:

```http
Authorization: Bearer <token>
```

Request body:

```json
{
  "name": "Acme Ltd",
  "email": "contact@acme.com",
  "phone": "123-456-7890"
}
```

Response:

```json
{
  "message": "Customer created successfully"
}
```

### List customers

`GET /customer`

Headers:

```http
Authorization: Bearer <token>
```

Response: list of customers owned by the authenticated user.

### Get customer by ID

`GET /customer/:id`

Headers:

```http
Authorization: Bearer <token>
```

Response: single customer object.

### Update customer

`PUT /customer/:id`

Headers:

```http
Authorization: Bearer <token>
```

Request body:

```json
{
  "name": "Acme Ltd Updated",
  "email": "contact@acme.com",
  "phone": "987-654-3210"
}
```

Response:

```json
{
  "message": "Customer updated successfully"
}
```

## Project routes

### Create project

`POST /project/`

Headers:

```http
Authorization: Bearer <token>
```

Request body:

```json
{
  "name": "Photo shape",
  "price": 2500.99,
  "startDate": "2026-07-01T18:00:00Z",
  "endDate": "2026-07-31T18:00:00Z"
}
```

Response:

```json
{
  "message": "Project created successfully"
}
```

### List projects

`GET /project`

Headers:

```http
Authorization: Bearer <token>
```

### Get project by ID

`GET /project/:id`

Headers:

```http
Authorization: Bearer <token>
```

### Update project

`PUT /project/:id`

Headers:

```http
Authorization: Bearer <token>
```

Request body:

```json
{
  "name": "Photo shape revised",
  "price": 2900.00,
  "startDate": "2026-07-01T18:00:00Z",
  "endDate": "2026-08-01T18:00:00Z"
}
```

Response:

```json
{
  "message": "Project updated successfully"
}
```

## B2B routes

### Create B2B

`POST /b2b/`

Headers:

```http
Authorization: Bearer <token>
```

Request body:

```json
{
  "customerId": 10001,
  "projectId": 10001,
  "status": "active"
}
```

Response:

```json
{
  "message": "b2b created successfully"
}
```

### Update B2B

`PUT /b2b/:id`

Headers:

```http
Authorization: Bearer <token>
```

Request body: any valid B2B fields to update.

Response:

```json
{
  "message": "B2B updated successfully"
}
```

### Sum payments

`GET /b2b/sum?status=active&start=2026-07-01T00:00:00&end=2026-07-28T23:59:59`

Headers:

```http
Authorization: Bearer <token>
```

Query params:

- `status` - B2B status value
- `start` - start datetime in `YYYY-MM-DDTHH:MM:SS`
- `end` - end datetime in `YYYY-MM-DDTHH:MM:SS`

Response:

```json
{
  "total": 12345.67
}
```

## Notes

- The database file is `go_job.db` in the project root.
- Authentication middleware validates the JWT and sets `userId` in the request context.
- All `/customer`, `/project` and `/b2b` routes require a valid token.
- Project date fields are stored as SQLite `DATETIME` values.
