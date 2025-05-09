package main

import "fmt"

// interface 本身是一個 tuple, 內含兩個值: type (實際型別) 跟 value (實際值),
// 它用來表示一組方法的集合, 任何型別實作該 interface 的所有方法就自動視為該 interface 的實作, 不需要顯式的 "implement",
// 所以 interface{} 也就可以代表任何型別, 也就像是 java 的 Object 代表所有物件的最上層.

// 也由於 go 的方法是逐個寫在 struct 之外, 會導致無法一眼看出該 struct 實作了哪些 interface,
// 而無法對 struct 勾勒出一個領域概念, 但這是從物件導向角度來看 struct,
// 我想 go 的核心理念並不是包裝"物件"而是包裝"行為", 也就是說以"行為"為設計導向,
// 所以從這個角度出發, 就不會覺得 interface 這樣設計很奇怪了,
// 原因是 interface 是要呈現一組領域方法, 只要任意型別有該 interface 的所有方法就代表它能提供該領域行為,
// 所以 interface 的方法就應該簡單明確不應該過多, 將行為領域劃分清楚

// 寫一寫體悟了 go 這樣設計 interface 與其 implement 的偉大之處,
// 是可以讓 "實作" 與 "介面" 也直接解耦合!! (超神~~~)
// 假設今天我要寫一個支援 10 個 log 介面的實作 lib, 暫時叫 kerker 好了,
// 在 java 的世界是這個 kerker 勢必要一起打包這 10 個介面的 lib,
// 然後使用這個 kerker 的 java 也間接地一定要包含這 10 個介面的 lib,
// 儘管他可能只用到了其中 1 個介面而已, 但不得已一定要把這 10 個介面的 lib 都帶進來,
// 否則 java 會編譯錯誤.
// 但在 go 的世界就沒這困擾了, 一個 struct 實作了那些介面,
// 取決於這個 struct 身上的方法與當前 runtime 所有介面的方法是不是一模一樣,
// 不需要"顯式"的包含"介面"資訊, 也就達成了跟"介面"解耦合的效果,
// 但當然該介面上的方法其參數與回傳都是原生型別, 那就會是真正的完全解耦合.
// 所以從這角度出發, 在設計介面的時候就要想想是不是這介面可以簡單到讓別人輕鬆替換掉實作,
// 如果可以, 那麼介面上的參數與回傳都應該是原生型別, 降低別人與你的耦合.

type dog struct{}

type cat struct{}

type animal interface {
	speak() string
}

// 不同介面有相同方法則會視為互相實作,
// 但實務上不應該這樣搞, 應該將 playful() 行為抽出為另一個 interface,
// tsundere 跟 passionate 再去包它, 所以要注意這個 package 底下所有介面的方法命名!

type tsundere interface { // 傲嬌
	animal
	playful() string
}
type passionate interface { // 熱情
	animal
	playful() string
}

// dog func
func (d dog) speak() string {
	return "汪汪"
}
func (d dog) playful() string {
	return "搖尾巴"
}

// cat func
func (c cat) speak() string {
	return "喵喵"
}
func (c cat) playful() string {
	return "打哈欠"
}

// func 綁定只能對 local type (同個 package) 使用
// func (i int) speck() string {} // 像這樣不行
// 所以如果要對 int 綁方法則可以用別名

type defiantInt int

func (i defiantInt) Hello() string { // 像這樣就可以對 int 綁方法了
	return "my value is " + fmt.Sprint(i)
}

// 實作靜態保護, 可以在編譯時檢查 struct 是否有實作 interface
var _ animal = (*dog)(nil)
var _ animal = (*cat)(nil)

func main() {
	dog := dog{}
	cat := cat{}
	makeAnimalSpeak(dog)
	makeAnimalSpeak(cat)
}

func makeAnimalSpeak(a animal) {
	fmt.Println(a.speak())

	// 將介面還原為原本的型別
	d, isDog := a.(dog) // 轉型, 要轉換的型別必須是子類
	c, isCat := a.(cat)
	fmt.Printf("isDog:%t(%T), isCar:%t(%T)\n", isDog, d, isCat, c)

	switch t := a.(type) { // 這個寫法只能在 switch 裡面使用
	case dog:
		fmt.Println("這是狗", t)
	case cat:
		fmt.Println("這是貓", t)
	default:
		fmt.Println("未知生物", t)
	}
}
