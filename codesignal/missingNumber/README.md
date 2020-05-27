# benchmark

make a large array (size: 100000000) for benchmark

```shell
$ go test -run=none -bench=.  -benchmem
goos: linux
goarch: amd64
pkg: github.com/hjcian/gophercises/codesignal/missingNumber
Benchmark_missingNumber/make()_and_two_loops-8                 7         150067186 ns/op        800006422 B/op         1 allocs/op
Benchmark_missingNumber/Trapezoid-8                           19          59140395 ns/op               0 B/op          0 allocs/op
Benchmark_missingNumber/Complement-8                          16          63049450 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/hjcian/gophercises/codesignal/missingNumber  6.236s
```