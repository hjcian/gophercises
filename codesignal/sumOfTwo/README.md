## benchmark notes

```bash
# array size = 1000
maxcian@LAPTOP-TNBCQ9US:~/go/src/github.com/hjcian/gophercises/codesignal/sumOfTwo$ ./bench.sh 
goos: linux
goarch: amd64
pkg: github.com/hjcian/gophercises/codesignal/sumOfTwo
BenchmarkV2-8              14287             85464 ns/op           24106 B/op         40 allocs/op
BenchmarkV3N2-8             9621            123387 ns/op           24346 B/op         44 allocs/op
BenchmarkV3N4-8            10000            107172 ns/op           24349 B/op         44 allocs/op
BenchmarkV3N8-8             9556            106511 ns/op           24352 B/op         44 allocs/op
BenchmarkV3N16-8           10000            108527 ns/op           24346 B/op         44 allocs/op
PASS
ok      github.com/hjcian/gophercises/codesignal/sumOfTwo       6.491s
```

```bash
# array size = 10000000
maxcian@LAPTOP-TNBCQ9US:~/go/src/github.com/hjcian/gophercises/codesignal/sumOfTwo$ ./bench.sh 
goos: linux
goarch: amd64
pkg: github.com/hjcian/gophercises/codesignal/sumOfTwo
BenchmarkV2-8                  1        2263844900 ns/op        214783800 B/op    251884 allocs/op
BenchmarkV3N2-8                1        2134073500 ns/op        214755072 B/op    251516 allocs/op
BenchmarkV3N4-8                1        1823819700 ns/op        214738744 B/op    251309 allocs/op
BenchmarkV3N8-8                1        1685308100 ns/op        214785416 B/op    251843 allocs/op
BenchmarkV3N16-8               1        1675334100 ns/op        214690344 B/op    250659 allocs/op
PASS
ok      github.com/hjcian/gophercises/codesignal/sumOfTwo       10.249s
```