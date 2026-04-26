# Discord Dice Bot

Discord メンションでダイスロールを行う Go 言語製のボットです。

## 使い方

ボットをメンションしてダイス表記を入力すると、ダイスを振って結果を返します。

```
<@bot> 4D6
→ [3, 5, 1, 6] 合計: 15
```

### ダイス表記

- `NDM` の形式: N 個の M 面ダイスを振る
- N（個数）: 1〜100
- M（面数）: 1〜100
- 例: `4D6`（6 面ダイスを 4 個）、`2D100`（100 面ダイスを 2 個）

### エラーケース

以下の入力はエラーメッセージを返します:

- 個数・面数が 1〜100 の範囲外
- 無効な表記（`d6`、`4d`、`abc` など）

## 構築・実行

### 環境要件

- Go 1.26 以上

### セットアップ

```bash
go mod tidy
```

### 実行

```bash
DISCORD_BOT_TOKEN=<bot_token> go run main.go
```

### ビルド

```bash
go build -o dice-bot .
```

### テスト

```bash
go test ./dice/
```

## 構成

```
main.go              — エントリポイント: discordgo セッション、メッセージハンドラ
dice/
  parser.go          — ダイス表記の解析
  validator.go       — 範囲検証
  roller.go          — crypto/rand によるダイスロール
  result.go          — 結果の構造体とフォーマット
  roll_test.go       — ユニットテスト
```
