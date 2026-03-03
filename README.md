# Go Reloaded - Text Completion Tool

A Go-based CLI tool that performs automated text editing, number conversion, and punctuation formatting.

## Features
- **Base Conversion**: Converts `(hex)` and `(bin)` to decimal.
- **Case Formatting**: Supports `(up)`, `(low)`, and `(cap)` (including multi-word support).
- **Article Correction**: Automatically switches `a` to `an` before vowels/h.
- **Punctuation**: Fixes spacing for common marks and handles single quotes.

## Usage
1. Place your text in `sample.txt`.
2. Run the program:
   ```bash
   go run . sample.txt result.txt
