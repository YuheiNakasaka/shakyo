# golangで複数単語のtweet検索をgoroutineでやる

## 写経元

[環境変数をmapにするやつ](https://coderwall.com/p/kjuyqw/get-environment-variables-as-a-map-in-golang)
[twitterで検索する](https://github.com/ChimeraCoder/anaconda)
[loopでgoroutine使うやつ](http://www.kaoriya.net/blog/2013/07/08/)

# 学び

- goの環境変数を扱うのはやや面倒。いちいちkeyvalのmapにセットしないと使い物にならん。
- slice便利
- sync版(5.425194秒)とasync版(4.459142秒)書いたけどあんまり速度は変わらなかった。ややasync版が速かったくらい。単語が10個だったからだと思う。もっと増やせば如実に結果でるんだろうけど、そうすると今度はツイートの制限に引っかかる。ツイート検索を並列実行してもあんまり意味ないのかも。
- 参照
  - http://qiita.com/sawanoboly/items/76abd295b32123c4568d
  - http://golang.jp/pkg/strings
  - https://coderwall.com/p/kjuyqw/get-environment-variables-as-a-map-in-golang