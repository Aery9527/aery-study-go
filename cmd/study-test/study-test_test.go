package main

import (
	"fmt"
	"os"
	"testing"
)

// 測試檔案必須以 "_test.go" 結尾
// 測試函式必須以 "Test" 開頭, 簽名為 func(t *testing.T)

// TestMain 這是一個特殊的函式, 一個 package 裡只能有一個 TestMain,
// 當定義了這個函式時, 它會取代預設的 go test 執行流程,
// 在所有測試開始之前執行初始化邏輯, 並在所有測試完成之後執行清理邏輯
func TestMain(m *testing.M) {
	// 做一些初始化行為, 例如 DB 連線
	fmt.Println("test initial...")
	code := m.Run() // 執行測試
	fmt.Println("test finish...")
	// 做一些收尾行為, 例如 關閉 DB 連線
	os.Exit(code)
}

// TestAdd 單元測試
func TestAdd(t *testing.T) {
	result := add(2, 3) // XXX test target
	if result != 5 {
		t.Errorf("expected 5, got %d", result)
	}
}

// BenchmarkAdd 效能測試
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(1, 2) // XXX test target
	}
}
