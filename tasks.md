# ascii-art-justify Task Cards

## TASK-01 Baseline Freeze
- Purpose: Confirm current project behavior before color changes.
- Scope: Run all existing tests.
- Method: Run tests first, record failures if any, and avoid code edits in this task.
- Done when: `go test ./...` passes.

## TASK-02 CLI Contract Definition
- Purpose: Lock accepted command forms and usage error behavior.
- Scope:
  - `go run . "text"`
  - `go run . --color=<color> "text"`
  - `go run . --color=<color> <substring> "text"`
  - Also supports optional trailing banner in each supported mode.
  - Invalid format prints an error plus usage guidance.
  - Method: Translate each rule into test cases before coding parser logic.
  - Done when: Rules are represented in tests.

## TASK-03 Parser Tests (Red)
- Purpose: Create failing tests for argument parsing.
- Scope:
  - Valid: `len==2`, `len==3`, `len==4` (after optional banner detection)
  - Invalid: bad flag format
  - Invalid: argument count `<2` or extra/unrecognized combinations
- Method: Add table-driven tests for args -> expected parse result or usage error.
- Done when: Tests fail for missing parser behavior.

## TASK-04 Parser Implementation (Green)
- Purpose: Implement minimum parser logic.
- Scope:
  - Parse single-string mode
  - Parse color whole-string mode
  - Parse color substring mode
  - Validate strict `--color=<color>` format
- Method: Implement a small parse function with explicit branches by arg count and flag prefix.
- Done when: TASK-03 tests pass.

## TASK-05 Color Behavior Tests (Red)
- Purpose: Define expected color behavior.
- Scope:
  - Whole-string coloring when substring omitted
  - All substring matches colored
  - Unmatched substring leaves output unchanged
- Method: Write unit tests around conversion output and assert ANSI start/reset placement.
- Done when: Tests fail before color implementation.

## TASK-06 Color Rendering (Green)
- Purpose: Implement ANSI color output.
- Scope:
  - Map supported color names to ANSI codes
  - Wrap target output with color and reset sequences
  - Keep non-target output unchanged
- Method: Build a deterministic color map and color only matched segments while preserving art layout.
- Done when: TASK-05 tests pass.

## TASK-07 CLI Integration Tests (Red)
- Purpose: Verify end-to-end argument flow.
- Scope:
  - Invalid format -> exact usage text
  - Single-string mode still works
  - Color modes route correctly
- Method: Add integration-style tests for main flow with controlled args and captured output.
- Done when: Integration tests fail before final wiring.

## TASK-08 Integration Implementation (Green)
- Purpose: Wire parser and rendering into main flow.
- Scope:
  - Preserve existing ascii-art behavior
  - Add color behavior without breaking base mode
- Method: Connect parse result to existing load/convert pipeline without changing banner loading behavior.
- Done when: Unit + integration tests pass.

## TASK-09 Cleanup (Clean)
- Purpose: Refactor for readability only.
- Scope: Small clarity refactors, no behavior change.
- Method: Extract tiny helpers, simplify names/branches, and rerun full tests after each small change.
- Done when: All tests still pass.

## TASK-10 Audit Verification
- Purpose: Validate against subject example behavior.
- Scope:
  - Run full tests
  - Manual check:
    - `go run . --color=red kit "a king kitten have kit"`
- Method: Verify exact behavior against subject wording and confirm no extra behavior was introduced.
- Done when: Output behavior matches project rules.

## TASK-11 Banner Contract Definition
- Purpose: Define how banner selection fits with existing color support.
- Scope:
  - `go run . "text" [BANNER]`
  - `go run . --color=<color> "text" [BANNER]`
  - `go run . --color=<color> <substring> "text" [BANNER]`
  - Invalid banner or format prints the exact usage message for the banner subject.
- Method: Lock accepted positional banner forms first, then keep parser behavior consistent with future extension work.
- Done when: Accepted forms and invalid cases are represented in tests.

## TASK-12 Banner Parsing Tests (Red)
- Purpose: Add failing tests for valid and invalid banner-aware CLI forms.
- Scope:
  - Valid: `standard`, `shadow`, `thinkertoy`
  - Invalid: unknown banner
  - Valid combinations with existing color modes
- Method: Extend table-driven parser tests with banner cases before parser changes.
- Done when: New tests fail against old parser behavior.

## TASK-13 Banner Parsing Implementation (Green)
- Purpose: Update parser to detect and return a trailing banner selection.
- Scope:
  - Default banner remains `standard`
  - Last argument is treated as banner only if valid
  - Existing color parsing remains compatible
- Method: Parse the trailing banner first, reduce effective arg count, then apply pattern-based parsing.
- Done when: TASK-12 tests pass.

## TASK-14 Banner Loading Integration
- Purpose: Load the selected banner template instead of always using `standard`.
- Scope:
  - Build the correct `banners/<name>.txt` path
  - Preserve plain and color rendering behavior
- Method: Use parsed banner name in main flow and keep banner loading centralized in the existing loader.
- Done when: Manual runs with different banners produce different correct outputs.

## TASK-15 Banner Verification
- Purpose: Validate banner support against subject examples and current optional features.
- Scope:
  - Manual checks for `standard`, `shadow`, and `thinkertoy`
  - Full test suite
  - One color + banner combination
- Method: Verify that banner selection and color mode both behave without regressing baseline flows.
- Done when: Outputs match expectations and tests remain green.

## TASK-16 Parser Refactor Decision Log
- Purpose: Record the decision to refactor CLI parsing before adding more flags.
- Scope:
  - Document that multiple future flags (`--output`, later `--align`) justify a structured parser result
  - Keep the change behavior-preserving
- Method: Log the decision first, then treat the refactor as a cleanup step before new feature work.
- Done when: The design decision is recorded in the AI usage index.

## TASK-17 Parser Struct Refactor (Clean)
- Purpose: Refactor `parseArgs` to return a small `cliOptions` struct instead of many values.
- Scope:
  - Keep current color and banner behavior unchanged
  - Move parser output into explicit named fields
  - Preserve all existing tests
- Method: Introduce a simple struct in `package main`, update parser and callers, and keep tests green throughout.
- Done when: `parseArgs` returns a struct and the current suite still passes.

## TASK-18 Output Contract Definition
- Purpose: Define how `--output=<fileName.txt>` fits with existing banner and color features.
- Scope:
  - `go run . --output=<fileName.txt> [STRING] [BANNER]`
  - No color alignment chaining is supported in the current mode.
  - Empty or malformed output flag is invalid
  - Final exact usage guidance is provided from the main error path.
  - Method: Lock a strict and extensible CLI shape before adding parser tests.
  - Done when: Accepted output forms are represented in tests.

## TASK-19 Output Parsing Tests (Red)
- Purpose: Add failing tests for valid and invalid `--output` flag parsing.
- Scope:
  - Valid output-only forms
  - Valid output + color forms
  - Invalid malformed `--output` forms
  - Empty output filename
- Method: Extend parser tests first so output support is driven by failing expectations.
- Done when: New output-related parser tests fail before implementation.

## TASK-20 Output Parsing Implementation (Green)
- Purpose: Extend the refactored parser to capture an output filename.
- Scope:
  - Parse `--output=<fileName.txt>`
  - Keep existing banner and color parsing compatible
  - No output file means normal stdout behavior
- Method: Parse recognized flags first, then parse remaining positional arguments.
- Done when: TASK-19 tests pass.

## TASK-21 Output Writing Tests (Red)
- Purpose: Add tests that verify rendered output can be written to a file.
- Scope:
  - File is created or overwritten
  - File contents match rendered ASCII art
  - Existing stdout mode remains unchanged when `--output` is absent
- Method: Use standard library file handling in tests and compare file contents against expected output.
- Done when: Output-file tests fail before implementation.

## TASK-22 Output Writing Implementation (Green)
- Purpose: Write rendered output to the requested file when `--output` is provided.
- Scope:
  - Preserve current rendering pipeline
  - Write the exact same output that would otherwise be printed
  - Keep non-output mode unchanged
- Method: Render once, join lines deterministically, then either write to file or print to stdout.
- Done when: Output-file tests pass and existing tests still pass.

## TASK-23 Align Contract Definition
- Purpose: Define how `--align=<type>` fits the existing one-mode CLI design.
- Scope:
  - Supported values: `left`, `right`, `center`, `justify`
  - `--align` is a standalone leading mode flag
  - It is not combined with `--color` or `--output`
  - Invalid align flag or value prints the final unified usage message
- Method: Lock the standalone-mode rule first so parser and runtime behavior stay deterministic.
- Done when: Accepted and rejected align forms are represented in tests.

## TASK-24 Align Parsing Tests (Red)
- Purpose: Add failing parser tests for valid and invalid `--align` forms.
- Scope:
  - Valid: `--align=left`, `--align=right`, `--align=center`, `--align=justify`
  - Invalid: malformed `--align` flag
  - Invalid: unsupported align value
- Method: Extend parser tests before changing `args.go`, following the same Red-first approach.
- Done when: New align parser tests fail before implementation.

## TASK-25 Align Parsing Implementation (Green)
- Purpose: Extend the parser to capture the selected alignment mode.
- Scope:
  - Populate `opts.Align`
  - Keep align as a standalone mode in `args[1]`
  - Preserve existing plain, color, output, and banner behavior
- Method: Add a dedicated align parsing branch in `args.go` without introducing multi-flag combinations.
- Done when: TASK-24 tests pass.

## TASK-26 Align Helper Tests (Red)
- Purpose: Define alignment behavior on rendered ASCII-art lines.
- Scope:
  - Left alignment pads lines to the target width (baseline behavior)
  - Right alignment pads to terminal width
  - Center alignment pads evenly
  - Justify distributes spacing across word gaps
- Method: Test alignment as a post-processing step on rendered `[]string`, independent from banner loading.
- Done when: Alignment helper tests fail before implementation.
Status: Completed

## TASK-27 Align Implementation (Green)
- Purpose: Implement line alignment for rendered ASCII-art output.
- Scope:
  - Apply alignment after rendering
  - Keep left as the baseline
  - Support `right`, `center`, and `justify` according to terminal width
- Method: Add a focused alignment helper (likely in `align.go`) that transforms `[]string` based on width and mode.
- Done when: TASK-26 tests pass.

## TASK-28 Terminal Width Handling
- Purpose: Determine terminal width using standard-library-compatible behavior.
- Scope:
  - Read terminal width for alignment decisions
  - Only align text that fits within available width
  - Keep behavior deterministic enough for tests
- Method: Isolate terminal-width lookup so alignment logic can be tested with explicit widths while runtime uses the current terminal size.
- Done when: Runtime alignment uses terminal width and tests can exercise alignment with controlled values.

## TASK-29 Align Verification
- Purpose: Validate alignment behavior against subject examples and the standalone-mode rule.
- Scope:
  - Manual checks for `left`, `right`, `center`, and `justify`
  - Verify banner selection still works with align mode
  - Full test suite
- Method: Confirm aligned ASCII-art output matches the intended placement and that no other modes regress.
- Done when: Alignment outputs look correct and tests remain green.
Status: Completed

## Completion Log
- [x] TASK-01 Baseline Freeze
- [x] TASK-02 CLI Contract Definition
- [x] TASK-03 Parser Tests (Red)
- [x] TASK-04 Parser Implementation (Green)
- [x] TASK-05 Color Behavior Tests (Red)
- [x] TASK-06 Color Rendering (Green)
- [x] TASK-07 CLI Integration Tests (Red)
- [x] TASK-08 Integration Implementation (Green)
- [x] TASK-09 Cleanup (Clean)
- [x] TASK-10 Audit Verification
- [x] TASK-11 Banner Contract Definition
- [x] TASK-12 Banner Parsing Tests (Red)
- [x] TASK-13 Banner Parsing Implementation (Green)
- [x] TASK-14 Banner Loading Integration
- [x] TASK-15 Banner Verification
- [x] TASK-16 Parser Refactor Decision Log
- [x] TASK-17 Parser Struct Refactor (Clean)
- [x] TASK-18 Output Contract Definition
- [x] TASK-19 Output Parsing Tests (Red)
- [x] TASK-20 Output Parsing Implementation (Green)
- [x] TASK-21 Output Writing Tests (Red)
- [x] TASK-22 Output Writing Implementation (Green)
- [x] TASK-23 Align Contract Definition
- [x] TASK-24 Align Parsing Tests (Red)
- [x] TASK-25 Align Parsing Implementation (Green)
- [x] TASK-26 Align Helper Tests (Red)
- [x] TASK-27 Align Implementation (Green)
- [x] TASK-28 Terminal Width Handling
- [x] TASK-29 Align Verification
