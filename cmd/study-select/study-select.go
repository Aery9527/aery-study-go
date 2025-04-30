package main

import (
	"fmt"
	"time"
)

// select 可以監聽多個 channel 的狀態, 只要有一個 channel 有資料就會執行對應的 case,
// 如果沒有任何 channel 有資料則會 block 等待或執行 default case

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch1)
		}
		ch2 <- 0
	}()

	time.Sleep(100 * time.Millisecond)
	fibonacci(ch1, ch2)
}

func fibonacci(ch1, ch2 chan int) {
	x, y := 1, 1
	for {
		select { // 在多個 channel 中選擇一個可用的 channel
		case ch1 <- x:
			x, y = y, x+y
		case v := <-ch2:
			fmt.Printf("ch2: %d\n", v)
			return
		case <-time.After(time.Second): // 超時機制, 如果 1 秒內沒有任何 channel 有資料則這個 ch 會被返回, 因此這個 case 會被執行
			fmt.Printf("timeout\n")
			return
		default: // 若 ch1 ch2 都 block 則會執行 default
			fmt.Printf("all channels are blocked...\n")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
