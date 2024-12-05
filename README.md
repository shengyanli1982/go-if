English | [中文](./README_CN.md)

<div align="center">
	<h1>go-if</h1>
    <p>A tiny, elegant implementation of conditional expressions for Go - your missing ternary operator.</p>
    <img src="assets/logo.png" alt="logo" width="350px">
</div>

[![Go Report Card](https://goreportcard.com/badge/github.com/shengyanli1982/go-if)](https://goreportcard.com/report/github.com/shengyanli1982/go-if)
[![Build Status](https://github.com/shengyanli1982/go-if/actions/workflows/test.yaml/badge.svg)](github.com/shengyanli1982/go-if/actions)
[![Go Reference](https://pkg.go.dev/badge/github.com/shengyanli1982/go-if.svg)](https://pkg.go.dev/github.com/shengyanli1982/go-if)

# Introduction

Missing the ternary operator in Go? Say hello to `go-if`! This tiny library brings elegant conditional expressions to Go, making your code cleaner and more expressive. It's like having the ternary operator, but in a way that feels natural in Go. 🎯

`go-if` offers:

1. Generic implementation that works with any type
2. Clean, simple syntax that's easy to read and understand
3. Zero external dependencies
4. Comprehensive test coverage
5. Type-safe operations
6. Perfect for one-liners without sacrificing readability

# Why Choose go-if?

-   **Simple API**: One function that does one thing perfectly
-   **Type Safety**: Fully generic implementation that works with any type
-   **Zero Dependencies**: No external packages required
-   **Production Ready**: Thoroughly tested with various types and edge cases
-   **Familiar Pattern**: Works like the ternary operator you know and love
-   **Go Idiomatic**: Designed to feel natural in Go code

# Installation

To install `go-if`, use the `go get` command:

```bash
go get github.com/shengyanli1982/go-if
```

# Quick Start

Here's how simple it is to use `go-if`:

```go
package main

import (
    "fmt"
    goif "github.com/shengyanli1982/go-if"
)

func main() {
    // Simple string condition
    result := goif.If(true, "yes", "no")
    fmt.Println(result) // Output: yes

    // With numbers
    count := goif.If(len(items) > 5, 10, 5)
    fmt.Println(count)

    // With complex types
    user := goif.If(isAdmin,
        User{Role: "admin", Permissions: []string{"read", "write"}},
        User{Role: "guest", Permissions: []string{"read"}},
    )
}
```

# Features and Use Cases

## 1. Basic Type Operations

Perfect for simple conditional assignments:

```go
// String operations
greeting := goif.If(isEnglish, "Hello", "Hola")

// Numeric operations
discount := goif.If(isPremiumUser, 0.2, 0.1)

// Boolean flags
isEnabled := goif.If(config.Debug, true, false)
```

## 2. Working with Complex Types

Handles any type elegantly:

```go
// Struct selection
config := goif.If(isProduction,
    Config{
        Host:    "prod.example.com",
        Timeout: 30,
    },
    Config{
        Host:    "dev.example.com",
        Timeout: 5,
    },
)

// Slice operations
permissions := goif.If(isAdmin,
    []string{"read", "write", "delete"},
    []string{"read"},
)

// Map selection
cache := goif.If(isDistributed,
    map[string]int{"capacity": 1000, "ttl": 3600},
    map[string]int{"capacity": 100, "ttl": 300},
)
```

## 3. Pointer and Interface Handling

Safely work with pointers and interfaces:

```go
// Pointer selection
var fallbackCache *Cache = nil
activeCache := goif.If(isCacheEnabled,
    mainCache,
    fallbackCache,
)

// Interface handling
var handler interface{} = goif.If(isAsync,
    AsyncHandler{},
    SyncHandler{},
)
```

## 4. Time and Duration Operations

Perfect for timing-related conditions:

```go
// Timeout duration
timeout := goif.If(isHighPriority,
    5 * time.Second,
    30 * time.Second,
)

// Time selection
deadline := goif.If(isUrgent,
    time.Now().Add(1 * time.Hour),
    time.Now().Add(24 * time.Hour),
)
```

# Performance Benchmarks

We've conducted comprehensive benchmarks comparing `go-if` with traditional if-else statements across different data types:

```bash
$ go test -bench=. -benchmem
goos: windows
goarch: amd64
pkg: github.com/shengyanli1982/go-if
cpu: 12th Gen Intel(R) Core(TM) i5-12400F
BenchmarkString/Generic_If-12                   1000000000               0.1293 ns/op          0 B/op          0 allocs/op
BenchmarkString/Traditional_IfElse-12           1000000000               0.1297 ns/op          0 B/op          0 allocs/op
BenchmarkInteger/Generic_If-12                  1000000000               0.1299 ns/op          0 B/op          0 allocs/op
BenchmarkInteger/Traditional_IfElse-12          1000000000               0.1310 ns/op          0 B/op          0 allocs/op
BenchmarkStruct/Generic_If-12                   1000000000               0.1285 ns/op          0 B/op          0 allocs/op
BenchmarkStruct/Traditional_IfElse-12           1000000000               0.1306 ns/op          0 B/op          0 allocs/op
BenchmarkSlice/Generic_If-12                    1000000000               0.1284 ns/op          0 B/op          0 allocs/op
BenchmarkSlice/Traditional_IfElse-12            1000000000               0.1281 ns/op          0 B/op          0 allocs/op
```

## Running Benchmarks

To run the benchmarks yourself:

```bash
go test -bench=. -benchmem
```

For specific type benchmarks:

```bash
# String operations only
go test -bench=BenchmarkString -benchmem

# Integer operations only
go test -bench=BenchmarkInteger -benchmem

# Struct operations only
go test -bench=BenchmarkStruct -benchmem

# Slice operations only
go test -bench=BenchmarkSlice -benchmem
```

## Benchmark Cases

The benchmarks cover four main categories:

1. **String Operations**: Comparing string conditional assignments
2. **Integer Operations**: Testing numeric value selections
3. **Struct Operations**: Evaluating complex struct type handling
4. **Slice Operations**: Measuring performance with slice assignments

Each benchmark compares:

-   `Generic_If`: Using our generic `If` function
-   `Traditional_IfElse`: Using traditional if-else statements

This helps you make informed decisions about when to use `go-if` in your codebase.

# Best Practices

## 1. Keep It Simple

The beauty of `go-if` lies in its simplicity. Use it for clear, straightforward conditions:

```go
// Good - clear and simple
message := goif.If(count > 0, "Items found", "No items")

// Avoid - overly complex
// Instead, use regular if statements for complex logic
result := goif.If(
    complexCondition() && anotherCondition(),
    someComplexFunction(),
    anotherComplexFunction(),
)
```

## 2. Type Consistency

Always ensure both return values are of the same type:

```go
// Good - both values are strings
path := goif.If(isWindows, "C:\\path", "/path")

// Bad - mixing types will not compile
// value := goif.If(condition, "string", 123)
```

# Limitations

Let's be clear about what `go-if` isn't:

-   Not a replacement for all if statements (use regular if statements for complex logic)
-   Not suitable for side effects (use regular if statements when you need to execute multiple statements)
-   Not designed for nested conditions (consider refactoring or using regular if statements)

# Contributing

Contributions to `go-if` are welcome! Please feel free to submit a Pull Request.

# License

`go-if` is released under the MIT License. See the LICENSE file for details.
