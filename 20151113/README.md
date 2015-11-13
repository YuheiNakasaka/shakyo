# Goでサンプルtwitter streamを取得

## 写経元

[http://ykicisk.hatenablog.com/entry/2015/11/11/004757](http://ykicisk.hatenablog.com/entry/2015/11/11/004757)

# 学び

- flagパッケージはコマンドラインオプションを自由に設定したり操作できて便利
- jsonのパースは正規のjson structと取得元のjsonを突き合わせてunmarshalする。シンプルで判りやすい。
- PublicStreamSampleでCがチャネルを返すようになってる
- 参照
  - [https://github.com/ChimeraCoder/anaconda/blob/4b2bf11c3f29812ce4507e50d43864e90f0499aa/streaming.go](https://github.com/ChimeraCoder/anaconda/blob/4b2bf11c3f29812ce4507e50d43864e90f0499aa/streaming.go)