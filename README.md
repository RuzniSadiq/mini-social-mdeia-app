# Mini Social Media App

A robust and scalable mini social media application built using the **Gin framework** for Go. This app supports features like posts, comments, likes, and user authentication, all encapsulated in a modular structure with reusable generic services and controllers.

---

## Features

- User Authentication (JWT-based)
- CRUD operations for:
  - Posts
  - Comments
  - Likes
- Modular architecture with generic controllers and services
- Middleware for authentication and ownership checks
- Dependency injection for better testability and flexibility
- Database integration with GORM
- Preload support for related entities
- Comprehensive error handling

---

## API Endpoints

### Authentication
- **POST /register**: Register a new user  
- **POST /login**: Authenticate a user  

### Posts
- **POST /posts**: Create a new post  
- **GET /posts**: List all posts  
- **GET /posts/:id**: Retrieve a specific post  
- **PUT /posts/:id**: Update a post (ownership check)  
- **DELETE /posts/:id**: Delete a post (ownership check)  

### Comments
- **POST /comments**: Create a new comment  
- **GET /comments**: List all comments  
- **GET /comments/:id**: Retrieve a specific comment  
- **PUT /comments/:id**: Edit a specific comment (ownership check) 
- **DELETE /comments/:id**: Delete a comment (ownership check)  

### Likes
- **POST /likes**: Like a post or comment  
- **GET /likes**: List all likes
- **GET /likes/:id**: Retrieve a specific like 
- **DELETE /likes/:id**: Remove a like (ownership check)  

---

## Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/) (v1.18 or later)
- [PostgreSQL](https://www.postgresql.org/) or your preferred database
- [Git](https://git-scm.com/)
- [Docker](https://www.docker.com/) (Optional, for containerized deployment)

---

## Installation

### 1. Clone the Repository

```bash
git clone https://github.com/<your-github-username>/mini-social-media.git
cd mini-social-media
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Set up Postgres DB
Run postgres DB server

### 3. Set Up Environment Variables
Create a .env file in the root directory with the following keys:
DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=
DB_NAME=
JWT_SECRET=
SERVER_HOST=
SERVER_PORT=

### 4. Run Database Migrations
go run cmd/migrate.go


### 5. Start the Server:
go run .

---

## Possible Improvements

### Unit Tests
Implement unit tests in order to test API endpoints and application flow.

### Pagination
Implement pagination for efficient data retrieval, allowing users to request a specific subset of results per page to enhance performance and usability.

### Sorting
Add sorting functionality to allow users to organize results based on different fields (e.g., alphabetically, numerically, by date) for easier data analysis and retrieval.

### Filtering
Enable filtering options to let users narrow down data by specific attributes, helping them find relevant information quickly.

### Aggregation
Implement aggregation techniques for summarizing and analyzing data, such as counting, averaging, or grouping, to provide insightful, high-level overviews of data.

### Validation
Introduce input validation to ensure that data entered by users meets the necessary format and constraints, improving data integrity and application stability.

### Swagger Setup
Set up Swagger for automatic generation of interactive API documentation, making it easier for developers to understand and use the available endpoints.

### Test Dockerfile
Ensure that the Dockerfile is properly set up for testing the application. This should include all necessary dependencies, environment variables, and commands to build and run the application in a containerized environment.
