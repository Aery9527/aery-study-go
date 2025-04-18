package study

import (
	_ "aery-study-go/util" // 這行的底線代表只要執行這個檔案就會執行 util/init_study.go 的 init() 函式
	"fmt"
)

func Use_side_effect_import() {
	fmt.Println("side_effect_import.go")
}
