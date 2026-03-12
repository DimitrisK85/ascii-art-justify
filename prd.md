# Product Requirements Document (PRD)
## ASCII-Art Tool

### Project Overview
A command-line program written in Go that converts text strings into ASCII art using predefined banner templates, with optional color rendering, output redirection, and line alignment.

### Objectives
- Convert an input string into ASCII-art blocks.
- Support selectable banners: `standard`, `shadow`, and `thinkertoy`.
- Support optional color rendering for the whole text or a substring.
- Support writing output to a `.txt` file.
- Support alignment modes: `left`, `right`, `center`, and `justify`.

### Technical Requirements

#### 1. Input Specifications
- Supported CLI forms:
  - `go run . "text"`
  - `go run . "text" [BANNER]`
  - `go run . --color=<color> "text"`
  - `go run . --color=<color> "text" [BANNER]`
  - `go run . --color=<color> "<substring>" "text"`
  - `go run . --color=<color> "<substring>" "text" [BANNER]`
  - `go run . --output=<filename>.txt "text"`
  - `go run . --output=<filename>.txt "text" [BANNER]`
  - `go run . --align=<left|right|center|justify> "text"`
  - `go run . --align=<left|right|center|justify> "text" [BANNER]`
- Reject invalid argument count and malformed flags.
- Validate `--color` values from the supported color list.
- Validate `--output` has a non-empty `.txt` filename.
- Validate `--align` values only from `left`, `right`, `center`, `justify`.
- Validate input runes:
  - printable ASCII `32..126` is allowed,
  - newline `10` is allowed for multi-line input.
- `\n` in input should be interpreted as a line break (including escaped `\\n` text passed by shell).

#### 2. Banner File Format
- Each character must have 8 rendered lines.
- Character definitions are separated by a blank line in banner files.
- Required banners:
  - `banners/standard.txt`
  - `banners/shadow.txt`
  - `banners/thinkertoy.txt`
- Banner files are input assets and should not be modified by the program.

#### 3. Output Requirements
- Conversion is row-based: each ASCII-art line is rendered across all characters of one input line.
- Each logical text line yields 8 ASCII-art lines.
- Empty input (`""`) produces no output.
- Single newline (`"\n"`) produces one blank line block.
- Consecutive newlines produce corresponding blank line blocks.
- `--color` inserts ANSI color start/end sequences around matched glyph regions.
- `--output=<filename>.txt` writes the same rendered output to file and prints a confirmation message.
- `--align` applies to final rendered lines using terminal width.
  - if rendered width is >= terminal width, lines are returned unchanged.
  - `justify` keeps left alignment for single-word input.

#### 4. Error Handling
- Invalid arguments must return a clear `Error: ...` message and reference usage instructions.
- Invalid banner argument is treated as invalid format unless recognized.
- Empty/invalid output file names and wrong extensions are rejected.
- Unsupported color must be rejected with supported-color guidance.
- Unsupported alignment value must be rejected with supported alignment guidance.
- Banner loader failures are reported as errors.
- Output write failures are reported as errors.

#### 5. Programming Constraints
- Language: Go
- Packages: Go standard library only
- Code follows readable, explicit Go style
- Unit tests and integration-like tests are required (TDD-aligned workflow)

### Functional Requirements

#### Core Functionality
1. **Parse Arguments**
   - Detect optional trailing banner first (`standard`, `shadow`, `thinkertoy`).
   - Recognize exactly one leading mode:
     - none
     - `--color=<color>`
     - `--output=<filename>.txt`
     - `--align=<left|right|center|justify>`
   - For `--color` with 3-arg form, parse substring + text.

2. **Load Banner**
   - Load selected banner file path `banners/<name>.txt`.
   - Build a rune-to-lines mapping for each printable ASCII glyph used.

3. **Render Text**
   - Split input by newline into logical lines.
   - Convert each line to ASCII-art by concatenating glyph rows.
   - Apply color masking and ANSI wrapping when `--color` is active.

4. **Post-process and Emit**
   - Apply requested alignment when `--align` is set.
   - Print final lines to stdout unless `--output` is set.
   - If `--output` is set, write the same rendered lines to the requested file.

### Test Cases (from Usage Examples)

| Input | Expected Behavior |
|-------|------------------|
| Input | Expected Behavior |
|-------|------------------|
| `""` | No output |
| `"\n"` | One blank line output |
| `"Hello\n"` | ASCII art of `Hello` (8 lines, with line break handling) |
| `"Hello\nThere"` | Two separate ASCII-art blocks |
| `"Hello\n\nThere"` | Two blocks with one blank block between |
| `--color=red "Hello"` | Entire output colored in red |
| `--color=blue "ell" "Hello"` | Only `ell` colored in blue |
| `--output=art.txt "Hello"` | Output written to `art.txt` |
| `--align=right "Hello"` | Right alignment to terminal width |
| `--align=justify "Hello there"` | Justified spacing behavior across words |
| `--align=left shadow "Hello"` | Left alignment with selected banner |

### Project Structure (Suggested)
```
ascii-art-justify/
├── main.go              # Entry point with color argument parsing
├── internal/
│   ├── banner/          # Banner loading logic
│   └── converter/       # Text to ASCII conversion with color support
├── banners/             # Banner template files
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
├── *_test.go            # Go test files (main, converter, align, banner loader)
├── ai/                  # AI conversation logs
│   └── task_decomposition_index.txt
├── agents.md            # Agent guidelines
├── tasks.md             # Task cards and completion status
└── prd.md               # This document
```

### Development Approach
- Follow TDD methodology (Red-Green-Refactor)
- Write tests before implementation
- Keep functions small and focused
- Use descriptive variable and function names
- Document complex logic with comments

### Success Criteria
- All usage examples produce correct output
- Color applied accurately to entire strings or specific substrings
- Substring coloring works by tracking character positions during conversion
- Unit tests cover core functionality including color features
- Code follows Go best practices
- No external dependencies used
- Clear `Error: ...` style messages are consistent
- Program handles edge cases correctly

### Non-Functional Requirements
- Performance: Handle strings up to reasonable length efficiently
- Maintainability: Clean, readable code structure
- Testability: Modular design for easy unit testing
- Portability: Works on any system with Go installed
