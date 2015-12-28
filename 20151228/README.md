# Let's EncryptでSSL設定をnginxで行う

## 参考

- [https://saku.io/using-lets-encrypt-with-nginx/](https://saku.io/using-lets-encrypt-with-nginx/)
- [http://inside.pixiv.net/entry/2015/12/10/153114](http://inside.pixiv.net/entry/2015/12/10/153114)

インストール

```
cd ~
git clone https://github.com/letsencrypt/letsencrypt.git
cd letsencrypt
```

init

```
./letsencrypt-auto
```

メアド登録と規約に承認

```
./letsencrypt-auto certonly --webroot --webroot-path WEBROOTのPATH -d ドメイン

例)
./letsencrypt-auto certonly --webroot --webroot-path /var/www/hoge/public -d hoge.hogehoge.com
```

DH鍵

```
sudo mkdir /etc/nginx/ssl # フォルダがない場合
sudo openssl dhparam 2048 -out /etc/nginx/ssl/dhparam.pem
```

nginxのconfファイルに設定

```
ssl_certificate /etc/letsencrypt/live/<ドメイン>/fullchain.pem;
ssl_certificate_key /etc/letsencrypt/live/<ドメイン>/privkey.pem;
ssl_session_timeout 1d;
ssl_session_cache shared:SSL:50m;
# 生成したDHパラメータを使用
ssl_dhparam /etc/nginx/ssl/dhparam.pem;
# 暗号化スイートを明示的に指定する
ssl_prefer_server_ciphers on;
ssl_ciphers 'ECDHE-RSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-ECDSA-AES256-GCM-SHA384:DHE-RSA-AES128-GCM-SHA256:DHE-DSS-AES128-GCM-SHA256:kEDH+AESGCM:ECDHE-RSA-AES128-SHA256:ECDHE-ECDSA-AES128-SHA256:ECDHE-RSA-AES128-SHA:ECDHE-ECDSA-AES128-SHA:ECDHE-RSA-AES256-SHA384:ECDHE-ECDSA-AES256-SHA384:ECDHE-RSA-AES256-SHA:ECDHE-ECDSA-AES256-SHA:DHE-RSA-AES128-SHA256:DHE-RSA-AES128-SHA:DHE-DSS-AES128-SHA256:DHE-RSA-AES256-SHA256:DHE-DSS-AES256-SHA:DHE-RSA-AES256-SHA:ECDHE-RSA-DES-CBC3-SHA:ECDHE-ECDSA-DES-CBC3-SHA:AES128-GCM-SHA256:AES256-GCM-SHA384:AES128-SHA256:AES256-SHA256:AES128-SHA:AES256-SHA:AES:CAMELLIA:DES-CBC3-SHA:!aNULL:!eNULL:!EXPORT:!DES:!RC4:!MD5:!PSK:!aECDH:!EDH-DSS-DES-CBC3-SHA:!EDH-RSA-DES-CBC3-SHA:!KRB5-DES-CBC3-SHA';
# POODLE対策でSSLv1〜3は有効にしない
ssl_protocols TLSv1 TLSv1.1 TLSv1.2;
# HSTSヘッダを設定 （SSL接続を強制する場合のみ）
#add_header Strict-Transport-Security 'max-age=31536000; includeSubDomains;';
```
