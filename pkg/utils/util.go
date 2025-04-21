package utils

import (
	"fmt"
	"os"
	"runtime"
)

// 沒有像 java 的 System.lineSeparator() 方便取得系統換行符號
func GetSystemLineSeparator() string {
	if os.PathSeparator == '\\' { // Windows 系統
		return "\r\n"
	}
	return "\n"
}

func PrintWhere() {
	PrintWhereLevel(1)
}

func PrintWhereLevel(level int) {
	_, file, line, _ := runtime.Caller(level + 1)
	fmt.Printf("at %s:%d\n", file, line)
}
