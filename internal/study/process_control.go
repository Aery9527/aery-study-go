package study

import (
	"fmt"
	"os"
)

func ShowIf(x int, y func() int) {
	if x > 10 {
		fmt.Println("x 比 10 大")
	} else if x == 10 {
		fmt.Println("x 等於 10")
	} else {
		fmt.Println("x 比 10 小")
	}

	if y := y(); y > 0 { // 簡短變數宣告在判斷內
		fmt.Println("正數")
	}
}

// ShowSwitch go 的 switch 會編譯成 binary search, 所以當分支很多時候會比 if 效率高
// 雖然目前資料看下來雖然沒有 jvm 的 switch 優化效率高, 但仍然有相同的使用情境
func ShowSwitch(day string) {
	switch day {
	case "Monday":
		fmt.Println("星期一")
	case "Friday":
		fmt.Println("星期五, 週末快樂")
	default:
		fmt.Println("其他日子:" + day)
	}

	score := 85
	switch { // 也可以不帶參數當作一般 if 使用, 只是某些情況下看起來會更整潔
	case score >= 90:
		fmt.Println("優等")
	case score >= 60:
		fmt.Println("及格")
		fallthrough // 無視下一個條件直接執行 XXX 這功能感覺很垃圾耶?
	case score <= 20:
		fmt.Println("當掉")
	default:
		fmt.Println("不及格")
	}
}

func ShowGoto() int {
	// goto 不能跨 function 跳, 不能跳進 block, 只能跳到同一層裡面定義的 label, 所以不大會造成以前常見的跳轉地獄
	// 之所以保留 goto 是在某些性能極端敏感的程式碼區塊能避開函式呼叫開銷, 或幫助生成更佳的機器碼
	// Go 的設計者(像是 Ken Thompson)是寫作業系統的大神, 對這種精確控制流程的能力非常執著
	// 而且有時候善用 goto 可以避免一堆巢狀 if-else 的情況達成有條理地跳轉

	file, err := os.Open("file.txt")
	if err != nil {
		goto ERROR
	}

	// do something with file

	file.Close()
	return 0

ERROR: // label 有區分大小寫
	fmt.Println("Something went wrong:", err)
	return 1
}

// ShowFor for 是唯一迴圈關鍵字
func ShowFor() {
	show := func(val any) {
		fmt.Print(val, ",")
	}

	for a := 10; a < 20; a++ {
		show(a)
	}
	fmt.Println()

	// 類似 while
	b := 0
	for b < 10 {
		show(b)
		b++
	}
	fmt.Println()

	// 無限迴圈
	//for {
	//}

	// 遍歷陣列
	nums := []int{9, 8, 7}
	for index, num := range nums { // 不需要 index 可以使用 _ 或乾脆不寫都可以
		fmt.Print(index, ":", num, ",")
	}
	fmt.Println()

	// 遍歷 map
	dict := map[string]int{"a": 1, "b": 2}
	for key, val := range dict {
		fmt.Print(key, val, ",")
	}
	fmt.Println()

	// 遍歷字串
	str := "Hello 世界"
	for index, ch := range str {
		fmt.Printf("[%d] \"%c\"\n", index, ch)
	}
	fmt.Println()
}
