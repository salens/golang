# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Structure

This is a Go learning project demonstrating core Go language concepts including:

- **Entry point**: `main.go` - Main application demonstrating Go fundamentals
- **Function examples**: `functionExamples.go` - Helper functions and examples
- **Data models**: `Model/myStructs.go` - Person struct definition
- **Module**: `github.com/salens/golang-project` (Go 1.25.0)

### Architecture

- Single-module Go project with a simple package structure
- Main package contains the primary executable code
- Model package contains data structures (Person struct)
- Demonstrates Go concurrency patterns with goroutines and channels

## Key Components

### Core Functionality
- **Person struct**: Basic data structure in `Model/myStructs.go:5-12` with ID, Name, and Age fields
- **Function examples**: Multiple utility functions in `functionExamples.go` including:
  - `ReturnPerson()`: Creates Person instances
  - `GetUserInfo()`: Returns formatted greetings and age calculations
  - `ReturnTwoInts()`: Example of multiple return values

### Concurrency Examples
- Goroutines with WaitGroup synchronization
- Channel communication patterns (unbuffered, buffered, directional)
- Anonymous functions with concurrent execution

## Development Commands

### Build and Run
```bash
# Build the project
go build -o myprogram .

# Run the built executable
./myprogram

# Build and run in one command
go run .
```

### Module Management
```bash
# Initialize module (already done)
go mod init github.com/salens/golang-project

# Download dependencies
go mod tidy

# Verify dependencies
go mod verify
```

### Code Quality
```bash
# Format code
go fmt ./...

# Vet code for issues
go vet ./...
```

## Code Patterns

When working with this codebase:

1. **Package imports**: Uses the custom module path `github.com/salens/golang-project/model`
2. **Struct initialization**: Person structs use field names explicitly
3. **Concurrency**: Heavy use of goroutines with proper synchronization using WaitGroups
4. **Channel patterns**: Demonstrates unbuffered, buffered, and directional channels
5. **Anonymous functions**: Extensive use for concurrent operations

## File Organization

- `main.go`: Primary execution logic with comprehensive Go feature demonstrations
- `functionExamples.go`: Utility functions and examples
- `Model/myStructs.go`: Data structure definitions
- `myprogram`: Compiled executable (generated)