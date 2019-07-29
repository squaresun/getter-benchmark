# Getter-benchmark

```shell
$ go test -bench=. .
goos: darwin
goarch: amd64
pkg: github.com/squaresun/getter-benchmark
BenchmarkServeInterface-8        3000000               410 ns/op
BenchmarkServeStruct-8          20000000               76.5 ns/op
PASS
ok      github.com/squaresun/getter-benchmark   3.270s
```
