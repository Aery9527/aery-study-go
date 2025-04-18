package util

import (
	"fmt"
	"os"
)

func init() {
	fmt.Println("util.go init()...")
}

// 沒有像 java 的 System.lineSeparator() 方便取得系統換行符號
func GetSystemLineSeparator() string {
	if os.PathSeparator == '\\' { // Windows 系統
		return "\r\n"
	}
	return "\n"
}
