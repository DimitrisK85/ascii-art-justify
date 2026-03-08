# Product Requirements Document (PRD)
## ASCII-Art-Color Generator

### Project Overview
A command-line program written in Go that converts text strings into ASCII art representations using predefined banner templates with optional color support for entire strings or specific substrings.

### Objectives
- Accept a string as a command-line argument
- Output the string in graphical ASCII art format
- Support multiple banner styles (standard, shadow, thinkertoy)
- Handle various input types: letters, numbers, spaces, special characters, and newlines
- Apply colors to entire output or specific substrings

### Technical Requirements

#### 1. Input Specifications
- Accept 1-3 command-line arguments:
  - 1 arg: text to convert
  - 2 args: `--color=<color>` flag and text (colors entire string)
  - 3 args: `--color=<color>` flag, substring to color, and text
- Support printable ASCII characters
- Handle special character `\n` for newlines
- Process empty strings and newline-only inputs
- Supported colors: red, green, yellow, blue, magenta, cyan, white, black, orange

#### 2. Banner File Format
- Each character has exactly 8 lines of height
- Characters are separated by a newline (`\n`)
- Three banner files must be supported:
  - `standard.txt`
  - `shadow.txt`
  - `thinkertoy.txt`
- Banner files should not be modified

#### 3. Output Requirements
- Each line of output corresponds to one line across all characters
- Characters are printed horizontally, line by line
- Empty input (`""`) produces no output
- Single newline (`"\n"`) produces one blank line
- Multiple newlines produce corresponding blank lines
- Trailing newline in output (8 lines per text line)
- ANSI color codes applied during conversion for accurate substring coloring

#### 4. Error Handling
- Return "ERROR" for any file format issues
- Return "ERROR" for invalid banner files
- Handle edge cases gracefully

#### 5. Programming Constraints
- Language: Go
- Packages: Standard library only
- Code must follow Go best practices
- Unit tests required (TDD approach)

### Functional Requirements

#### Core Functionality
1. **Parse Command-Line Arguments**
   - Read 1-3 arguments from command line
   - Parse `--color=<color>` flag if present
   - Identify substring to color (if 3 args provided)
   - Validate input exists

2. **Load Banner File**
   - Read appropriate banner file
   - Parse character definitions
   - Store character mappings

3. **Process Input String**
   - Split input by `\n` for multi-line support
   - Identify character positions to color based on substring matches
   - Convert each character to its ASCII art representation
   - Apply ANSI color codes during conversion
   - Handle spaces and special characters

4. **Generate Output**
   - Combine character representations horizontally
   - Insert color codes at appropriate positions
   - Print 8 lines for each line of text
   - Add blank line for standalone `\n`

### Test Cases (from Usage Examples)

| Input | Expected Behavior |
|-------|------------------|
| `""` | No output |
| `"\n"` | Single blank line |
| `"Hello\n"` | ASCII art of "Hello" (8 lines) |
| `"hello"` | ASCII art of "hello" (8 lines) |
| `"HeLlO"` | Mixed case ASCII art (8 lines) |
| `"Hello There"` | ASCII art with space (8 lines) |
| `"1Hello 2There"` | ASCII art with numbers (8 lines) |
| `"{Hello There}"` | ASCII art with special chars (8 lines) |
| `"Hello\nThere"` | Two separate ASCII art blocks |
| `"Hello\n\nThere"` | Two blocks with blank line between |
| `--color=red "Hello"` | Red colored ASCII art of "Hello" |
| `--color=blue "ell" "Hello"` | ASCII art with only "ell" in blue |

### Project Structure (Suggested)
```
ascii-art-color/
├── main.go              # Entry point with color argument parsing
├── internal/
│   ├── banner/          # Banner loading logic
│   ├── converter/       # Text to ASCII conversion with color support
│   └── color/           # Color parsing and ANSI code handling
├── banners/             # Banner template files
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
├── tests/               # Unit tests
├── ai/                  # AI conversation logs
│   └── ai.txt
├── agents.md            # Agent guidelines
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
- Error handling returns "ERROR" appropriately
- Program handles edge cases correctly

### Non-Functional Requirements
- Performance: Handle strings up to reasonable length efficiently
- Maintainability: Clean, readable code structure
- Testability: Modular design for easy unit testing
- Portability: Works on any system with Go installed
