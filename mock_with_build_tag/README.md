build tagでソースコードを切り替えることで、同名の関数を定義してもエラーにならない。これを利用してテスト時はmockコードに切り替える

$ go run --tags 'test' main.go
test
$ go run --tags 'production' main.go
production
