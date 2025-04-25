# Cook Concurrent Generation Example

This example demonstrates how to use the Cook Generator for efficient concurrent pattern generation.

## What This Example Shows

1. How to initialize a `CookGenerator` once and share it across multiple goroutines
2. How to concurrently generate multiple different pattern sets
3. How to collect and aggregate results from concurrent operations
4. How to use different method combinations with the generator

## Running the Example

```bash
cd v2/examples/concurrent
go run main.go
```

## Key Concepts

The `CookGenerator` provides an efficient way to reuse shared resources across multiple operations. This example demonstrates:

- Using `cook.NewGenerator()` to create a single generator with pre-initialized resources
- Running concurrent generation tasks with goroutines
- Using Go's concurrency patterns (WaitGroup, channels) for coordinating work
- Generating patterns with various method combinations:
  - Basic patterns
  - Range patterns
  - Method transformations (uppercase, md5, base64, etc.)
  - Column-specific methods

This approach is particularly useful for applications that need to generate large numbers of patterns efficiently, leveraging multiple cores for parallel processing.
