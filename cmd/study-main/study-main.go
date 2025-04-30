// 一定要是 main package, 才能執行 main()
package main

import (
	kerker "aery-study-go/pkg/utils" // 給一個別名, 預設使用最後一層 folder 當作 namespace, 但若有衝突時就可以使用別名化解
	"flag"
	"fmt"
	"os"
	"strings"
	// . "fmt" 就像 java 的 import static 一樣, 這樣就可以直接用 Println() 而不用帶 package 了
	// _ "fmt" (side effect import), import 一個 package 但沒打算用裡面的任何 symbol(function, struct, constant 等), 只想讓它的 init() 執行就可以這樣寫
)

// GO 有兩個保留函數 init() 跟 main(), 它們不能帶參數或回傳值

// init() 不能帶參數或回傳值, 會在 main() 被呼叫之前執行
// 可以有多個 init(), 甚至在同一個檔案中也可以定義多次
// 每個檔案的 init() 執行順序依照 Go compiler 決定的 import 順序來定
// 主要用途就是為了模組初始化, 不需要特別去呼叫它, Go compiler 會搞定
func init() {
	kerker.WrapPrint("func init()", func() {
		fmt.Println("~~~yo~~~yo~~~yo~~~")
	})
}

// 程式進入點
func main() {
	//os.Args[0] = "Aery Handsome~~~" // XXX 這居然是可變的!?
	cmdArgs := os.Args // 取得 cmd 傳進來的參數, [0] 永遠是當前執行檔案完整路徑與名稱, [1] 開始才是傳進來的參數 (這樣挺不錯的, 夠直覺)
	lineSeparator := kerker.GetSystemLineSeparator()
	kerker.WrapPrint("參數列表", func() {
		fmt.Printf("%s\n", strings.Join(cmdArgs, lineSeparator))
	})

	// 使用 flag 套件來解析命令列參數, cmd 格式為 -key=value, 遇到非這個格式就會停下不解析往後的參數了
	// flag.Parse() // error at here, 因為沒有先把所有key給接起來
	name := flag.String("name", "黑山老妖", "名字")
	age := flag.Int("age", 9527, "年齡")
	//height := flag.Int("age", 0, "身高") // 所有出現的參數都要接下否則會報錯, 因為 flag 屬於"強驗證型"套件
	flag.Parse() // 執行解析, 將值塞到上述的指標裡
	kerker.WrapPrint("flag.Parse()", func() {
		fmt.Printf("嗨 %d 歲的 %s ❤️\n", *age, *name)
	})

	// 取得所有解析的 cmd args, 但一定要先執行過 flag.Parse(), 否則啥屁也拿不到
	kerker.WrapPrint("flag.VisitAll()", func() {
		flag.VisitAll(func(f *flag.Flag) {
			fmt.Printf("flag(-%s) value(%s) usage(%s)\n", f.Name, f.Value, f.Usage)
		})
	})
}
