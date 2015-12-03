# fake_sqs

仕事で書いたのでコードは載せられない

## 写経元

[]()

# 学び

- sqsと接続してる部分のtestを書くためにfake_sqsを検討した
- 他にElasticMQがあるがそっちはjavaで動かすのでちょっと大げさな感じがする
  - [http://qiita.com/halhide/items/59665e7c5ee1d2b20603](http://qiita.com/halhide/items/59665e7c5ee1d2b20603)
- in memoryで使えるのでテスト(開発)用キューとしては十分な感じ
- しかしながらrspec+sporkとの連携がまだうまくいかないとこもあって検証中
  - [https://github.com/iain/fake_sqs](https://github.com/iain/fake_sqs)
  - [http://qiita.com/iemon7stars/items/d4efdd8872d287906d29](http://qiita.com/iemon7stars/items/d4efdd8872d287906d29)
  - [http://dev.classmethod.jp/cloud/ruby-amazon-sns-sqs/](http://dev.classmethod.jp/cloud/ruby-amazon-sns-sqs/)