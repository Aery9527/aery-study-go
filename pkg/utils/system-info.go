package utils

import "os"

// GetSystemLineSeparator 沒有像 java 的 System.lineSeparator() 方便取得系統換行符號
func GetSystemLineSeparator() string {
	isWindows := os.PathSeparator == '\\'
	if isWindows {
		return "\r\n"
	} else {
		return "\n"
	}
}
