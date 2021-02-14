# gosample

## メモ

### Frameworks and Drivers Layer

* DB
* Web
* Devices
* UI
  * cli

### Interface Adapters Layer

* Controllers
  * encrypt
  * decrypt
* Gateways
* Presenters
  * subcmd

### Application Business Rules Layer

* Use Cases:
  * read config

### Enterprise Business Rules:

* Entities
  * config

## Golangの利用開始

* `$GOPATH`を確認

~~~
echo %GOPATH%
~~~

* 上記の場所に移動

~~~
cd %GOPATH
~~~

* 作業ディレクトリ生成

~~~
mkdir gosample
~~~

* 作業ディレクトリに移動

~~~
cd gosample
~~~

* go.mod初期化

~~~
go mod init gosample
~~~

* main.go作成

~~~
# example

package main

import (
	"fmt"
)

func main() {
	fmt.Println("test")
}
~~~

* main.goの検証

~~~
go run main.go
~~~

* 自作パッケージ向けディレクトリ作成

~~~
mkdir printer
~~~

* 自作パッケージ向けディレクトリに移動

~~~
cd printer
~~~

* 自作パッケージのプログラムコード作成

~~~
# example

package printer

import "fmt"

func Print() {
	fmt.Println("Call: Print()")
}
~~~

* 自作パッケージの検証

~~~
go build
~~~

* main.goで自作パッケージを参照

~~~
# example

package main

import (
	"fmt"
	"gosample/printer"
)

func main() {
	fmt.Println("test")
	printer.Print()
}
~~~

* main.goを検証

~~~
go run main.go
~~~