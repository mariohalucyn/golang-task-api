Simple Todo application built with Go, Gorilla Mux, GORM, and PostgreSQL. The application allows users to sign up, log in, create, update, and delete todo items. It also includes email verification and JWT-based authentication.

## Features

- User registration and login
- Email verification
- JWT-based authentication
- Create, update, and delete todo items
- List all todos for a user

## Prerequisites

- Go 1.23.0 or later
- PostgreSQL
- Docker (optional, for running PostgreSQL in a container)

## Setup

1. Clone the repository:

   ```sh
   git clone https://github.com/mariohalucyn/todo-app.git
   cd todo-app
   ```

2. Create a `.env` file in the root directory and add the following environment variables:

   ```env
   DSN=your_postgres_dsn
   FROM_EMAIL=your_email@example.com
   POSTMARKAPP_USERNAME=your_postmark_username
   POSTMARKAPP_PASSWORD=your_postmark_password
   FRONTEND_ADDRESS=http://localhost:3000
   ```

3. Generate ECDSA keys for JWT signing:

   ```sh
   openssl ecparam -name prime256v1 -genkey -outform der -out ec-priv-key.pem
   openssl ec -in ec-priv-key.pem -pubout > ec-pub-key.pem
   ```

4. Run PostgreSQL using Docker (optional):

   ```sh
   docker-compose up -d
   ```

5. Install dependencies:

   ```sh
   go mod tidy
   ```

6. Build and run the application:

   ```sh
   go build -o todo-app
   ./todo-app
   ```

## API Endpoints

- `POST /api/signup`: Register a new user
- `POST /api/login`: Log in a user
- `GET /api/verify`: Verify user email
- `PUT /api/update-user`: Update user details
- `GET /api/authorization`: Check user authorization
- `GET /api/logout`: Log out a user
- `POST /api/create-todo`: Create a new todo item
- `GET /api/get-todos`: Get all todos for a user
- `PUT /api/update-todo/{id}`: Update a todo item
- `DELETE /api/delete-todo/{id}`: Delete a todo item
