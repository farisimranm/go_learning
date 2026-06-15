# Go Learning Project

A Golang learning project by following this [Youtube tutorial](https://[example.com](https://youtu.be/8uiZC0l4Ajw)) from Alex Mux.

## Topics & File Paths

| # | Topic | File |
| --- | ------- | ------ |
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

```plaintext
go_learning/
├── go.mod                          # Go module definition
├── README.md                       # This file
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

- Of course you need to install Go first
