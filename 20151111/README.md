#  MySQL データベースの使用

## 写経元

[https://github.com/astaxie/build-web-application-with-golang/blob/master/ja/05.2.md](https://github.com/astaxie/build-web-application-with-golang/blob/master/ja/05.2.md)

# 学び

- golangにおけるsql driverはgoが提供するインターフェースを実装する
- コロンイコール(:=)は変数の宣言と代入を同時に行うもの。res, err := stmt.Exec("razokulover", "研究開発部門", "2015-11-11")ではコロンイコールなのに、res, err = stmt.Exec("razokulover2", id)ではコロンイコールじゃないのはなぜかと思ったけど前者ではresが初出。後者は２回目の登場なので宣言が入らなかったのだった。
- ScanはsqlパッケージのRowsのメソッドで、レコードのカラムの値をScanの引数の変数に代入する
- 参照
  - [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
  - [変数の宣言と代入](http://qiita.com/ohkawa/items/e7a056708ece04d06310#%E5%A4%89%E6%95%B0%E3%81%AE%E5%AE%A3%E8%A8%80%E3%81%A8%E4%BB%A3%E5%85%A5)
  - [https://golang.org/pkg/database/sql/#Row.Scan](https://golang.org/pkg/database/sql/#Row.Scan)