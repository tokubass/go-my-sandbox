3行まとめ
- build tagでソースコードをファイル単位で切り替えることで、同名の関数を定義してもエラーにならない。これを利用してテスト時はmockコードに切り替える。
- tagはファイルの先頭行にコメントで記載
- `go test --tag 'test'` でtag指定できる ( go runやbuildのときは必要ない。`!test`にしているから）

production実装とmockが乖離しないようにinterfaceを定義しておくことを強く推奨。

###  --tag 'test' の場合
`main_test.go`と`client/mock_client.go`が選ばれる

mockの実装をグローパル変数に保持しているの単純な実装。parallelに実行されるとダメ。

### --tag 'use_lock'の場合
`main_lock_test.go`と`client/mock_client_lock.go`が選ばれる

parallel対応のため`mutex.Lock`を使用した。

LockするのでDBのmockなど頻繁に使うものには速度面で辛い。

またテスト側でUnlockする必要がある。

## 実行結果
```console
$ go run main.go
production

$ make test-all
go test --tags 'test' -v ./...
=== RUN   TestFoo
--- PASS: TestFoo (0.00s)
PASS
ok  	my/go-my-sandbox/mock_with_build_tag	0.018s
?   	my/go-my-sandbox/mock_with_build_tag/client	[no test files]
?   	my/go-my-sandbox/mock_with_build_tag/domain	[no test files]
go test --tags 'use_lock' -v ./...
=== RUN   TestFoo
=== RUN   TestFoo/test_1
=== RUN   TestFoo/test_2
=== RUN   TestFoo/test_3
=== RUN   TestFoo/test_4
=== RUN   TestFoo/test_5
=== RUN   TestFoo/test_6
=== RUN   TestFoo/test_7
=== RUN   TestFoo/test_8
--- PASS: TestFoo (0.00s)
    --- PASS: TestFoo/test_1 (0.00s)
    --- PASS: TestFoo/test_4 (0.00s)
    --- PASS: TestFoo/test_2 (0.00s)
    --- PASS: TestFoo/test_5 (0.00s)
    --- PASS: TestFoo/test_8 (0.00s)
    --- PASS: TestFoo/test_3 (0.00s)
    --- PASS: TestFoo/test_7 (0.00s)
    --- PASS: TestFoo/test_6 (0.00s)
PASS
ok  	my/go-my-sandbox/mock_with_build_tag	0.018s
?   	my/go-my-sandbox/mock_with_build_tag/client	[no test files]
?   	my/go-my-sandbox/mock_with_build_tag/domain	[no test files]
```

build tagの解説は以下
https://golang.org/pkg/go/build/#hdr-Build_Constraints

build tagとpackage名の間に空行が必要なことに注意（lintで防げると思いますが）

# memo

Golang: Don’t afraid of makefiles
https://frasco.io/golang-dont-afraid-of-makefiles-785f3ec7eb32
