package study

import "fmt"

func ShowPanic() {
	// go 沒有像 java 一樣的 exception 機制
	// 依靠的是 panic 和 recover 機制
	// 這應當是最後手段, 正常流程不應該出現

	defer func() {
		// 如果這邊沒有調用 recover() 則會一路拋出到最上層
		// 看是直接中斷程序或者是到 goroutine 的最上層
		if x := recover(); x != nil {
			fmt.Printf("recover() 捕獲到 panic: %v (type:%T)\n", x, x)
		}
	}()

	fmt.Println("b1")
	goPanic() // panic 會強制中斷執行(就像 java 的 throw 一樣), 直到被 recover 處理
	fmt.Println("b2")

	defer func() {
		// 這裡調用 recover() 並沒有效果, 原因是在觸發 panic 之前並沒有定義這個 defer
	}()
}

func goPanic() {
	fmt.Println("c1")
	panic("oops") // 這有點像 java throw 的感覺
	fmt.Println("c2")
}
