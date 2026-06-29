# Go Jobs

A full-stack application with a Go API and an Angular frontend for managing customers, projects, and B2B records.

## Project structure

- api/: Go backend using Gin and SQLite
- app/: Angular frontend

## Requirements

For Docker-based execution:
- Docker
- Docker Compose

For local development without Docker:
- Go
- Node.js
- npm

## Run with Docker Compose

From the project root, run:

```bash
docker compose up --build
```

After the containers start, open:

- Frontend: http://localhost:4200
- API: http://localhost:3000

To stop the services, press Ctrl+C in the terminal where Docker Compose is running, or run:

```bash
docker compose down
```

## Main API endpoints

- POST /signup
- POST /auth
- GET/POST/PUT /customer
- GET/POST/PUT /project
- POST/PUT /b2b
- GET /b2b/sum

## Local development

### Backend

```bash
cd api
go run ./cmd
```

### Frontend

```bash
cd app
npm install
npm start -- --host 0.0.0.0 --port 4200
```

If you want to build the Angular app for production, run:

```bash
cd app
npm run build
```
