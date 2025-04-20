package study

import "fmt"

func ShowFunc() {
	//_, remainder := divide(10, 3) // 使用 _ 可以忽略回傳值
	quotient, remainder := divide(10, 3)
	fmt.Printf("商: %d, 餘數: %d\n", quotient, remainder)

	a, b := split(10)
	fmt.Printf("x: %d, y: %d\n", a, b)

	d1, d2 := deferTest()
	fmt.Printf("d1: %d, d2: %s\n", d1, d2)

	result, err := errorTest()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("result: %s\n", result)
	}
}

func divide(a, b int) (int, int) {
	return a / b, a % b
}

// 命名回傳值, 要嘛全命名, 要嘛全都不要命名
func split(sum int) (x, y int) { // 採用命名回傳值, 若只有最後一個寫型態, 則是一個語法糖, 表示所有回傳值都是該型態
	x = sum * 4 / 9 // 整數除法運算跟 java 一樣, 結果無條件捨去, 一樣返回 int
	y = sum
	return // 會自動 return x, y
}

func deferTest() (x int, y string) {
	// defer 在當前 func 結束後, 但真正 return 前執行
	// 只可以影響 "命名回傳值" 的內容, 多個 defer 採 LIFO(後進先出) 執行
	// 使用情境:想統一加 log/metrics(但不影響回傳值)/安全清理資源/優雅地處理error/recover
	defer func() {
		x += 10
		y += "?"
	}()
	defer func() {
		x += 2
		y += "!"
	}()
	return
}

// 錯誤用回傳值處理
func errorTest() (string, error) {
	// 假設某種錯誤發生
	return "", fmt.Errorf("出錯囉")
}
