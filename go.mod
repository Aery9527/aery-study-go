// 模組名稱, 通常是 repo URL, 之後 import 就會基於這個路徑
module aery-study-go

// 指定這個 module 使用哪個 Go 版本, 不是專案執行版本, 而是 module 功能版本, 影響一些語法或行為
go 1.24

//require (
//    github.com/gin-gonic/gin v1.9.1
//    golang.org/x/crypto v0.17.0
//)

// 用來替代某個模組的來源, 例如開發本地 fork
//replace github.com/original/lib => ../local-lib

// 排除某個 module 的特定版本（很少用）
//exclude github.com/some/dependency v1.2.3

require github.com/cockroachdb/errors v1.12.0

require (
	github.com/cockroachdb/logtags v0.0.0-20230118201751-21c54148d20b // indirect
	github.com/cockroachdb/redact v1.1.5 // indirect
	github.com/getsentry/sentry-go v0.27.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
)
