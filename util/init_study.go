package util

import "fmt"

// init() 是 Go 語言的特殊函式, 在 main() 函式之前執行, 主要用來初始化套件或變數
// init() 函式不需要參數, 也不會回傳值, 可以有多個 init() 函式, 但會按照出現的順序執行
// init() 函式在 package 被載入時自動執行, 這樣可以確保在使用 package 之前, 所有的初始化工作都已經完成
func init() {
	fmt.Println("init_study.go init()...")
}
