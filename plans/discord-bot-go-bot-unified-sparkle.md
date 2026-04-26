# Discord Dice Bot Implementation Plan

## Context
Create a Discord dice bot in Go that responds when mentioned with dice notation (e.g., `4D6`, `2D100`). The bot rolls the dice and returns individual results plus total. Project is currently empty — everything must be created from scratch.

**Response format:** `[3, 5, 1, 6] 合計: 15`
**Trigger:** Bot mention only (no slash commands)
**Reply type:** Normal message reply to original

## File Structure

```
go.mod
go.sum
main.go
dice/
  parser.go      — Regex parsing of "NDN" notation
  validator.go   — Range validation (count 1-100, sides 1-100)
  roller.go      — Core rolling logic via crypto/rand
  result.go      — Result struct, String(), RollNotation() helper
  roll_test.go   — Unit tests
```

## Implementation Steps

### Step 1: Initialize Go module
```bash
go mod init github.com/nekonenene/discord-dice-bot-qwen-claude-code
```

### Step 2: Create `dice/parser.go`
- Regex: `(?i)^(\d+)D(\d+)` to match `NDN` notation (case-insensitive D)
- `Parse(input string) (count, sides int, err error)` — extracts and converts to int
- Returns error if regex doesn't match or conversion fails

### Step 3: Create `dice/validator.go`
- `Validate(count, sides int) error`
- Count must be 1-100, sides must be 1-100
- Descriptive error messages for each violation

### Step 4: Create `dice/roller.go`
- `Roll(count, sides int) ([]int, error)` using `crypto/rand`
- Use `rand.Int(rand.Reader, big.NewInt(int64(sides)))` for bias-free generation
- Results in range [1, sides]

### Step 5: Create `dice/result.go`
- `Result` struct with Count, Sides, Values, Total fields
- `String() string` → `"[v1, v2, ..., vn] 合計: total"`
- `RollNotation(input string) (Result, error)` — orchestrates Parse → Validate → Roll → format

### Step 6: Create `dice/roll_test.go`
- Table-driven tests for Parse, Validate, Roll bounds, RollNotation, Result.String()

### Step 7: Create `main.go`
- Read `DISCORD_BOT_TOKEN` from env var
- Create `discordgo.Session`, register `OnMessageCreate` handler
- Handler logic:
  1. Ignore non-mention messages
  2. Strip mention prefix, extract dice notation
  3. Call `dice.RollNotation()`
  4. Reply with result string or error message
- Graceful shutdown on SIGINT/SIGTERM

### Step 8: Build and verify
```bash
go mod tidy
go fmt ./...
go vet ./...
go test ./dice/
go build -o dice-bot .
```

## Key files to modify
- `main.go` (new)
- `dice/parser.go` (new)
- `dice/validator.go` (new)
- `dice/roller.go` (new)
- `dice/result.go` (new)
- `dice/roll_test.go` (new)
- `go.mod` (new)

## Verification
1. `go test ./dice/` — all unit tests pass
2. `go build` — compiles without errors
3. Manual test: run bot with valid token, mention with `4D6`, verify response `[x, x, x, x] 合計: N`
4. Edge cases: `0D6` → error, `101D6` → error, `4D0` → error, `4D101` → error, `abc` → error
