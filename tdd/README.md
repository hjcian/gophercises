# Source
- https://github.com/quii/learn-go-with-tests
- https://quii.gitbook.io/learn-go-with-tests/

# Notes
- GODOC
    - 使用 `godoc -http=:6060` 查看 godoc 
        - 前提 `sudo apt install golang-golang-x-tools` 才能使用 godoc
    - 定義 ExampleFunc ，並包含 // Output: 來要求 test 也要執行 ExampleFunc，讓測試例正確
- `reflect.DeepEqual`
    - 複雜的類型深度比較，但只能在 runtime 時檢查型別，不會在編譯時期檢查是否為可比較的型別
- `t.Errorf("%#v got %.2f want %.2f", shape, got, want)`
    - %v - 只印出 value
    - %#v - 一併印出 struct/field name (debug 好用)

# Progress
- https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/arrays-and-slices
