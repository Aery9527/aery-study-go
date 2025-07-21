package main

import "fmt"

// go 沒有像 java 一樣的 try-catch 錯誤處理機制
// 而是依賴 func 的多個回傳值, 用回傳的 error 型別來表示錯誤

// 而 panic(這個就像 java 的 throw 了) 和 recover 機制是最後手段, 正常流程不應該出現
// 因為在任何一個 goroutine 中, 如果發生 panic 沒有被 recover, 整個 process 會直接中斷, exit code 是 2
// 所以 panic 會是"超級嚴重的錯誤", 程度比 java OutOfMemoryError 還要嚴重的概念

func main() {
	defer func() {
		// 如果這邊沒有調用 recover() 則會一路拋出到最上層
		// 看是直接中斷程序或者是到 goroutine 的最上層
		if x := recover(); x != nil {
			fmt.Printf("recover() 捕獲到 panic: %v (type:%T)\n", x, x)
		}
	}()

	fmt.Println("a1")
	_, err := goError0() // 錯誤依賴回傳值處理
	if err != nil {
		fmt.Printf("a2 : %v\n", err)
	} else {
		fmt.Printf("a3 : %v\n", err)
	}

	fmt.Println("b1")
	goPanic()         // panic 會強制中斷執行(就像 java 的 throw 一樣), 直到被 recover 處理
	fmt.Println("b2") // 不會執行

	defer func() {
		// 這裡調用 recover() 並沒有效果, 原因是在觸發 panic 之前並沒有定義這個 defer
	}()
}

func goPanic() {
	fmt.Println("c1")
	panic("oops")     // 就像 java 的 throw exception
	fmt.Println("c2") // 不會執行
}

func goError0() (int, error) {
	_, err := goError1()
	return 0, fmt.Errorf("goError0: %w", err) // 包裝錯誤, %w 可以讓錯誤被包裝起來, 方便後續處理
}

func goError1() (int, error) {
	// 錯誤訊息有規範
	// 1.開頭不能是大寫字母(因為錯誤通常是被包在其他上下文中, 不需要大寫)
	// 2.不能用標點符號結尾, 例如："." "!" "?"
	return 9527, fmt.Errorf("error1")
}
