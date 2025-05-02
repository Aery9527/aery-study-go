// 模組名稱, 通常是 repo URL, 之後 import 就會基於這個路徑
module github.com/Aery9527/aery-study-go

// 指定這個 module 使用哪個 Go 版本, 不是專案執行版本, 而是 module 功能版本, 影響一些語法或行為
go 1.24

// 指定此 module 所依賴的其他 module 和其版本
//require (
//    github.com/gin-gonic/gin v1.9.1
//    golang.org/x/crypto v0.17.0
//)

// 用來替代某個模組的來源, 例如開發本地 fork
//replace github.com/original/lib => ../local-lib

// 排除某個 module 的特定版本（很少用）
//exclude github.com/some/dependency v1.2.3
