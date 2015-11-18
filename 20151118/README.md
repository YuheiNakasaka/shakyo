# jsonとsortのサンプル
今日時間なくて手抜き

## 写経元

[http://blog.restartr.com/2014/08/13/golang-json-marshal-unmarshal/](http://blog.restartr.com/2014/08/13/golang-json-marshal-unmarshal/)
[https://gobyexample.com/sorting](https://gobyexample.com/sorting)

# 学び

- 基本のsortは楽。mapのkeyでソートとかどうやるか。追記)=>やってみた。keyのスライスを作ってそれをソートして、その結果をfor rangeで新規にmapにつっこみ直す。
- jsonのタグづけで`json:hoge`みたいにダブルクォーテーションがない場合はエラーにならずに無視されて、keyが構造体のフィールドを使うみたいなので注意が必要
- var hoge map[string]int はnilなので代入できない