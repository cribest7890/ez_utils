# ez_utils

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

## Installation

Requires Go 1.18+

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
