package study

import "fmt"

// interface 本身是一個 tuple, 內含兩個值: type (實際型別) 跟 value (實際值),
// 它是一組方法的集合, 任何型別實作該 interface 的所有方法就自動視為該 interface 的實作, 不需要顯式的 "implement",
// 所以 interface{} 也就可以代表任何型別, 也就像是 java 的 Object 代表所有物件的最上層.

// 也由於 go 的方法是逐個寫在 struct 之外, 會導致無法一眼看出該 struct 實作了哪些 interface,
// 而無法對 struct 勾勒出一個領域概念, 但這是從物件導向角度來看 struct,
// 我想 go 的核心理念並不是包裝"物件"而是包裝"行為", 也就是說以"行為"為設計導向,
// 所以從這個角度出發, 就不會覺得 interface 這樣設計很奇怪了,
// 原因是 interface 只是要呈現一組方法, 只要任意型別有該方法就代表它能提供該行為,
// 所以 interface 的方法就應該簡單明確不應該過多, 將行為領域劃分清楚

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

func (d dog) speak() string {
	return "汪汪"
}
func (d dog) playful() string {
	return "搖尾巴"
}

func (c cat) speak() string {
	return "喵喵"
}
func (c cat) playful() string {
	return "打哈欠"
}

// 實作靜態保護, 可以在編譯時檢查 struct 是否有實作 interface
var _ animal = (*dog)(nil)
var _ animal = (*cat)(nil)

func ShowInterface() {
	dog := dog{}
	cat := cat{}
	makeAnimalSpeak(dog)
	makeAnimalSpeak(cat)
}

func makeAnimalSpeak(a animal) {
	fmt.Println(a.speak())

	// 將介面還原為原本的型別
	d, isDog := a.(dog)
	c, isCat := a.(cat)
	fmt.Printf("isDog:%t(%T), isCar:%t(%T)\n", isDog, d, isCat, c)

	switch t := a.(type) {
	case dog:
		fmt.Println("這是狗", t)
	case cat:
		fmt.Println("這是貓", t)
	default:
		fmt.Println("未知生物", t)
	}
}
