# Qwen 3.5 on Claude Code による Discord ダイスボット

[Qwen3.6-35B-A3B-UD-MLX-4bit](https://huggingface.co/unsloth/Qwen3.6-35B-A3B-UD-MLX-4bit) をローカルLLMとして使用し Claude Code を動かした。

指示は以下の通りおこない、その後の修正はおこなっていない。

```
Discord bot を作りたいです。Go言語です。
ダイスロールの結果を返すbotです。

例えば「4D6」のようにメンションされると、6面ダイスを4個振った結果を返します。
「2D100」のようにメンションされると、100面ダイスを2個振った結果を返します。

ダイスの面数は1~100までの整数、ダイスの個数は1~100までの整数で、
それ以外を指定された場合はエラーメッセージを返すようにしたいです。

コーディングルールは ~/Programs/.editorconfig を参照してください。
コードは必ず go fmt で整形してください。
```

## 良かった点

- メンション時にのみメッセージを返すことはできている

## 悪かった点

- 正常に動かない！！
- `discordgo` のリポジトリ名を間違え `go get github.com/bwmarr0/discordgo` を実行しようとして失敗し、迷走し続けた（正しいリポジトリは `github.com/bwmarrin/discordgo` ）
  - `cp ~/.go/pkg/mod/github.com/bwmarrin/discordgo@v0.29.0 ~/.go/pkg/mod/github.com/bwmarr0/discordgo@v0.29.0` をして無理やり解決しようとしたので、正しいリポジトリを教えることにした

## 作成した Discord bot とのやり取りの様子

> +++++++++++++ Qwen No. 3 +++++++++++++

> @shamiko-bot 20D100

ダイスを指定してください。例: `4D6`

> 4D6

> @shamiko-bot 4D6

ダイスを指定してください。例: `4D6`
