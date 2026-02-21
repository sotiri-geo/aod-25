# Advent of Code 2025 (Go)

Personal repository for solving Advent of Code 2025 puzzles in Go.

## Project goals

- Practice problem solving with clear, testable implementations.
- Keep each day isolated in its own folder (`day01`, `day02`, etc.).
- Cover core logic with unit tests and runnable entrypoints.

## Tech stack

- Language: Go
- Module: `github.com/sotiri-geo/aod-2025`
- Go version: `1.25.3` (from `go.mod`)

## Repository structure

```text
aod-2025/
├── common/
│   └── asserts.go
├── day01/
│   ├── input.txt
│   ├── main.go
│   └── main_test.go
├── day02/
│   ├── input.txt
│   ├── input-test.txt
│   ├── main.go
│   └── main_test.go
├── day03/
│   ├── input.txt
│   ├── main.go
│   └── main_test.go
├── day04/
│   ├── input.txt
│   ├── main.go
│   └── main_test.go
├── go.mod
└── README.md
```

## Implemented days

### Day 01

- Parses dial rotations like `L30` / `R12`.
- `Part1`: counts wraps to zero with a strict zero-hit incrementor.
- `Part2`: counts passes through zero with a pass-through incrementor.
- Key APIs:
  - `ParseRotation(rotation string) (int, error)`
  - `Part1(input []string) (int, error)`
  - `Part2(input []string) (int, error)`

Run:

```bash
go run ./day01
```

### Day 02

- Parses numeric ranges like `10-20`.
- `Part1`: sums values whose decimal string is exactly two repeated halves (example: `1212`).
- `Part2`: sums values whose decimal string is made of any repeating substring pattern.
- Key APIs:
  - `GetRange(input string) ([]string, error)`
  - `IsRepeatedTwice(input string) bool`
  - `IsRepeated(input string) bool`
  - `Part1(input []string) (int, error)`
  - `Part2(input []string) (int, error)`

Run:

```bash
go run ./day02
```

### Day 03

- Treats each input line as a bank of digits.
- `Part1`: finds best 2-digit value per bank and sums results.
- `Part2`: builds max 12-digit value per bank (monotonic-stack style selection) and sums results.
- Key APIs:
  - `MaxTwoDigit(input []int) int`
  - `MaxTwelveDigit(input []int) int`
  - `Part1(input [][]int) int`
  - `Part2(input [][]int) int`

Run:

```bash
go run ./day03
```

### Day 04

- Models the warehouse as a `Grid` of bytes (`@` for roll of paper).
- For each roll, checks 8-neighborhood adjacency.
- `Part1`: counts currently accessible rolls (`< 4` neighboring rolls).
- `Part2`: repeatedly removes accessible rolls from the grid until no more can be removed, then returns total removed.
- Key APIs:
  - `NewGrid(input [][]byte) *Grid`
  - `CanAccessRollOfPaper(row, col int) bool`
  - `ProcessRollOfPaper() int`
  - `ProcessRollOfPaperWithUpdate() int`

Run:

```bash
go run ./day04
```

## How to run locally

From repo root:

```bash
# Run a specific day
go run ./day01
go run ./day02
go run ./day03
go run ./day04

# Run tests for one day
go test ./day01
go test ./day02
go test ./day03
go test ./day04

# Run all tests/packages
go test ./...
```

## Current test/build status (as of 2026-02-19)

Latest full run in this repository:

- `day01`: passing
- `day02`: passing
- `day03`: passing
- `day04`: passing
- `common`: no test files

## Conventions in this repo

- One folder per puzzle day.
- Each day has:
  - `main.go` with parsing + part implementations + runnable `main()`
  - `input.txt` with puzzle input
  - `main_test.go` for unit tests
- Shared helpers can live in `common/`.

## Notes

- Inputs are read from relative paths like `./dayXX/input.txt`, so run commands from repository root.
- This repository currently includes days 1 to 4.
