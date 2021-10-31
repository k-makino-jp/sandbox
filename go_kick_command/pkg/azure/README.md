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
```



