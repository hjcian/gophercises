# benchmark

make a large array (size: 100000000) for benchmark

```shell
goos: linux
goarch: amd64
pkg: github.com/hjcian/gophercises/codesignal/missingNumber
Benchmark_missingNumber/make()_and_two_loops-8                 7         150511186 ns/op        800006157 B/op         1 allocs/op
Benchmark_missingNumber/Trapezoid-8                           18          58250744 ns/op               0 B/op          0 allocs/op
Benchmark_missingNumber/Complement-8                          16          67570538 ns/op               0 B/op          0 allocs/op
Benchmark_missingNumber/Cathy-8                                1        9618034700 ns/op              32 B/op          1 allocs/op
Benchmark_missingNumber/Lisa-8                                 7         154618286 ns/op        800006157 B/op         1 allocs/op
PASS
ok      github.com/hjcian/gophercises/codesignal/missingNumber  17.062s
```