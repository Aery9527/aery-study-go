package main

import "fmt"

func main() {
	// chan <TYPE>, TYPE 是這個 channel 傳遞資料的型別
	// 無緩衝概念就是像是 A 有東西要給 B, 但 B 一直沒來拿, 所以 A 就只能一直拿在手上等然後無法去做其他事情, 這就是阻塞(block)
	// 有緩衝概念就是有個籃子, A 有東西要給 B 就先放到這個籃子裡, A 就可以去做其他事情了, B 也可以隨時來籃子拿東西,
	// 相對的當籃子滿了的話 A 就必須能在手上, 此時就又是 阻塞(block) 狀態了
	// 當以 B 的角度也是一樣的概念
	latch := make(chan any) // 無緩衝 channel, 傳送接收資料都是 block 的, 這樣的機制使得多個 goroutine 之間同步非常簡單
	ch := make(chan int, 2) // 有緩衝 channel, 可以讓 goroutine 之間有限量的不阻塞溝通

	go goChannel(ch, latch)

	fmt.Printf("ch: %v\n", <-ch)
	fmt.Printf("ch: %v\n", <-ch)
	//fmt.Printf("ch: %v\n", <-ch) // 這邊再取就會 deadlock, 因為只有 2 個buffer, 上面取完後由於另外一個 goroutine 要等下面喚醒, 所以兩邊就互卡了
	latch <- true // 如果沒有人收走, 這邊就會 block

	close(latch)
	close(ch)
	//close(ch) // 再次關閉會 panic

	//ch <- 9527 // 對 close 的 channel 寫入資料會 panic
	fmt.Printf("ch: %v\n", <-ch) // 關閉之後就會回傳型態的零值(zero value)

	// 可以使用 range 不斷取出 channel 裡的資料, 直到 channel 被關閉
	for v := range ch {
		fmt.Println(v)
	}
}

func goChannel(ch chan int, latch chan any) {
	ch <- 1
	ch <- 2
	_ = <-latch // 若沒有值可以取則會block
}
