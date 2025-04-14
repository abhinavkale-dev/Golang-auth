# Golang Authentication System

A modern authentication system built with Go, Prisma, and React.

## Overview

This project provides a full-stack authentication system with:
- Go backend with Gin web framework
- Prisma ORM for database operations
- PostgreSQL database for data storage
- React frontend with TypeScript

## Features

- User registration (signup)
- User authentication (signin)
- Session management
- Secure password handling

## Project Structure

```
.
├── server/           # Go backend
│   ├── controllers/  # Request handlers
│   ├── models/       # Data models
│   ├── prisma/       # Prisma ORM
│   └── utils/        # Helper functions
└── client/           # React frontend
    ├── public/       # Static files
    └── src/          # React components and logic
```

## Prerequisites

- Go (1.18+)
- Node.js (18+)
- PostgreSQL
- Git

## Setup & Installation

### Backend

1. Navigate to the server directory:
   ```
   cd server
   ```

2. Install Go dependencies:
   ```
   go mod download
   ```

3. Set up your PostgreSQL database and update the connection string in `.env`

4. Generate Prisma client:
   ```
   go run github.com/steebchen/prisma-client-go generate
   ```

5. Run database migrations:
   ```
   go run github.com/steebchen/prisma-client-go migrate dev
   ```

6. Start the backend server:
   ```
   go run main.go
   ```

### Frontend

1. Navigate to the client directory:
   ```
   cd client
   ```

2. Install Node.js dependencies:
   ```
   npm install
   ```

3. Start the development server:
   ```
   npm run dev
   ```

## API Endpoints

- `POST /signup` - Create a new user account
- `POST /signin` - Authenticate a user
- `POST /signout` - End a user session

## Development

- Backend runs on: http://localhost:8080
- Frontend runs on: http://localhost:5173

