# Library Management System

A Go-based library management system implemented to learn and practice layered architecture design principles.

## Project Overview

This project implements a basic library management system using Go, focusing on clean architecture principles and separation of concerns. It uses in-memory storage for simplicity while maintaining a structure that could be easily extended to use persistent storage.

## Architecture

The project follows a layered architecture pattern with the following components:

### Project Structure
```
├── cmd/
│   └── main.go           # Application entry point
├── internal/
│   ├── domain/          # Core business entities and rules
│   │   ├── books.go
│   │   ├── loans.go
│   │   └── members.go
│   ├── repository/      # Data access layer
│   │   ├── book_loan_repo.go
│   │   ├── book_repo.go
│   │   ├── member_repo.go
│   │   └── repo.go
│   └── service/        # Business logic layer
│       ├── book_service.go
│       ├── loan_service.go
│       └── member_service.go
└── pkg/
    └── generator.go    # Utility functions
```

### Layers
1. **Domain Layer**: Contains the core business entities and rules
   - Books, Members, and Loans entities
   - Business rules and validation logic

2. **Repository Layer**: Handles data storage and retrieval
   - In-memory implementation using Go maps
   - Abstracts data storage operations
   - Implements CRUD operations for each entity

3. **Service Layer**: Implements business logic
   - Coordinates between different repositories
   - Handles complex operations involving multiple entities
   - Enforces business rules

## Features

- **Book Management**
  - Create, read, update, and delete books
  - Track book metadata (title, author, published year, genre)

- **Member Management**
  - Member registration and profile management

- **Loan Management**
  - Book borrowing and returning functionality
  - Track book lending status

## Learning Objectives

- Understanding and implementing layered architecture
- Separation of concerns
- Domain-driven design principles
- Clean code practices
- Go programming patterns and best practices

## Future Improvements

- Implement persistent storage (e.g., SQLite, PostgreSQL)
- Add authentication and authorization
- Implement fine calculation system
- Add API/CLI interface
- Add unit tests and integration tests