package utils

import (
	"fmt"
	"runtime"
)

// PrintWhere 取得當前執行的檔案名稱與行號
func PrintWhere() {
	PrintWhereAt(1)
}

// PrintWhereAt 取得當前執行的檔案名稱與行號, at 代表要往上找幾層
func PrintWhereAt(at int) {
	_, file, line, _ := runtime.Caller(at + 1)
	fmt.Printf("at %s:%d\n", file, line)
}
