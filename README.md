# Go Language Learning Repository

Welcome to this Go language learning repository! This project contains practical exercises and examples to help master the Go programming language from basics to more advanced concepts.

## 📁 Repository Structure

```
Go/
├── lab_01/          # Introduction to Go basics
│   ├── main.go      # Basic Hello World example
│   └── variables.go # Variables and data types introduction
└── README.md        # This file
```

## 🎯 Labs Overview

### Lab 01: Go Basics
- **Objective**: Introduction to Go syntax and basic programs
- **Topics Covered**:
  - Hello World program
  - Basic package structure
  - Introduction to fmt package for printing
  - Variables and data types

## 🚀 Getting Started

### Prerequisites

Before running these programs, ensure you have Go installed on your system:

1. Download Go from [https://golang.org/dl/](https://golang.org/dl/)
2. Follow the installation instructions for your operating system
3. Verify installation:
   ```bash
   go version
   ```

### Running Programs

To run any Go program in this repository:

```bash
# Run a single file
go run filename.go

# Example: Running from lab_01
cd lab_01
go run main.go
go run variables.go
```

Or from the root directory:
```bash
go run ./lab_01/main.go
go run ./lab_01/variables.go
```

## 📚 Learning Path

This repository follows a structured learning approach:

1. **Lab 01**: Fundamentals
   - Basic syntax
   - Package management
   - Variable declaration
   - Data types
   - Basic I/O operations

## 🛠️ Tools & Commands

Common Go commands used in this project:

```bash
go run      # Compile and run a Go program
go build    # Compile Go source code
go test     # Run tests
go fmt      # Format code
go vet      # Analyze code for potential bugs
```

## 💡 Tips for Learning Go

- Go is designed to be simple and readable
- Pay attention to error handling (Go's approach is different from many languages)
- Practice writing small programs to reinforce concepts
- Use `go fmt` to maintain consistent code style
- Refer to the [official Go documentation](https://golang.org/doc/) for detailed information

## 📖 Additional Resources

- [Official Go Documentation](https://golang.org/doc/)
- [A Tour of Go](https://tour.golang.org/)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Effective Go](https://golang.org/doc/effective_go)

## 🔄 Progression

As you progress through the labs, concepts build upon each other. Ensure you understand each section before moving to the next.

## 📝 Notes

- All programs use the `main` package and `main()` function as entry points
- The `fmt` package is used for formatted I/O operations
- Code follows Go conventions and best practices

---

**Happy Learning!** 🎉

For any questions or issues, refer to the official Go documentation or community forums.
