# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

Single-file Discord dice bot (`main.go`) using `github.com/bwmarrin/discordgo`. Go 1.26.2.

## Commands

- **Run**: `DISCORD_BOT_TOKEN=<token> go run main.go`
- **Build**: `go build -o discord-dice-bot`
- **Format**: `go fmt main.go`
- **Check**: `go vet ./...`

## Editorconfig

`.editorconfig` applies Go-specific rules: tab indent, 4-space indent size. Use `go fmt` for final formatting.

## Architecture

All logic is in `main.go`:

- `parseDice()` - regex-based parser for `NDM` format (e.g. `4D6`). Validates count and faces are 1~100.
- `rollDice()` - rolls N dice with M faces using `math/rand`.
- `formatResult()` - formats results as `NDM: [x,y,z] (合計: sum)`.
- `sendEmbed()` - helper that sends a `discordgo.MessageEmbed` to a channel.
- `main()` - sets up discordgo session, registers `MessageCreate` handler that triggers on bot mention + dice string, listens with `select {}`.

The bot responds only when mentioned (matches `<@id>` or `<@!id>` prefix), ignoring its own messages.
