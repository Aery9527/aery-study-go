# go 語言學習筆記

- 主要參考 [這裡](https://willh.gitbook.io/build-web-application-with-golang-zhtw) 還有對 AI 提問整理出來的內容
- 由於本身熟捻 java, 因此會與 java 對比語法/觀念內容

---

### 概觀

- [目錄結構](./directory-structure.md)
- [go module]() *還未整理
- 對於沒用到的 var, import 會報錯, 強制清理無用的 code 保持乾淨
- `_` 單純一個底線的變數是 **匿名變數**, 其主要用途是當變數用不到時可避免代碼中的雜訊, java 21 也導入了這個東西
- func 可以有多回傳值, `r1, r2 := func(arg1 string, arg2, int) (return1 int, return2 string)`
- 沒有三元判斷子, 也沒有像 java `() -> {}` 匿名函數語法糖, 只能使用 `func() {}` 來表示匿名函數
- 僅有 public/private 兩個可見性 scope, 用命名開頭字母大小寫決定, 大寫開頭是 public, 小寫開頭是 private
- package 僅有一層, 其概念類比 java 的一個 "class", 也就說散在各檔的東西只要是同個 package 就是同個 scope
- package 與所在該層的 folder name 無相關性, 但習慣上會保持一致, 且同個 folder 下的檔案 package 要全部一樣

---

### 語法與特性

- [main()](./cmd/study-main/study-main.go) : go 的進入點
- [基本型別](./cmd/study-var/study-var.go)
    - [nil](./cmd/study-nil/study-nil.go) : 類似 java 的 null, 表示一個型別是"零值"或"空值"的概念
    - [var iota](./cmd/study-iota/study-iota.go) : 類似 java enum 的概念
    - [var array](./cmd/study-array/study-array.go) : 同 java array
    - [var slice](./cmd/study-slice/study-slice.go) : 類似 java ArrayList
    - [var map](./cmd/study-map/study-map.go) : 同 java HashMap(無序)
    - [var struct{}](./cmd/study-struct/study-struct.go) : 同 java 16 的 record
    - [interface](./cmd/study-interface/study-interface.go) : 類似 java 的 interface, 但概念上並不是包裝"物件", 而是包裝"行為"
    - [make()](./cmd/study-make/study-make.go) : 用於建立型別 map/slice/channel 的記憶體分配, 回傳相對應型別的初始化結構
    - [new()](./cmd/study-new/study-new.go) : 用於分配所有型別的記憶體分配, 回傳一個指標
    - [reflect](./cmd/study-reflect/study-reflect.go) : runtime 取得變數型別相關資訊, 框架的基礎大多依賴 reflect 機制
    - [type](./cmd/study-type/study-type.go) : 是一種可以為任何型別添加別名的宣告
    - [generics](./cmd/study-generics/study-generics.go) : 1.18 開始支援泛型, 有比 java 更彈性的泛型限制
- [func(){}](./cmd/study-func/study-func.go) : 如何定義函數與使用, 包含 `defer` 說明
- [流程控制](./cmd/study-process/study-process.go) : if, switch, for, goto
- [錯誤處理](./cmd/study-error/study-error.go)
- [全域變數衝突]() *待整理
- [package]() *待整理 : public/private 的展示
- [goroutine](./cmd/study-goroutine/study-goroutine.go) : go 的多工處理
    - [channel](./cmd/study-channel/study-channel.go) : goroutine 之間的溝通管道
    - [select](./cmd/study-select/study-select.go) : 多個 channel 的選擇器, 當多個 channel 都 block 時, 會等待直到某個 channel 被 unblock
    - [context](./cmd/study-context/study-context.go) : 用來在多個 goroutine 之間傳遞 cancel 或 timeout 訊號用的, 其本質上是一個 chain
    - [goroutine-pool](./cmd/study-goroutine-pool/study-goroutine-pool.go) : 實現一個簡單的 goroutine pool 當作練習
    - [lock]() *待整理
    - [atomic]() *待整理

---

### 天生體質 Go vs Java

|         | go                | java         |                                                                                                                                                                                                                                                                                                                                                   |
|---------|-------------------|--------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| 編譯器     | go compiler       | javac + JVM  | - go 使用自家編寫的編譯器(gc, Go Compiler), 透過 `go build` 將 source code 編譯成 native binary 直接跑在 OS 上 <br/> - go 編譯速度極快, 大型專案也能在幾秒內編譯完成, 還有 Go modules + increment build system 加速重複編譯 <br/> - go 是 AOT(Ahead Of Time) compiler, 在執行前會一次編譯所有 source code 為 "一個" native binary 檔案 <br/> - go 指定編譯平台 `GOOS=linux GOARCH=amd64 go build -o my_app_linux main.go` |
| 執行方式    | native binary     | JVM bytecode | - go 執行速度快, 因為直接是 native binary 直接跑在 OS 上, java 則還隔了一層 JVM                                                                                                                                                                                                                                                                                        |
| 跨平台     | cross-compilation | JVM 負責抽象     | - go 要跨平台執行就要分別編譯, java 則是編譯一次到處透過 JVM 執行                                                                                                                                                                                                                                                                                                         
| runtime | 自帶 runtime        | 依賴 JVM       | - go 自帶 runtime, 但不需要額外安裝, 因為編譯時會直接打包進 binary 裡面 <br/> - java 需要安裝 JVM, 並且要確保版本相容性                                                                                                                                                                                                                                                                |
| GC      | 在自帶的 runtime 裡    | 依賴 JVM       | 狀況同上                                                                                                                                                                                                                                                                                                                                              |
| 速度      | AOT               | JIT          | - go 透過 AOT 提前準備好一切, 第一次編譯好直接使用<br/> - java 透過 JIT 才跑得快, 但相對的啟動就慢                                                                                                                                                                                                                                                                                 |
| mem     |                   |              |                                                                                                                                                                                                                                                                                                                                                   |

### 多工體質 Go vs Java

|             | go                                                         | java                                                       |
|-------------|------------------------------------------------------------|------------------------------------------------------------|
| 語法          | `go function()`                                            | `new Thread().start()`                                     |
| 效能          | 協程, 開銷低, 由 go runtime 調度                                   | OS thread, 開銷大. 但 jdk 21 的 virtual thread 也是協程模式, 開銷也降低了很多 |
| 同步          | channel                                                    | synchronized, Lock, Future, BlockingQueue                  | 
| 協調/溝通       | channel, select                                            | wait/notify, Future, ExecutorService, BlockingQueue        | 
| thread pool | 自行實作 或 [第三方 lib (ants)](https://github.com/panjf2000/ants) | ExecutorService, ThreadPoolExecutor                        | 
| 讓出 CPU      | `runtime.Gosched()`                                        | `Thread.yield()`                                           | 

- 為什麼 go 還需要 thread pool? 協程不是交給 go runtime 協調就好了嗎?
    - goroutine 是很輕沒錯, 但每個 goroutine 啟動時還是會佔用 stack(預設 2KB 起跳, 動態增長), 加上 runtime 調度與 context switch 等, 量多大一樣會OOM
    - 協程適合 I/O-heavy 系統, 若是 CPU-bound 系統 thread 的 context switch 反而成為瓶頸
- goroutine pool 功能:
    - 限流: 可限制處理 request 的 goroutine 數量, 避免 QPS 突然飆高出現 OOM
    - 資源分配: 可限制對 DB 或 API 操作的 goroutine 數量, 避免後端系統被打爆
    - 併發控制: 避免 goroutine 氾濫

- package 不能有巢狀命名(GO哲學), 所以當多維度交錯時則應該以 **領域** 為主, 例如 user/order 跟 service/repository 交錯時, 應以 user/order 為 package
  劃分
  ```
  internal/
  ├── user/
  │   ├── service.go        // package user
  │   ├── repository.go     // package user
  │   └── model.go          // package user
  ├── order/
  │   ├── service.go        // package order
  │   ├── repository.go     // package order
  │   └── model.go          // package order
  ```
  > 這樣就以 **領域**(user/order) 劃分所有面向(SRP), 業務邏輯就可以內斂. \
  > 若以 **角色**(service/repository) 劃分 package, 業務邏輯就會被分散, 過度抽象化. \
  > **領域** **角色** 可以簡單用 **業務需求** 或 **系統需求** 來區分:
  > - user/order: 是業務邏輯劃分出來的 **領域** 概念
  > - service/repository: 是程式系統操作上劃分出來的 **角色** 概念.
- package: 一律小寫避免複數與底線, 與 folder 沒有正相關, 但習慣上一樣維持 folder 跟 package 相同, 方便管理與理解
- file: 不建議camelCase, 社群偏好小寫+無底線命名, 但官方沒有明文禁止使用底線
- var/method: camelCase(`getUserByID()`, `getUserByIDAndName()`), 只要是大寫開頭就是 public, 小寫開頭就是 private 的概念
- struct/interface: PascalCase(`OrderItem `, `UserService`), 命名結構建議為 **領域 + 行為**
- **go.mod** 定義用了哪些 module/版本是多少(類似 maven pom.xml), **go.sum** 是記錄這些 module 的內容 checksum 確保下載來的沒被改動
- Go 採用 MVS (Minimal Version Selection) 的版本解決策略, 例如 A 相依 C:v1.1, B 相依 C:v1.2, 那麼整體會使用 C:v1.2, 因為不支援多版本共存
- 為避免上述的 A 使用到 C:v1.2 而炸掉, 所以 Go 社群推崇 semver (Semantic Versioning, 語意化版本), 也就是說小版號不應該有 breaking change,
  而是不向下相容時跳大版號, 因此不同版號而炸掉是開發者的問題!
- GO method 傳遞變數時是 pass by value, 只有傳指標才會有 reference 的效果
- 傳指標時會做 escape analysis (逃逸分析), 如果其內容離開 scope 會被放到 heap 上, 後續自動 GC

---

### 為什麼 native binary 在當年不受歡迎, 而是由 java 引領風騷 20 年? 百轉千迴又重新站上舞台?

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