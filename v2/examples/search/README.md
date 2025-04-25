# Cook Search Example

This example demonstrates how to use the Search functionality with the CookGenerator pattern.

## What This Example Shows

1. How to initialize a `CookGenerator` once and reuse it for efficient operations
2. How to use the `Search` method to find ingredients in the Cook database
3. How to process and display search results
4. How to use the same generator instance for both searching and pattern generation

## Running the Example

```bash
cd v2/examples/search
go run main.go
```

## Key Concepts

The `CookGenerator` provides an efficient way to reuse shared resources across multiple operations. This example demonstrates:

- Using `cook.NewGenerator()` to create a generator with pre-initialized resources
- Using `generator.Search(query)` to search for ingredients
- Processing the `CookIngredient` results returned from the search
- Using the same generator for pattern generation with `generator.Generate(pattern)`

This approach is particularly useful for applications that need to perform multiple Cook operations without the overhead of re-initializing resources for each operation.
