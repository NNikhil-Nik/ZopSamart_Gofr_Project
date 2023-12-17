# ZopSamart_Gofr_Project
# Library Management System

A Library Management System (LMS) is a software application designed to automate and streamline the operations of a library. It serves as a centralized platform to manage various library processes, making it easier for librarians and library patrons to access, organize, and maintain library resources.

# To use GoFr:
```bash
go get gofr.dev
```
# UML Diagrams
- Use Case Diagram
  ![diagram](https://github.com/NNikhil-Nik/ZopSamart_Gofr_Project/assets/71598429/2519d34b-af76-4847-9676-a95e2a465478)

## API Calls

## 1. Adding Books(POST) 
![image](https://github.com/NNikhil-Nik/ZopSamart_Gofr_Project/assets/71598429/b2326185-9454-470a-8e9b-251d02ff676b)

## 2. Get All Books(GET) 
  ![image](https://github.com/NNikhil-Nik/ZopSamart_Gofr_Project/assets/71598429/0de4556c-0980-454c-a944-b5bb15eb3f8b)

## 3. Update Quantity of Books(PUT)
![image](https://github.com/NNikhil-Nik/ZopSamart_Gofr_Project/assets/71598429/1222966f-9b52-467c-b296-c9277876fc83)

# Library Management System in Go (Golang)

## Overview

This project is a library management system implemented in Go (Golang) framework GoFr using a MySQL database. It provides functionality for managing books and users, as well as establishing relationships between them.

## Project Structure

The project is organized into several packages:

- **models:** Defines data structures, including `Book` and `User`.
- **db:** Handles database operations, including CRUD operations for books and users.
- **handlers:** Implements HTTP request handlers for various operations.

## Functionality

### Books Management

- **GetBooks:** Retrieve a list of all books.
- **GetBook/{id}:** Retrieve a specific book by ID.
- **AddBook:** Add a new book to the library.
- **UpdateBook/{id}:** Update information for an existing book.
- **DeleteBook/{id}:** Delete a book from the library.

### Users Management

- **GetUsers:** Retrieve a list of all users.
- **GetUser/{id}:** Retrieve a specific user by ID.
- **AddUser:** Add a new user.
- **UpdateUser/{id}:** Update information for an existing user.
- **DeleteUser/{id}:** Delete a user.

### Book-User Relationship

Users have a list of associated books, forming a one-to-many relationship between users and books.

## API Endpoints

### Books API

- `GET /books`: Retrieve a list of all books.
- `GET /books/{id}`: Retrieve a specific book by ID.
- `POST /books`: Add a new book to the library.
- `PUT /books/{id}`: Update information for an existing book.
- `DELETE /books/{id}`: Delete a book from the library.

### Users API

- `GET /users`: Retrieve a list of all users.
- `GET /users/{id}`: Retrieve a specific user by ID.
- `POST /users`: Add a new user.
- `PUT /users/{id}`: Update information for an existing user.
- `DELETE /users/{id}`: Delete a user.

## Getting Started

To run the project locally, follow these steps:

1. Install Go and MySQL.
2. Set up the database with the required tables.
3. Clone the repository.
4. Configure database connection details in the `config.go` file.
5. Run the application using `go run main.go`.

Feel free to explore and contribute to the project!

