|         | go                | java         |                                                                                                                                                                                                                                                                                                                                                   |
|---------|-------------------|--------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| 編譯器     | go compiler       | javac + JVM  | - go 使用自家編寫的編譯器(gc, Go Compiler), 透過 `go build` 將 source code 編譯成 native binary 直接跑在 OS 上 <br/> - go 編譯速度極快, 大型專案也能在幾秒內編譯完成, 還有 Go modules + increment build system 加速重複編譯 <br/> - go 是 AOT(Ahead Of Time) compiler, 在執行前會一次編譯所有 source code 為 "一個" native binary 檔案 <br/> - go 指定編譯平台 `GOOS=linux GOARCH=amd64 go build -o my_app_linux main.go` |
| 執行方式    | native binary     | JVM bytecode | - go 執行速度快, 因為直接是 native binary 直接跑在 OS 上, java 則還隔了一層 JVM                                                                                                                                                                                                                                                                                        |
| 跨平台     | cross-compilation | JVM 負責抽象     | - go 要跨平台執行就要分別編譯, java 則是編譯一次到處透過 JVM 執行                                                                                                                                                                                                                                                                                                         
| runtime | 自帶 runtime        | 依賴 JVM       | - go 自帶 runtime, 但不需要額外安裝, 因為編譯時會直接打包進 binary 裡面 <br/> - java 需要安裝 JVM, 並且要確保版本相容性                                                                                                                                                                                                                                                                |
| GC      | 在自帶的 runtime 裡    | 依賴 JVM       | 狀況同上                                                                                                                                                                                                                                                                                                                                              |
| 速度      | AOT               | JIT          | - go 透過 AOT 提前準備好一切, 第一次編譯好直接使用<br/> - java 透過 JIT 才跑得快, 但相對的啟動就慢                                                                                                                                                                                                                                                                                 |

- 執行入口一定要 `package main`, 且該 package 只能有一個 `main` 函式, 否則會報錯
- `package main` 是特殊的 package, 只能當作程式進入點, 無法被其他 package 引用
- package 跟 folder 沒有正相關, 但習慣上一樣維持 folder 跟 package 相同, 方便管理與理解
- 沒有三元判斷子, 沒有 `() -> {}` 的 lambda 語法糖
-

|         | go             | java            |                                       |
|---------|----------------|-----------------|---------------------------------------|
| package | 跟 folder 沒有正相關 | 必須跟 folder 相同結構 | 但習慣上一樣維持 folder 跟 package 相同, 方便管理與理解 |

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