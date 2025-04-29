package goroutine

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"time"
)

// 用 go 關鍵字就可以啟動一個 goroutine, 就像 java new Thread().start() 一樣 fork 出去開始執行,
// goroutine 不像 java 有 thread name, 在多執行緒下執行時透過 log 方便追蹤識別,
// 只能透過 context 攜帶 metadata, 或自行傳入值識別

func ShowGoroutine() {
	// 設定要用多少 CPU core 數來跑所有 goroutines
	runtime.GOMAXPROCS(runtime.NumCPU()) // 1.5 之後預設為機器 CPU code 數, 在這之前預設為1, 因此不需要特別設定了

	// 設定 goroutine 的 stack size, 預設是 2KB 會視情況自動增長,
	// 這可用來強制限制所有 goroutine 的 stack size, 可在測試階段找出潛在 stack overflow 風險,
	// 因此不應該用在正式環境, 除非有特殊安全/資源控制需求.
	// 這類似 java 的 -Xss 參數, 只是 java 再啟動後就無法修改, debug.SetMaxStack(n) 則可以動態修改,
	// 但是這個值只能變大不能變小, 且只影響之後產生的 goroutine, 已經產生的不受影響
	debug.SetMaxStack(256 * 1024) // 1.21+

	go goString("A") // fork 一個 goroutine 執行
	goString("B")    // 當前 goroutine 執行
	time.Sleep(100 * time.Millisecond)
}

func goString(s string) any {
	for i := 0; i < 5; i++ {
		runtime.Gosched() // 同 java 的 yield(), 讓出 CPU 時間給其他 Goroutines 使用, 自己重新排隊等待執行
		fmt.Println(s)
	}
	return nil // 如果 func 有回傳值, 則用 go 會拿不到該值, 需要該值則需要包裝起來用 channel 回傳
}
