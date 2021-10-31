# Azurite を使用してローカルの Azure Queue Storage にアクセスする

## Reference

* [azurite](https://docs.microsoft.com/ja-jp/azure/storage/common/storage-use-azurite?tabs=docker-hub#authorization-for-tools-and-sdks)

## 手順

### HTTPS 設定

* [mkcert installation](https://github.com/FiloSottile/mkcert#installation) を参考に以下の手順を実行する。

```bash
$ git clone https://github.com/FiloSottile/mkcert && cd mkcert
$ go build -ldflags "-X main.Version=$(git describe --tags)"
$ ./mkcert -install
$ ./mkcert 127.0.0.1
=> 127.0.0.1.pem and 127.0.0.1-key.pem created.
```

### Docker Compose 実行

```
$ mkdir azurite
$ mv 127.0.0.1.pem azurite/127.0.0.1.pem
$ mv 127.0.0.1-key.pem azurite/127.0.0.1-key.pem
$ docker-compose up -d
```



