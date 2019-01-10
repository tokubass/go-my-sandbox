3行まとめ
- build tagでソースコードをファイル単位で切り替えることで、同名の関数を定義してもエラーにならない。これを利用してテスト時はmockコードに切り替える。
- tagはファイルの先頭行にコメントで記載
- go test --tag 'test' でtag指定できる ( go runやbuildのときは必要ない。`!test`にしているから）

production実装とmockが乖離しないようにinterfaceを定義しておくことを強く推奨。

```console
$ go run main.go
production

$ make test
go test --tags 'test' -v ./...
ok  	my/go-my-sandbox/mock_with_build_tag	0.017s
?   	my/go-my-sandbox/mock_with_build_tag/client	[no test files]
?   	my/go-my-sandbox/mock_with_build_tag/domain	[no test files]
```

build tagの解説は以下
https://golang.org/pkg/go/build/#hdr-Build_Constraints

build tagとpackage名の間に空行が必要なことに注意（lintで防げると思いますが）

# memo

Golang: Don’t afraid of makefiles
https://frasco.io/golang-dont-afraid-of-makefiles-785f3ec7eb32
