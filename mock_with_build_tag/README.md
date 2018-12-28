build tagでソースコードを切り替えることで、同名の関数を定義してもエラーにならない。これを利用してテスト時はmockコードに切り替える

```console
$ go run --tags 'test' main.go
test
$ go run --tags 'production' main.go
production
```

build tagの解説は以下
https://golang.org/pkg/go/build/#hdr-Build_Constraints

build tagとpackage名の間に空行が必要なことに注意（lintで防げると思いますが）
