package study

import (
	"aery-study-go/util"
	"flag"
	"fmt"
	"os"
	"strings"
)

func ShowArgs() {
	//os.Args[0] = "Aery Handsome~~~" // XXX 居然是可變的!?
	cmdArgs := os.Args // 取得 cmd 傳進來的參數, [0] 永遠是當前執行檔案完整路徑與名稱, [1] 開始才是傳進來的參數 (這樣挺不錯的, 夠直覺)
	lineSeparator := util.GetSystemLineSeparator()
	fmt.Printf("參數列表：%s%s\n", lineSeparator, strings.Join(cmdArgs, lineSeparator))
	fmt.Println()

	// 使用 flag 套件來解析命令列參數, cmd 格式為 -key=value, 遇到非這個格式就會停下不解析往後的參數了
	//flag.Parse() // error at here, 因為沒有先把所有key給接起來
	name := flag.String("name", "黑山老妖", "名字")
	age := flag.Int("age", 9527, "年齡")
	//height := flag.Int("age", 0, "身高") // 所有出現的參數都要接下否則會報錯, 因為 flag 屬於"強驗證型"套件
	flag.Parse() // 執行解析, 將值塞到上述的指標裡
	fmt.Printf("嗨 %d 歲的 %s ❤️\n", *age, *name)
	fmt.Println()

	// 取得所有解析的 cmd args, 但一定要先執行過 flag.Parse(), 否則啥屁也拿不到
	flag.VisitAll(func(f *flag.Flag) { // 沒有像 java -> 可以縮成一行的語法糖
		fmt.Printf("flag(-%s) value(%s) usage(%s)\n", f.Name, f.Value, f.Usage)
	})
}
