# ApplyMethods Example

This example demonstrates how to use the `ApplyMethods` function from the Cook generator to directly apply transformation methods to a list of strings.

## What it does

The example:
1. Creates a new Cook generator
2. Defines a list of input strings
3. Applies methods like "upper", "reverse", and "leet" to transform the strings
4. Displays both the original and transformed strings

## Running the example

```bash
cd v2/examples/apply_methods
go run main.go
```

## Expected output

```
Original strings:
  - test
  - hello
  - world

After applying methods (upper, reverse):
  - TSET
  - OLLEH
  - DLROW

Applying different methods (leet, lower):
  - t3$t
  - h3ll0
  - w0rld
```

Note: The actual output may vary depending on how the methods are configured in your Cook installation.
