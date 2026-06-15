# Go Learning Project

This project is a comprehensive introduction to the Go programming language (Golang), covering fundamental to intermediate concepts through hands-on tutorial files. Each tutorial builds on previous knowledge and explores a specific topic.

## Topics & File Paths

| # | Topic | File |
|---|-------|------|
| 1 | **Hello World** – Basic program structure, `package main`, `import`, and `fmt.Println()` | [tutorial_1/main.go](tutorial_1/main.go) |
| 2 | **Variables & Data Types** – Declaring variables (`int`, `float64`, `string`, `rune`, `bool`), zero values, and formatted printing | [tutorial_2/main.go](tutorial_2/main.go) |
| 3 | **Functions & Error Handling** – Defining functions with parameters and return types, multiple return values, error handling with `errors.New()`, `if/else`, and `switch` statements | [tutorial_3/main.go](tutorial_3/main.go) |
| 4 | **Collections & Loops** – Arrays (fixed-length), slices (dynamic-length, `append`, `make`, capacity preallocation performance), maps (key-value store, iteration, deletion), and loop styles (`for` with condition, `for` with index, infinite loop with `break`) | [tutorial_4/main.go](tutorial_4/main.go) |
| 5 | **Strings & Runes** – Strings as byte arrays, UTF-8 indexing, rune literals, string concatenation vs `strings.Builder` performance | [tutorial_5/main.go](tutorial_5/main.go) |
| 6 | **Structs** – Defining structs, creating instances (with/without field names), reassigning fields, nested structs, embedded structs (field promotion), and anonymous structs | [tutorial_6/main.go](tutorial_6/main.go) |
| 6a | **Interfaces** – Defining interfaces, implementing interfaces on structs, using interfaces as function parameters for polymorphism | [tutorial_6a/main.go](tutorial_6a/main.go) |
| 7 | **Pointers** – Pointer declaration (`*T`), referencing (`&`), dereferencing (`*`), heap allocation with `new()`, and slice reference semantics | [tutorial_7/main.go](tutorial_7/main.go) |
| 7a | **Pointers & Functions** – Passing by value vs passing by pointer, modifying data via pointers in functions | [tutorial_7a/main.go](tutorial_7a/main.go) |
| 8 | **Goroutines & Concurrency** – Concurrent execution with `go` keyword, `sync.WaitGroup` for synchronization, `sync.Mutex` for preventing race conditions, performance comparison (without goroutine, with goroutine, with race conditions, with mutex locks) | [tutorial_8/main.go](tutorial_8/main.go) |
| 9 | **Channels** – Channel creation (`make(chan T)`), sending/receiving values (`<-`), closing channels (`close`), iterating with `range`, buffered channels vs unbuffered channels | [tutorial_9/main.go](tutorial_9/main.go) |
| 9a | **Select Statement** – Using `select` with multiple channels to race between goroutines and send the first-available notification | [tutorial_9a/main.go](tutorial_9a/main.go) |
| 10 | **Generics (Functions)** – Writing type-specific functions vs generic functions with type constraints (`[T int \| float32 \| float64]`) | [tutorial_10/main.go](tutorial_10/main.go) |
| 10a | **Generics + JSON** – Loading JSON files into typed slices using a generic `loadJSON[T]()` function with `encoding/json` and `os` | [tutorial_10a/main.go](tutorial_10a/main.go) |
| 10b | **Generic Structs** – Defining structs with generic type fields (`desktop[T windowsOS \| macOS]`) and instantiating with concrete types | [tutorial_10b/main.go](tutorial_10b/main.go) |
| 11 | **HTTP API (REST)** – Building a REST API with `chi` router, custom middleware (authorization), handler pattern, mock database with interface, request/response models, and structured logging via `logrus` | [tutorial_11/](tutorial_11/) |

## Project Structure

```
go_learning/
├── go.mod                          # Go module definition
├── readme.md                       # This file
├── tutorial_1/                     # Hello World
│   └── main.go
├── tutorial_2/                     # Variables & Data Types
│   └── main.go
├── tutorial_3/                     # Functions & Error Handling
│   └── main.go
├── tutorial_4/                     # Collections & Loops
│   └── main.go
├── tutorial_5/                     # Strings & Runes
│   └── main.go
├── tutorial_6/                     # Structs
│   └── main.go
├── tutorial_6a/                    # Interfaces
│   └── main.go
├── tutorial_7/                     # Pointers
│   └── main.go
├── tutorial_7a/                    # Pointers & Functions
│   └── main.go
├── tutorial_8/                     # Goroutines & Concurrency
│   └── main.go
├── tutorial_9/                     # Channels
│   └── main.go
├── tutorial_9a/                    # Select Statement
│   └── main.go
├── tutorial_10/                    # Generics (Functions)
│   └── main.go
├── tutorial_10a/                   # Generics + JSON
│   ├── main.go
│   ├── contactInfos.json
│   └── purchaseInfos.json
├── tutorial_10b/                   # Generic Structs
│   └── main.go
└── tutorial_11/                    # HTTP API (REST)
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── api/
    │   └── api.go                  # Request/response models & error helpers
    ├── cmd/
    │   └── api/
    │       └── main.go             # Entry point: chi router, logrus, HTTP listen
    └── internal/
        ├── handlers/
        │   ├── api.go              # Route definitions with chi
        │   └── get_coin_balance.go # GET /account/coins handler
        ├── middleware/
        │   └── authorization.go    # Auth middleware (token + username validation)
        └── tools/
            ├── database.go         # DatabaseInterface definition & constructor
            └── mockdb.go           # In-memory mock implementation
```

## Key Concepts Covered

| Concept | Tutorial(s) |
|---------|-------------|
| Package & import system | 1 |
| Variable declaration & zero values | 2 |
| Basic data types (int, float, string, rune, bool) | 2, 5 |
| Functions (parameters, return values, multiple returns) | 3 |
| Error handling & custom errors | 3 |
| Control flow (if/else, switch) | 3 |
| Arrays (fixed-length) | 4 |
| Slices (dynamic, append, make, capacity) | 4 |
| Maps (CRUD, iteration, existence check) | 4 |
| For loops (condition, index, infinite + break) | 4 |
| Strings (byte arrays, UTF-8, runes, concatenation, Builder) | 5 |
| Structs (fields, nesting, embedding, anonymous) | 6 |
| Interfaces (definition, implementation, polymorphism) | 6a |
| Pointers (declaration, &, *, new, slice semantics) | 7 |
| Pass-by-value vs pass-by-pointer | 7a |
| Goroutines (go keyword, async execution) | 8 |
| WaitGroup (goroutine synchronization) | 8 |
| Mutex (race condition prevention) | 8 |
| Channels (unbuffered, buffered, close, range) | 9 |
| Select statement (multi-channel race) | 9a |
| Generics (type constraints on functions) | 10 |
| Generics + JSON deserialization | 10a |
| Generics (type constraints on structs) | 10b |
| HTTP server with chi router | 11 |
| Route grouping & middleware | 11 |
| Handler pattern (http.ResponseWriter, *http.Request) | 11 |
| Request query parsing (gorilla/schema) | 11 |
| Custom middleware (authorization) | 11 |
| Interface-based database abstraction | 11 |
| Mock database pattern | 11 |
| Structured logging (logrus) | 11 |
| JSON request/response encoding | 11 |

## How to Run

```bash
# Run a single-file tutorial (1–10b)
go run tutorial_1/main.go
go run tutorial_2/main.go
# ... etc.

# Run the REST API (tutorial_11)
cd tutorial_11 && go run ./cmd/api
```

## Prerequisites

- [Go](https://go.dev/dl/) 1.25+ (the project uses Go modules)

## Learning Path

The tutorials follow this recommended order:

1. **Basics** – Tutorials 1–3 (Hello World, Variables, Functions)
2. **Data Structures** – Tutorials 4–5 (Collections, Strings)
3. **OOP-Style Constructs** – Tutorials 6–6a (Structs, Interfaces)
4. **Memory & References** – Tutorials 7–7a (Pointers)
5. **Concurrency** – Tutorials 8–9a (Goroutines, Channels, Select)
6. **Generics** – Tutorials 10–10b (Functions, JSON, Structs)
7. **Real-World Application** – Tutorial 11 (REST API with router, middleware, handlers, mock DB)