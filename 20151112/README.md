# Goでルーターを書く

## 写経元

[https://github.com/bmizerany/pat](https://github.com/bmizerany/pat)

# 学び

- Goのカスタムルーターはルーティングパターンとそれに対応するServeHTTPメソッドハンドラを定義したServeHTTPメソッドを実装したstructを作れば良い
- ルーターは色々あるが、違いは基本的にルーティングのマッチ方法やパターンの定義書式の違いしかない
- 参照
  - [Goのhttpパッケージ詳細](https://github.com/astaxie/build-web-application-with-golang/blob/master/ja/03.4.md)