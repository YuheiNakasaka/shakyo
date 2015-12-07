# Rubyのdsl、puma編

## 写経元

[http://joe-re.hatenablog.com/entry/2015/04/01/184856](http://joe-re.hatenablog.com/entry/2015/04/01/184856)

# 学び

- tap
  - レシーバの評価結果をブロックにして渡す。主にメソッドチェーンの目的で使われる。
  ```ruby
  # tapなし
  s = "hello".upcase
  p s
  ss = s.reverse
  p ss

  # tapあり
  "hello".upcase.tap{|s| p s} # HELLO
  .reverse.tap{|s| p s} # OLLEH
  ```
  - [tapのいろんな使い方](http://melborne.github.io/2012/10/29/rubys-new-control-structure-by-tap-and-break/)
- instance_eval
  - instance_evalを使っているコンテキスト(例えばclassのself)でinstance_evalの引数に渡したブロックの処理を評価し実行する
  ```ruby
  class Hoge
    def initialize(name)
      @name = name
    end
  end

  hoge = Hoge.new('razoku')
  hoge.instance_eval{p @name} #razoku
  ```
  - [instance_evalの説明](http://docs.ruby-lang.org/ja/2.1.0/method/BasicObject/i/instance_eval.html)