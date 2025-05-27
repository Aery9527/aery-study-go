package main

import "fmt"

func main() {
	d := dog{3}
	c := cat{3}

	// XXX 為什麼 d 跟 c 都可以傳入 playWith(pet) ?
	playWith(d)
	playWith(&c)

	// XXX 輸出為何?
	fmt.Printf("Dog weight: %d\n", d.weight)
	fmt.Printf("Cat weight: %d\n", c.weight)
}

func playWith(p pet) {
	p.speak()
	p.feed(2)
}

// XXX 對於 interface 有何想法? 它可以幫你達成甚麼樣的目的?
type pet interface {
	speak() string
	feed(foodWeight int)
}

// XXX dog 跟 cat 都有 weight, 該怎麼調整消除這種重複

type dog struct {
	weight int
}

func (d dog) speak() string {
	return "汪汪"
}

func (d dog) feed(foodWeight int) {
	d.weight += foodWeight
}

type cat struct {
	weight int
}

func (c *cat) speak() string {
	return "喵喵"
}

func (c *cat) feed(foodWeight int) {
	c.weight += foodWeight
}
