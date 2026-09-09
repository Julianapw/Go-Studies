# Go-Studies

A study repository for the **Go (Golang)** programming language, built around the **"8-Week Structured Go Mentorship"** track from the [Practice Golang for Beginners](https://practice-golang-for-beginners.github.io/) project.

## About

This repository gathers the exercises, solutions, and mini-projects developed throughout an 8-week structured mentorship, guiding the learner from the fundamentals of Go all the way to building a production-ready backend service.

## Objective

To consolidate the core concepts of Go in practice — syntax, interfaces, concurrency, context handling, HTTP API design, and engineering best practices — through weekly exercises and applied projects.

## Concept Roadmap (8 Weeks)

| Week | Topic | Key Concepts |
|---|---|---|
| **1** | Basics | Basic syntax, types, variables, control flow |
| **2** | Structs & Functions | Structs, methods, functions, and code organization |
| **3** | Interfaces & Errors | Interfaces, idiomatic error handling in Go |
| **4** | Collections, Pointers & Memory | Slices, maps, pointers, and memory management |
| **5** | Concurrency | Goroutines, channels, and the *Concurrent URL Fetcher* mini-project |
| **6** | Context, Timeouts & Cancellation | Context propagation, timeouts, and operation cancellation |
| **7** | HTTP Services | Building HTTP services and REST APIs in Go |
| **8** | Production Ready | Clean architecture, logging, configuration, graceful shutdown, and testing — culminating in the final mini-project: a **Task Management Service** |

## Final Mini-Project: Task Management Service

The mentorship's closing mini-project consists of building a **production-ready REST API in Go**, applying all the concepts covered across the 8 weeks. Key points addressed include:

- Health check and task CRUD endpoints (`/health`, `/tasks`, `/tasks/{id}`)
- Layered architecture (handler, service, repository)
- Idiomatic error handling (no use of `panic`)
- Structured logging
- Configuration management via environment variables
- Graceful shutdown with OS signal handling
- Unit tests for each application layer
- Optional extras: Docker, CI with GitHub Actions, benchmarking, and profiling

## Technologies

* [Go (Golang)](https://go.dev/)

## How to Run

1. Make sure you have [Go installed](https://go.dev/doc/install) on your machine.
2. Clone this repository:
   ```bash
   git clone https://github.com/Julianapw/Go-Studies.git
   cd Go-Studies
   ```
3. Navigate to the desired chapter or project folder and run:
   ```bash
   go run main.go
   ```

## Notes

This is a study repository: the code reflects the learning process throughout the mentorship and may be refined continuously.

## License

This project is free to use for educational and study purposes.