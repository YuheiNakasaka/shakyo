# gitのhook

## 写経元

[]()

# 学び

- gitのhookにはcommit前/後,push前/後などで独自のshell scriptを実行できる機能がある
  - [Gitのカスタマイズ - Gitフック](http://git-scm.com/book/ja/v2/Git-%E3%81%AE%E3%82%AB%E3%82%B9%E3%82%BF%E3%83%9E%E3%82%A4%E3%82%BA-Git-%E3%83%95%E3%83%83%E3%82%AF)
- post-commitのタイミングでslackに通知するという機能を考えた
- post-receive(pushし終えた時)にslackにに変更を通知する奴はすでにあった
  - [git-slack-hook](https://github.com/chriseldredge/git-slack-hook)
- よく考えるとcommitした内容をslackで受け取っても特にうれしいこと無い
- だからpre-commitのサンプルだけ動かした
  - 元ネタ [gitのhookでphpのテストするやつ](http://blog.manaten.net/entry/645)