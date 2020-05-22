# benchmark

make a large array (size: 100000000) for benchmark

```shell
$ go test -run=none -bench=.  -benchmem
goos: linux
goarch: amd64
Benchmark_missingNumber/make()_and_two_loops-6                 1        6311308578 ns/op        800006144 B/op         1 allocs/op
Benchmark_missingNumber/Trapezoid-6                           19          58002254 ns/op               0 B/op          0 allocs/op
Benchmark_missingNumber/Complement-6                          19          59698539 ns/op               0 B/op          0 allocs/op
PASS
ok      _/home/maxcian/projects/gophercises/codesignal/missingNumber    13.638s
```