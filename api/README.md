# Go Jobs API

A small Go REST API for user signup, login, and customer management using Gin and SQLite.

## Architecture

```
Client
  ├─ POST /signup       -> create user
  ├─ POST /auth         -> login and receive JWT
  └─ /customer/*        -> protected routes using JWT token

Protected request flow:
  Client -> Authorization: Bearer <token> -> Gin middleware -> route handler -> SQLite DB
```

## Features

- `POST /signup` -> create a new user
- `POST /auth` -> authenticate existing users and return JWT
- `POST /customer` -> create a customer (authenticated)
- `GET /customer` -> list authenticated user's customers
- `GET /customer/:id` -> get a single customer by ID (authenticated)
- `PUT /customer/:id` -> update a customer by ID (authenticated)

## Run the app

```bash
cd /media/felipe/Lab2/www/projects/go-jobs/api
go run cmd/main.go
```

The server listens on port `3000`.

## Authentication

- Login with `POST /auth` and receive a JWT token.
- Send the token in the `Authorization` header for all `/customer` routes.

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

## Notes

- The database file is `go_job.db` in the project root.
- Authentication middleware validates the JWT and sets `userId` in the request context.
- Customer routes are scoped under `/customer` and require a valid token.
