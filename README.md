## ez_utils
<p align="center">
  <a href="https://golang.dev">
    <img src="https://img.shields.io/badge/go-1.21-blue?logo=go&logoColor=white&style=for-the-badge" alt="Go">
  </a>
  <a href="https://github.com/cribest7890/ez_utils/releases">
    <img src="https://img.shields.io/github/v/release/cribest7890/ez_utils?style=for-the-badge&label=Latest%20Release&color=2ea44f" alt="Latest Release">
  </a>
  <a href="https://github.com/cribest7890/ez_utils/blob/main/LICENSE">
    <img src="https://img.shields.io/github/license/cribest7890/ez_utils?style=for-the-badge&color=1f6feb" alt="License">
  </a>
</p>

![reference(https://pkg.go.dev/badge/github.com/cribest7890/ez_utils.svg)](https://pkg.go.dev/github.com/cribest7890/ez_utils)  

A high-performance utility library to reduce boilerplate in Go applications. Built for production services and CLI tools.

## Included packages

### `colors` — Terminal Text Coloring
Simplified terminal text coloring using standard hex codes (`#000000`) with automatic terminal capability detection.

Features:
- Supports 24-bit truecolor and 8-color fallback modes
- 80+ predefined color constants (Red, Blue, Crimson, ForestGreen, etc.)
- ANSI escape sequence generation
- Color conversion utilities (`ToRGB`, `ToRGBA`, `HexToAnsi`)
- Convenience functions: `ColorPrint`, `ColorPrintln`, `SupportsTrueColor()`

### `scanner` — Type-Safe CLI Input
A robust CLI input scanner featuring native EOF handling, string sanitization, and type-safe parsing with validation loops.

Functions:
- `Scanln()` — Read any string (allows empty input)
- `Nextstring()` — Read non-empty strings with validation
- `Nextint()` — Read and parse integers
- `Nextfloat32()` — Read and parse float32 numbers
- `Nextbool()` — Read booleans (accepts: true/false, 1/0, y/n, yes/no)

### `input` — Keyboard Event Handling
Real-time keyboard input detection with multi-threaded event state tracking for interactive CLI applications.

Functions:
- `IsPressing(key)` — Check if a key is currently being pressed
- `Pressed(key)` — Check if a key was pressed since last check
- `ResetEvents()` — Clear all key press events

Supported keys: arrow keys, space, enter, escape, backspace, tab, and alphanumeric keys.

### `convert` — Flexible Type Conversion Utilities
Helpers to coerce values into common Go types with sensible defaults.

Functions:
- `ToInt(value any, defaultValue int) int` — Convert many types to int
- `ToFloat64(value any, defaultValue float64) float64` — Convert to float64
- `ToString(value any) string` — Convert to string
- `ToBool(value any, defaultValue bool) bool` — Convert to bool

### `files` — File helper utilities
Convenience wrappers for common file operations.

Functions:
- `Exists(path string) bool` — Check if a file or directory exists
- `Write(path, content string) error` — Write (creates dirs as needed)
- `Append(path, content string) error` — Append to a file (creates if missing)
- `Read(path string) (string, error)` — Read file contents
- `Clear(path string) error` — Truncate file
- `Delete(path string) error` — Remove file if exists

### `random` — Secure random helpers
Random generation utilities built on crypto/rand.

Functions:
- `Int(min, max int) int` — Random int in range
- `Float64() float64` — Random float in [0,1)
- `Float64Range(min, max float64) float64` — Random float in range
- `Bool() bool` — Random boolean
- `String(length int) string` — Random alphanumeric string
- `Choice[T any](slice []T) T` — Pick random element
- `Shuffle[T any](slice []T) []T` — Shuffle slice

### `sys` — OS helpers
Small utilities for OS-specific behaviors and helpers (platform-specific files included).


## Installation

Requires Go 1.21+

```bash
go get github.com/cribest7890/ez_utils@latest
```

(Alternatively use Go modules: `go install github.com/cribest7890/ez_utils@latest`)

## Quick start

### Colors example

```go
package main

import (
	"github.com/cribest7890/ez_utils/colors"
)

func main() {
	colors.ColorPrintln("Success!", colors.Green)
	colors.ColorPrintln("Error message", colors.Red)
	colors.ColorPrintln("Warning", colors.Orange)
}
```

### Scanner example

```go
package main

import (
	"fmt"
	"github.com/cribest7890/ez_utils/scanner"
)

func main() {
	fmt.Print("Enter your name: ")
	name := scanner.Nextstring()

	fmt.Print("Enter your age: ")
	age := scanner.Nextint()

	fmt.Printf("Hello %s, you are %d years old!\n", name, age)
}
```

### Convert example

```go
package main

import (
	"fmt"
	"github.com/cribest7890/ez_utils/convert"
)

func main() {
	fmt.Println(convert.ToInt("42", 0))         // 42
	fmt.Println(convert.ToFloat64("3.14", 0.0)) // 3.14
	fmt.Println(convert.ToBool("true", false))  // true
}
```

### Files example

```go
package main

import (
	"fmt"
	"github.com/cribest7890/ez_utils/files"
)

func main() {
	_ = files.Write("tmp/example.txt", "hello world")
	content, _ := files.Read("tmp/example.txt")
	fmt.Println(content)
}
```

### Random example

```go
package main

import (
	"fmt"
	"github.com/cribest7890/ez_utils/random"
)

func main() {
	fmt.Println(random.Int(1, 10))
	fmt.Println(random.String(8))
}
```

### Input example

```go
package main

import (
	"github.com/cribest7890/ez_utils/input"
)

func main() {
	if input.Pressed(input.KeyEnter) {
		// Handle enter key press
	}
	if input.IsPressing(input.KeyArrowUp) {
		// Handle up arrow held down
	}
}
```

## Contributing

Contributions welcome — please open issues or pull requests. Follow standard Go formatting (gofmt) and include tests for new features when possible.

## License

This project is licensed under the GNU Affero General Public License v3.0 — see the LICENSE file for details.
