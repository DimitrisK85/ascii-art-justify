# ASCII-Art-Color

A Go program that converts text into ASCII art using banner files with optional color support.

## Usage

```bash
# Basic usage
go run . "your text here"

# Color entire string
go run . --color=<color> "your text here"

# Color specific substring
go run . --color=<color> "substring" "your text here"
```

## Examples

```bash
# Basic ASCII art
go run . "Hello"

# Color entire string in red
go run . --color=red "Hello"

# Color only "ell" substring in blue
go run . --color=blue "ell" "Hello"
```

Output:
```
 _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                
```

## Features

- Converts printable ASCII characters (32-126) to ASCII art
- Supports newlines with `\n`
- Color support for entire strings or specific substrings
- Available colors: red, green, yellow, blue, magenta, cyan, white, black, orange
- Uses standard Go library only

## Testing

```bash
go test ./...
```

## Project Structure

- `main.go` - Entry point
- `internal/banner/` - Banner file loader
- `internal/converter/` - Text to ASCII art converter with color support
- `internal/color/` - Color parsing and ANSI code handling
- `banners/` - ASCII art banner files (standard, shadow, thinkertoy)
