# 簡単なhttpサーバー

## 写経元

[http://qiita.com/taizo/items/bf1ec35a65ad5f608d45](http://qiita.com/taizo/items/bf1ec35a65ad5f608d45)
[https://github.com/golang-samples/http/tree/master/file](https://github.com/golang-samples/http/tree/master/file)

# 学び

- httpを使ったサーバー、割と簡単にかけてよさそ
- 型にServeHTTPを生やしたhandlerをかけばhttp.Handleのhandlerに登録できる。便利。
- http以外のプロトコルに応じるサーバーもhttp.RegisterProtocolを使えばかける