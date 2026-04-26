# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Discord dice bot in Go. Responds to bot mentions containing dice notation (e.g., `4D6`, `2D100`) by rolling the dice and replying with individual results and total (e.g., `[3, 5, 1, 6] 合計: 15`).

## Architecture

```
main.go              — Entry point: discordgo session, message handler, reply logic
dice/
  parser.go          — Regex parsing of "NDN" notation → (count, sides)
  validator.go       — Range validation (count/sides 1-100)
  roller.go          — Core rolling via crypto/rand (bias-free via big.Int)
  result.go          — Result struct, String() formatting, RollNotation() orchestrator
  roll_test.go       — Unit tests
```

The `dice` package is Discord-agnostic and can be tested independently. `main.go` is thin glue that handles the Discord event loop and delegates all business logic to `dice.RollNotation()`.

## Common Commands

| Task | Command |
|---|---|
| Run tests | `go test ./dice/` |
| Run tests verbose | `go test ./dice/ -v` |
| Build | `go build -o dice-bot .` |
| Format | `go fmt ./...` |
| Vet | `go vet ./...` |
| Add dependency | `go get <module>` then `go mod tidy` |
| Run bot | `DISCORD_BOT_TOKEN=<token> go run main.go` |

## Coding Conventions

- Go files use **tabs** for indentation (per `.editorconfig`, matches `gofmt`)
- All code must be formatted with `go fmt` before committing
- Module path: `github.com/nekonenene/discord-dice-bot-qwen-claude-code`
- Bot token must be set via `DISCORD_BOT_TOKEN` environment variable (never hardcoded)
