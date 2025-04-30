package main

import "fmt"

// go 沒有像 java 一樣的 try-catch 錯誤處理機制
// 而是依賴 func 的多個回傳值, 回傳 error 型別來表示錯誤
// 然後 panic(這個就像 java 的 throw 了) 和 recover 機制是最後手段, 正常流程不應該出現

func main() {
	defer func() {
		// 如果這邊沒有調用 recover() 則會一路拋出到最上層
		// 看是直接中斷程序或者是到 goroutine 的最上層
		if x := recover(); x != nil {
			fmt.Printf("recover() 捕獲到 panic: %v (type:%T)\n", x, x)
		}
	}()

	fmt.Println("a1")
	_, err := goError() // 錯誤依賴回傳值處理
	if err != nil {
		fmt.Println("a2")
	} else {
		fmt.Println("a3")
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

func goError() (int, error) {
	return 9527, fmt.Errorf("error")
}
