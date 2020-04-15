benchmark notes
```bash
maxcian@LAPTOP-TNBCQ9US:~/go/src/github.com/hjcian/gophercises/codesignal/sumOfTwo$ ./bench.sh 
goos: linux
goarch: amd64
pkg: github.com/hjcian/gophercises/codesignal/sumOfTwo
BenchmarkV1-8                246           4125539 ns/op             199 B/op          0 allocs/op
BenchmarkV2-8               5428            219467 ns/op           24111 B/op         40 allocs/op
BenchmarkV3N2-8             2362            486274 ns/op           48408 B/op         85 allocs/op
BenchmarkV3N4-8             1875            643828 ns/op           48384 B/op         85 allocs/op
BenchmarkV3N8-8             2548            523778 ns/op           48412 B/op         85 allocs/op
PASS
ok      github.com/hjcian/gophercises/codesignal/sumOfTwo       6.609s
```