## ez_utils

A high-performance utility library designed to eliminate boilerplate code in Go applications. Built for production-scale services and CLI tools.

### Included Packages:

#### **`colors`** — Terminal Text Coloring
Simplified terminal text coloring using standard Hex codes (`#000000`) with automatic terminal capability detection.
- Supports 24-bit truecolor and 8-color fallback modes
- 80+ predefined color constants (Red, Blue, Crimson, ForestGreen, etc.)
- ANSI escape sequence generation
- Color conversion utilities (`ToRGB`, `ToRGBA`, `HexToAnsi`)
- Functions: `ColorPrint`, `ColorPrintln`, `SupportsTrueColor()`

#### **`scanner`** — Type-Safe CLI Input
A robust CLI input scanner featuring native EOF handling, string sanitization, and type-safe parsing with validation loops.
- `Scanln()` — Read any string (allows empty input)
- `Nextstring()` — Read non-empty strings with validation
- `Nextint()` — Read and parse integers
- `Nextfloat32()` — Read and parse float32 numbers
- `Nextbool()` — Read booleans (accepts: true/false, 1/0, y/n, yes/no)

#### **`input`** — Keyboard Event Handling
Real-time keyboard input detection with multi-threaded event state tracking for interactive CLI applications.
- `IsPressing(key)` — Check if a key is currently being pressed
- `Pressed(key)` — Check if a key was pressed since last check
- `ResetEvents()` — Clear all key press events
- Supports: arrow keys, space, enter, escape, backspace, tab, and alphanumeric keys

### Installation
```bash
go get github.com/cribest7890/ez_utils/@latest
```

### Quick Start

**Colors Example:**
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

**Scanner Example:**
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

**Input Example:**
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

### License
This project is licensed under the GNU Affero General Public License v3.0 — see LICENSE file for details.
