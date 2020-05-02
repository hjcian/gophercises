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
- `t.Helper()`
    - 告知 testing 工具這個 function 是 helper function，錯誤追蹤時不需要深入到這裡
- `t.Errorf("%#v got %.2f want %.2f", shape, got, want)`
    - %v - 只印出 value
    - %#v - 一併印出 struct/field name (debug 好用)
- mocking 心法
    1. 撰寫實作與測試時，就可以思考未來如果要重構這一段，是否需要更改測試？如果需要更改測試，表示此實作太專注於**細節**而非**行為**
    2. 測試應專注在**行為**，而非程式碼的實作細節。確實地感受它。能夠提高程式碼撰寫品質、可讀性、可維護性等軟體工程層面的好處
    3. 故若測試都是針對程式碼實作細節，會造成未來重構程式碼時也要重構一堆測試，這的確有害
    4. 一個測試使用了三次 mocking 技巧表示程式碼可能太複雜了
    

# Progress
- https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/concurrency
