package main

import (
	"fmt"
	"sync"
)

// Once 是一個結構體, 用來確保某個函數只會執行一次
// 具有原子性操作, 確保在多執行緒環境下, 只會執行一次
// 如果第一次執行 panic 且在外部被 recover, 仍然是算"第一次"被執行

func main() {
	times := 0
	action := func(id string) {
		times++
		//fmt.Println(id, "執行中", times) // 這段打開可以觀察到後續不會再執行第二次
		//if times == 1 {
		//	panic("爆炸啦")
		//}

		fmt.Println(id, "應該只會出現一次")
	}

	var once sync.Once

	var wg sync.WaitGroup
	wg.Add(3)

	test := func(id string) {
		defer wg.Done() // 執行完就 -1
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(id, "recover error:", r)
			}
		}()

		fmt.Println(id, "開始執行")
		once.Do(func() { action(id) })
	}

	go test("A")
	go test("B")
	go test("C")

	wg.Wait() // 等全部 goroutine 結束
	fmt.Println("finish")
}
