# goji読む、サンプル使う
今日もサンプル打つしかできなかった。移動中にこれ読んだ。[13 どのようにしてWebフレームワークを設計するか](https://astaxie.gitbooks.io/build-web-application-with-golang/content/ja/13.0.html)

## 写経元

[https://github.com/zenazn/goji](https://github.com/zenazn/goji)

# 学び

- Goでアプリ作る時にフレームワークを使おうと思ったのだがまだRubyでいうRails、PythonでいうDjangoみたいなデファクトはないみたい。
- Revel,Beego,Gojiとか色々あるけどパッとみ重すぎず、適度に抽象化されてて機能も十分という意味でGojiを使ってみた。
- Router読んでみたけどnet/httpに準拠してるのはいい感じ。
- Einhorn(rubyのgem。新しいソケットを作って既存の動いてるプロセスからコピーしてbindする)を使うとゼロダウンタイムでサーバーのリロードができるみたいなのでそこも実際に運用するとき良さそう
  - [Einhornでサーバのリロードやってる](http://takeshiyako.blogspot.jp/2015/10/golang-goji-einhorn-graceful-restart.html)