# Toolkit

A collection of useful Go utility packages to simplify common tasks.

## Installation

```bash
go get github.com/grzegab/toolkit
```

## Packages

### `datex`
Utility functions for date and time manipulation.
- **Filter**: Filter dates relative to a base date (Older/Newer). Supports both `time.Time` and `string`.
- **Parse**: Parse RFC3339 date strings.
- **Manipulation**: Functions like `RemoveHours`, `RoundToHour`, and `RoundToDay`.

### `grpcx`
Safe type conversion utilities, especially useful when dealing with gRPC generated types or `any` types.
- **ConvertInt**: Safely convert various numeric types (int, int32, int64, float32, float64) and their pointers to `int`.
- **ConvertFloat**: Safely convert various numeric types and their pointers to `float64`.

### `pointerx`
Simple pointer utilities.
- **ToPointer**: A generic helper to get a pointer to any value.

### `slicex`
Generic operations and utilities for slices.
- **Filter**: Filter slice elements based on a predicate.
- **Intersect / Difference**: Set operations for slices.
- **Unique**: Remove duplicates while preserving order.
- **RemoveZero**: Remove zero values from a slice.
- **Sort**: Specialized sorting for strings and dates.
- **Convert**: Utilities like `IntsToStrings`.

### `stringx`
String manipulation and generation.
- **RandomString / RandomSmallString**: Generate random alphanumeric or lowercase strings.
- **Reverse**: Reverse a string (rune-aware).
- **Slug**: Create URL-friendly slugs from strings (includes Polish character transliteration).

### `uuidx`
Wrappers for UUID generation (using `github.com/google/uuid`).
- Supports versions: V1, V2, V3, V4, V5, V6, and V7.

## Usage Examples

### String Slugification
```go
import "github.com/grzegab/toolkit/stringx"

slug := stringx.Slug("Za żółtą gęsią jaźń")
// Output: "za-zolta-gesia-jazn"
```

### Slice Filtering
```go
import "github.com/grzegab/toolkit/slicex"

numbers := []int{1, 2, 3, 4, 5}
even := slicex.Filter(numbers, func(n int) bool {
    return n % 2 == 0
})
// Output: [2, 4]
```

### Generic Pointer Helper
```go
import "github.com/grzegab/toolkit/pointerx"

p := pointerx.ToPointer(42) // returns *int
```

### Safe gRPC Conversions
```go
import "github.com/grzegab/toolkit/grpcx"

var val *int32 = nil
result := grpcx.ConvertInt(val) // returns 0 instead of panicking
```

## License
MIT
