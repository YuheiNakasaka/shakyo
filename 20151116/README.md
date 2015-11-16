# Goでgrep実装

## 写経元

[https://github.com/bose999/GoGrep/blob/master/gogrep.go](https://github.com/bose999/GoGrep/blob/master/gogrep.go)

# 学び

- mapつかない奴は配列。map[]int{}とかはいわゆるハッシュ。まだ間違える。
- Printf,Println
- scanner.NewScannerしてscanner作ってScan()でイテレートしてText()で行の文字を読み込む流れ
- 無名goroutine関数即時実行慣れてない
- channelはcloseされると0を返す。for{}でchannelが閉じるのを待って0の時にbreakするという書き方があるのか。
- 参照
  - [https://golang.org/pkg/bufio/#Scanner.Text](https://golang.org/pkg/bufio/#Scanner.Text)
  - [Goでchannelがcloseしてるかどうか知りたい というアンチパターン](https://beatsync.net/main/log20150325.html)