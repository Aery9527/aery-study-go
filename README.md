# go 學習筆記

- 主要參考 [這裡](https://willh.gitbook.io/build-web-application-with-golang-zhtw) 還有對 AI 提問整理出來的內容
- 由於本身熟捻 java, 因此會與 java 對比語法/觀念內容

---

### 概觀

- [目錄結構](directory-structure.md) : 完全不同於 java 的 **src/main/java**, **src/test/java** 結構, test 是跟 source 放在一起的
- [go module](go-module.md) : 就類似 java 的 maven 或 gradle 的相依管理
- go 對於沒用到的 var, import package 會報錯, 強制開發者清理無用的 code 保持乾淨
- 僅有 public/private 兩個可見性 scope (以 package 為單位), 命名時用開頭字母大小寫決定: 大寫開頭是 public / 小寫開頭是 private
- go 在 `func` 間傳遞變數是 **pass by value**, 只有傳指標才會有 reference 的效果
- 命名原則與習慣
    - 要能明確表達 **職責** 與 **行為**
    - 善用語意明確縮寫, HTTP/ID/URL 應大寫: `UserID`/`HTTPClient`
    - var/method 習慣用 camelCase, ex: `userName`/`getUserByID()`
    - struct 習慣用 PascalCase (UpperCamelCase), 以 **領域(domain):`User`/`Order`** 為主, 避免使用 service/manager/data 等 suffix
    - interface 習慣用 **領域(domain) + 行為命名 + er 結尾:`FileReader`/`OrderCreator`**, 內包含精簡且功能明確的 func 為佳, 勿包含太多概念的 func
    - folder/package/file 社群偏好小寫+無底線命名, 但官方沒有明文禁止使用底線. [更多概念看這裡:study_package.md](package.md)
- 禁止循環相依
    - go 編譯會事先掃過所有套件的 import, 決定依賴圖 (dependency graph), 若有循環會無法決定誰先編譯誰
    - 且實務上循環相依意味著設計不良, 代表功能過度耦合, 需要將依賴的部分拆至另外 package
    - 另外像 `init()` 這種 compiler 代勞優先執行的 `func` 也會在循環相依的狀況下會有非預期情況發生

---

### 語法與特性

- [main()](cmd/study-main/study-main.go) : go 的進入點
- [basic var](cmd/study-var/study-var.go) : 基本型別
    - [point](cmd/study-point/study-point.go) : 指標
    - [nil](cmd/study-nil/study-nil.go) : 類似 java 的 null, 表示一個型別是"零值"或"空值"的概念
    - [var iota](cmd/study-iota/study-iota.go) : 類似 java enum 的概念
    - [var array](cmd/study-array/study-array.go) : 同 java array, 長度不可變
    - [var slice](cmd/study-slice/study-slice.go) : 類似 java ArrayList, 長度可變
    - [var map](cmd/study-map/study-map.go) : 同 java HashMap(無序)
    - [var struct{}](cmd/study-struct/study-struct.go) : 類似 java 16 的 record
    - [var interface](cmd/study-interface/study-interface.go) : 類似 java 的 interface, 但概念上並不是包裝"物件", 而是包裝"行為"
    - [make()](cmd/study-make/study-make.go) : 用於建立 slice/map/channel 這三種型別的記憶體分配, 回傳東西的實際上是一個 struct
    - [new()](cmd/study-new/study-new.go) : 用於分配所有型別的記憶體分配, 回傳一個指標
    - [reflect](cmd/study-reflect/study-reflect.go) : runtime 取得變數型別相關資訊, 框架的基礎大多依賴 reflect 機制
    - [type](cmd/study-type/study-type.go) : 是種可以為任何型別添加別名的宣告, EX: `type age int` 就可以宣告 age 型別的變數 `var aery age = 18`
    - [generics](cmd/study-generics/study-generics.go) : 在 `[]` 內定義泛型, EX: `func funcName[K string, V any](m map[K]V)`
- [func(){}](cmd/study-func/study-func.go) : 如何定義函數與使用
- [error handling](cmd/study-error/study-error.go) : 錯誤處理
- [process control](cmd/study-process/study-process.go) : 流程控制 (if, switch, for, goto)
- [global variable cover](cmd/study-global-variable-cover/study-global-variable-cover.go) : 全域變數覆蓋問題
- [package](cmd/study-package/study-package.go) 概念就像 java 一個 "class" 的 scope, 也就是說散在各檔案的東西只要是同個 package 就是同個
  scope. [更多概念說明](package.md)
- [godoc](cmd/study-godoc/study-godoc.go) : 類似 javadoc 的文件撰寫規範, 但沒有原生工具生成像 javadoc 的一份 html,
  但有官方的線上工具 [pkg.go.dev](https://pkg.go.dev/) 自動處理
- [go test](cmd/study-test/study-test_test.go) : 內建類似 junit 的測試,
  也可使用第三方工具 [[github.com/stretchr/testify/mock](https://github.com/stretchr/testify)] [[golang/mock](https://github.com/golang/mock)]
- [goroutine](cmd/study-goroutine/study-goroutine.go) : go 的多工處理 (multithreading)
    - [channel](cmd/study-channel/study-channel.go) : goroutine 之間的溝通管道
    - [select](cmd/study-select/study-select.go) : 多個 channel 的選擇器, 當多個 channel 都 block 時, 會等待直到某個 channel 被 unblock
    - [context](cmd/study-context/study-context.go) : 用來在多個 goroutine 之間傳遞 cancel 或 timeout 訊號用的, 其本質上是一個 chain
    - [goroutine-pool](cmd/study-goroutine-pool/study-goroutine-pool.go) : 實現一個簡單的 goroutine pool 當作練習
    - [lock]() <待整理>
    - [atomic]() <待整理>

| 關鍵字                                                   | 功能                                            | java 與之對應的                                 |
|-------------------------------------------------------|-----------------------------------------------|--------------------------------------------|
| [`var`](cmd/study-var/study-var.go) (基本型別關鍵字也在此)      | 宣告變數(可變)                                      | 直接使用型別宣告變數, 沒有一個特定的變數關鍵字                   |
| [`const`](cmd/study-var/study-var.go)                 | 宣告變數(不可變)                                     | 用 `final` 關鍵字定義                            |
| [`if` `else`](cmd/study-process/study-process.go)     | 條件判斷                                          | 一樣                                         |
| [`for`](cmd/study-process/study-process.go)           | 開啟迴圈, go 裡唯一迴圈關鍵字                             | 一樣, 但 go 的 for 語法多一點, 因為涵蓋了 java 的 `while` |
| [`range`](cmd/study-process/study-process.go)         | 搭配 `for` 使用可遍歷有個數的型別(array, slice, map, chan) | 沒有, 但類似的功能要視物件有不同的操作方式                     |
| [`break`](cmd/study-process/study-process.go)         | 中斷迴圈                                          | 一樣                                         |
| [`continue`](cmd/study-process/study-process.go)      | 跳過迴圈, 回到 `for` 判斷式                            | 一樣                                         |
| [`goto`](cmd/study-process/study-process.go)          | 流程跳躍                                          | 早期有, 但後來好像取消功能了                            |
| [`switch`](cmd/study-process/study-process.go)        | 條件判斷                                          | 一樣, 但語法有一點不同                               |
| [`fallthrough`](cmd/study-process/study-process.go)   | 搭配 `switch` 使用, 可接續執行下面的 `case`               | 預設就是這樣的行為, 與之相反需要使用 `break` 跳出 `switch`    |
| [`default`](cmd/study-process/study-process.go)       | 預設條件, `select` 跟 `swtich` 會搭配著用               | 一樣                                         |
| [`case`](cmd/study-context/study-context.go)          | 選擇條件, `select` 跟 `swtich` 會搭配著用               | 搭配 `swtich` 使用                             |
| [`select`](cmd/study-context/study-context.go)        | 在多個 `chan` 中選擇非 block 的 `chan` 執行             | 沒有類似功能                                     |
| [`func`](cmd/study-func/study-func.go)                | 宣告函數                                          | 用語法格式定義 method, 沒有特定關鍵字                    |
| [`return`](cmd/study-func/study-func.go)              | 退出 `func`                                     | 一樣                                         |
| [`defer`](cmd/study-error/study-error.go)             | 退出 `func` 前會執行區塊程式碼                           | 類似 `finally` 區塊                            |
| [`go`](cmd/study-goroutine/study-goroutine.go)        | 讓 `func` 非同步執行                                | `new Thread().start()`                     |
| [`interface`](cmd/study-interface/study-interface.go) | 定義介面                                          | 一樣, 但概念不太一樣                                |
| [`map`](cmd/study-map/study-map.go)                   | 宣告 `map` 型別                                   | `Map<String, String> map = ...`            |
| [`struct`](cmd/study-struct/study-struct.go)          | 定義 `struct` 型別                                | 類似 java 16 的 record                        |
| [`chan`](cmd/study-channel/study-channel.go)          | 宣告 `chan`(channel) 型別                         | java 有一堆物件達成的功能, go 由一個 `chan` 完成          |
| [`package`](cmd/study-package/study-package.go)       | 定義檔案的 package                                 | 有同樣關鍵字, 但意義完全不同                            |
| [`import`](cmd/study-package/study-package.go)        | 引用其他模組 `package`                              | 一樣                                         |
| [`type`](cmd/study-type/study-type.go)                | 為任何型別定義別名                                     | 要達到類似的功能必須額外定義 class, 非常麻煩                 |

---

### 社群常見框架

---

### 天生體質 Go vs Java

|         | go                  | java                    |                                                                                                                                                                                                                                                                                                                                                   |
|---------|---------------------|-------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| 編譯器     | go compiler         | javac + JVM             | - go 使用自家編寫的編譯器(gc, Go Compiler), 透過 `go build` 將 source code 編譯成 native binary 直接跑在 OS 上 <br/> - go 編譯速度極快, 大型專案也能在幾秒內編譯完成, 還有 Go modules + increment build system 加速重複編譯 <br/> - go 是 AOT(Ahead Of Time) compiler, 在執行前會一次編譯所有 source code 為 "一個" native binary 檔案 <br/> - go 指定編譯平台 `GOOS=linux GOARCH=amd64 go build -o my_app_linux main.go` |
| 執行方式    | native binary       | JVM bytecode            | - go 執行速度快, 因為直接是 native binary 直接跑在 OS 上, java 則還隔了一層 JVM                                                                                                                                                                                                                                                                                        |
| 跨平台     | cross-compilation   | JVM 負責抽象                | - go 要跨平台執行就要分別編譯, java 則是編譯一次到處透過 JVM 執行                                                                                                                                                                                                                                                                                                         |
| runtime | 自帶 runtime          | 依賴 JVM                  | - go 自帶 runtime, 但不需要額外安裝, 因為編譯時會直接打包進 binary 裡面 <br/> - java 需要安裝 JVM, 並且要確保版本相容性                                                                                                                                                                                                                                                                |
| GC      | 在自帶的 runtime 裡      | 依賴 JVM                  | 狀況同上                                                                                                                                                                                                                                                                                                                                              |
| 速度      | AOT                 | JIT                     | - go 透過 AOT 提前準備好一切, 第一次編譯好直接使用<br/> - java 透過 JIT 才跑得快, 但相對的啟動就慢                                                                                                                                                                                                                                                                                 |
| 記憶體     | 自動調節, OS 能給多少就能吃到多少 | 透過 `-Xmx` -`Xms` 限制上下用量 | - 不過 go 可以透過內部程式設定 `debug.SetMemoryLimit(512 << 20) // 限制為 512MB`<br/> - 當設置了 `SetMemoryLimit` 且用量達標, go 會優雅執行退出的流程, 否則當 OS 無法供應記憶體則會觸發 oom killer 刪除整個 process                                                                                                                                                                                   |

### 多工體質 Go vs Java

|             | go                                                         | java                                                       |
|-------------|------------------------------------------------------------|------------------------------------------------------------|
| 語法          | `go function()`                                            | `new Thread().start()`                                     |
| 效能          | 協程, 開銷低, 由 go runtime 調度                                   | OS thread, 開銷大. 但 jdk 21 的 virtual thread 也是協程模式, 開銷也降低了很多 |
| 同步          | channel                                                    | synchronized, Lock, Future, BlockingQueue                  | 
| 協調/溝通       | channel, select                                            | wait/notify, Future, ExecutorService, BlockingQueue        | 
| thread pool | 自行實作 或 [第三方 lib (ants)](https://github.com/panjf2000/ants) | ExecutorService, ThreadPoolExecutor                        | 
| 讓出 CPU 時間片  | `runtime.Gosched()`                                        | `Thread.yield()`                                           | 

- 為什麼 go 還需要 thread pool? 協程不是交給 go runtime 協調就好了嗎?
    - goroutine 是很輕沒錯, 但每個 goroutine 啟動時還是會佔用 stack(預設 2KB 起跳, 動態增長), 加上 runtime 調度與 context switch 等, 量多大一樣會OOM
    - 協程適合 I/O-heavy 系統, 若是 CPU-bound 系統 thread 的 context switch 反而成為瓶頸
- goroutine pool 功能:
    - 限流: 可限制處理 request 的 goroutine 數量, 避免 QPS 突然飆高出現 OOM
    - 資源分配: 可限制對 DB 或 API 操作的 goroutine 數量, 避免後端系統被打爆
    - 併發控制: 避免 goroutine 氾濫

---

### 為什麼 native binary 在當年不受歡迎, 而是由 java 引領風騷 20 年? 如今百轉千迴又重新站上舞台?

- 當年處於 OS 百家爭鳴時代, 所以 JVM 跨平台的方案就成了救星, 大家只要寫一次就可以在不同 OS 上執行了, 而如今基本上已經被 Windows, Unix, Linux 三分天下了.
- 而且當年 Compiler 太重/部署太複雜, 需要針對不同平台打包/靜態連結非常麻煩等問題, 不像 java 有一個 *只要有 JVM 就能跑* 的便利性. 而且如今也有了
  container 利器, 執行平台已不再是問題.
- 早期應用主要是桌面軟體/資料庫系統, 後來才著重於 web service, 而那時候 C/C++ 並沒有很好的網路 library. java 則推出 EJB, J2EE 等企業標準架構, 為當年開發
  web 系統或大型應用程式提供了良好的環境.
- C/C++ 的記憶體管理問題（pointer, memory leak）令人苦惱, Java 提供 GC（Garbage Collection）讓工程師更能專注在商業邏輯而不是在 debug segmentation fault.
- Go 是 Ken Thompson & Rob Pike(Unix 之父們)在 Google 工作時設計出來的, 他們被龐大的 C++ build system & Java boilerplate 折磨到受不了, 才設計了 Go：
  > 我要一個能像 C 一樣快, 像 Python 一樣簡潔, 還能用來寫 server 的語言!

| 時代     | 特點                           | native binary 是否合適 |
|--------|------------------------------|--------------------|
| 1990s  | 平台多, 企業要安全, 跨平台是王道           | 不合適(native 編譯太重)   |
| 2000s  | Java 壟斷企業界, VM 解放部署壓力        | JVM 模型最合適          |
| 2010s+ | 雲端化, 微服務, 平台統一, container 興起 | 開始大放異彩             |