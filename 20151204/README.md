# rubyのthread

本番コードで使ったのでコードは載せられない

## 写経元

[]()

# 学び

- 状況としてはとあるアプリ内で外部連携先(redis)にやや特殊な接続をする実装をしなければならなかった
- 素直に同期的な実装をするとアプリのレスポンスはそことの連携部分で40%くらいを占めることに
- そこでthreadで非同期に実行させることにした
- Threadを使う際の注意は
  - RubyのThreadはthread内部で処理が通常 or 例外によって終了したらkillされる
  - threadを増やしすぎないこと => 最悪、memoryを使い切る
  - thread内からDBと接続する場合はActiveRecord::Base.connection_pool.with_connectionにブロックを渡してその内部で処理を書く。そうしないとthread内でDBへの接続ごとに別コネクションを貼ることになり、最悪、DBのmax_connectionに達して死ぬ
  - thread外のリソース(例えば配列)に複数のthread内からアクセスするとリソースの同じアドレスを上書きしてしまうことがある
  - 基本的にはthread内で発生した例外はその親プロセスでは捕捉できない。どうしてもthread内の例外をつかみたい時は、下記のようにするとthread内の例外を親にパスしてもう一度例外を発してくれる。
  ```ruby
  begin
    t = Thread.new do
      # something error
      raise 'error occured'
    end
    t.join
  rescue => e
    Rails.logger.fatal(e) # error occured
  end
  ```
- 参照
  - [class Thread](http://docs.ruby-lang.org/ja/2.1.0/class/Thread.html)
  - [Rubyの並行処理で学んだこと――パート１(翻訳版)](http://www.engineyard.co.jp/blog/2013/ruby-concurrency/)
  - [【Ruby on Rails】非同期処理についてその②。　〜マルチスレッドとかマルチプロセスを実装してみた〜](http://shirusu-ni-tarazu.hatenablog.jp/entry/2013/07/02/042448)
