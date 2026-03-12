# ASCII-Art-Color

A Go program that converts text into ASCII art using banner files with optional color support.

## How it works

- Reads CLI arguments.
- Validates printable ASCII input.
- Loads the selected banner (`standard`, `shadow`, `thinkertoy`).
- Converts the text to 8-line ASCII art glyphs.
- Supports output alignment.
- Can save output to a `.txt` file.

Notes:
- `\n` in CLI input is interpreted as a newline.
- Supported flags are handled as a single leading mode (`--color`, `--output`, or `--align`).

## Usage (quick start)

Start with the simple forms and add banners/options as needed.

### 1) No flag

```bash
go run . "text"
```

### 1.1) No flag with banner

```bash
go run . "text"                 # uses default banner: standard
go run . "text" shadow
go run . "text" thinkertoy
```

### 2) `--color=<color>`

```bash
go run . --color=<color> "text"                  # color the whole string
go run . --color=<color> "substring" "text"     # color a specific substring
```

### 2.1) `--color` with banner

```bash
go run . --color=red "Hello"
go run . --color=blue "ell" "Hello" shadow
go run . --color=green "e" "Hey there" thinkertoy
```

Supported colors: `black`, `red`, `green`, `yellow`, `blue`, `magenta`, `cyan`, `white`

### 3) `--output=<filename>.txt`

```bash
go run . --output=<filename>.txt "text"
```

### 3.1) `--output` with banner

```bash
go run . --output=art.txt "Hello"
go run . --output=art.txt "Hello" shadow
```

### 4) `--align=<mode>`

```bash
go run . --align=left "text"
go run . --align=right "text"
go run . --align=center "text"
go run . --align=justify "text"
```

### 4.1) `--align` with banner

```bash
go run . --align=center "Hello" 
go run . --align=justify "Hello world" shadow
```

> The alignment mode is standalone: it is used as the main mode flag for the command.

## Common all-in-one format

```bash
go run . [mode] [mode-args] "text" [banner]
```

Where:
- `mode` is one of: `--color=<color>`, `--output=<filename>.txt`, `--align=<left|right|center|justify>`
- `mode-args` depends on the mode:
  - none for `--output` and `--align`
  - substring for `--color` when doing substring coloring
- `banner` is optional and can be `standard`, `shadow`, or `thinkertoy` (default: `standard`)

## Examples

```bash
# Newline in input

go run . "Hello\nThere"

# Different banner

go run . "Hello" shadow

go run . --output=hello-art.txt "Hello there" thinkertoy

# Alignment with color

go run . --color=magenta "ell" "hello" shadow
```

## Example output

```text
 _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                
```

## Features

- Converts printable ASCII characters (32-126) to ASCII art.
- Supports `\n` for line breaks.
- Color mode for full string or substring matches.
- Alignment support: `left`, `right`, `center`, `justify`.
- Banner selection: `standard` (default), `shadow`, `thinkertoy`.
- Output to file with `--output=<filename>.txt`.
- Uses only the Go standard library.

## Testing

```bash
go test ./...
```

## Project Structure

- `main.go` - Entry point
- `internal/banner/` - Banner file loader
- `internal/converter/` - Text-to-ASCII conversion and color mapping
- `banners/` - ASCII art banner files (`standard`, `shadow`, `thinkertoy`)
- `output.go` - File output helper
- `align.go` - Alignment helpers
- `args.go` - CLI argument parser
