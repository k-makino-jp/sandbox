# 自己署名認証局のサーバー証明書を生成する

## 自己署名認証局を構築する

* 作業ディレクトリを作成する

~~~
$ mkdir -p /etc/pki/CA/private
~~~

* 自己署名認証局の秘密鍵を作成する

~~~
$ openssl genrsa -aes256 -out /etc/pki/CA/private/cakey.pem 2048
~~~

* 証明書発行要求ファイルを作成する

~~~
$ openssl req -new \
-key /etc/pki/CA/private/cakey.pem \
-out /etc/pki/CA/cacert.csr
~~~

* 自己署名証明書を作成する

~~~
$ openssl x509 -days 365 -req \
-in      /etc/pki/CA/cacert.csr \
-signkey /etc/pki/CA/private/cakey.pem \
-out     /etc/pki/CA/cacert.pem
~~~

* 自己署名認証局の運用に必要なファイルを作成する

~~~
$ touch /etc/pki/CA/index.txt
$ echo 00 > /etc/pki/CA/serial
~~~

## 自己署名認証局でサーバー証明書を発行する

* サーバーの秘密鍵を作成する

~~~
$ openssl genrsa -aes256 -out /etc/pki/tls/private/privkey.pem 2048
~~~

* 証明書発行要求ファイルを作成する

~~~
$ openssl req -new \
-key /etc/pki/tls/private/privkey.pem \
-out /etc/pki/tls/certs/domain_name.csr
~~~

* サーバー証明書を作成する

~~~
$ openssl ca -days 365 -policy policy_anything \
-in  /etc/pki/tls/certs/domain_name.csr \
-out /etc/pki/tls/certs/domain_name.crt.pem
~~~

## パスフレーズなしの秘密鍵を作成する

~~~
$ openssl rsa \
-in /etc/pki/tls/private/privkey.pem \
-out /etc/pki/tls/private/privkey-nopass.pem
~~~
