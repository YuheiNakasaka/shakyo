# Twitter Streaming APIを触る

## 写経元

[https://github.com/sferik/twitter/blob/master/examples/Configuration.md](https://github.com/sferik/twitter/blob/master/examples/Configuration.md)

# 学び

- Streaming APIにバグがあって、時期バージョンで治るらしい。今はこの[issue](https://github.com/sferik/twitter/issues/709)のようにgemをハードコードする必要あり
- RTされたTweetもStreamに乗って表示される。その際は「RT @そのツイートした人のscreen_name ツイート本体」のような形式。
