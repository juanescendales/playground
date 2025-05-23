# Go Simple Cache

## Overview

`go-simple-cache` is a playground project designed to explore and implement a local in-memory cache in Go. The primary goal is to practice Go programming skills, delve into system design concepts, and experiment with layered architectural patterns.

This project serves as a learning sandbox, intentionally employing a more complex scaffolding than a simple cache might require, to better understand and apply principles of layered architecture.

## Core Idea

The main idea is to build a flexible cache system that can support various caching strategies. Currently, the focus has been on laying the groundwork for this system and implementing the first caching strategy.

## Features Implemented

*   **Local In-Memory Cache**: Stores key-value pairs in memory.
*   **LRU (Least Recently Used) Caching Strategy**: The first implemented eviction policy. When the cache reaches its capacity, the least recently used item is removed to make space for new items.
*   **Layered Architecture**: The project structure attempts to follow principles of layered architecture, separating concerns into distinct layers (e.g., application, domain, infrastructure).
*   **HTTP API Endpoints**:
    *   `POST /key`: Add a key-value pair to the cache.
        *   Request Body: `{"key": "your_key", "value": "your_value"}`
        *   Response: `201 Created` with the added key-value pair.
    *   `GET /key?key=your_key`: Retrieve a value from the cache by its key.
        *   Response: `200 OK` with `{"value": "retrieved_value"}`.
    *   `GET /status`: Get the current status of the cache, including ordered keys and current size.
        *   Response: `200 OK` with `{"keys": ["key1", "key2"], "size": 2}`.
*   **Configurable Cache Capacity**: Cache capacity can be set via `internal/app/config/config.yml`.
*   **Graceful HTTP Server Shutdown**: Implemented robust server shutdown handling OS signals.

## Future Work / Planned Features

*   Implementation of other caching strategies:
    *   LFU (Least Frequently Used)
    *   FIFO (First-In, First-Out)
*   Benchmarking different cache strategies.

## Purpose & Learning Goals

This project is primarily a learning exercise focusing on:
*   **Go Programming**: Enhancing proficiency in Go, including concurrency, interfaces, and standard library usage.
*   **System Design**: Applying fundamental system design principles to a practical problem.
*   **Layered Architectures**: Understanding and implementing separation of concerns and dependency management in a layered system.
*   **Caching Mechanisms**: Gaining a deeper understanding of how different caching strategies work and their trade-offs.

## Getting Started

(This section can be expanded with build and run instructions as the project evolves.)

1.  **Configuration**:
    *   The cache capacity can be configured in `internal/app/config/config.yml`. The default is currently set to 3.
2.  **Running the application**:
    *   Ensure you have Go installed.
    *   Navigate to the project root directory.
    *   Run the application using: `go run ./cmd/server/` (assuming your main package is in `cmd/server/main.go`).

### Example API Calls (using curl)

*   **Add a key-value pair:**
    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"key":"hello","value":"world"}' http://localhost:8080/key
    ```
*   **Get a value by key:**
    ```bash
    curl http://localhost:8080/key?key=hello
    ```
*   **Get cache status:**
    ```bash
    curl http://localhost:8080/status
    ```