package main

import (
	"aery-study-go/pkg/where"
	"fmt"
	"sync"
	"time"
)

// 讀寫鎖: 用於讀多寫少的情境, 允許同時多個 goroutine 讀取,
// 但僅一個 goroutine 可以寫入, 此時就如同互斥鎖一樣
// 不可在 RLock 內寫資料, 會有 data race 的問題

// 在鎖的範圍內對所有變數的操作都有立即可見性,
// 不會有 cpu cache 的問題, 所以在退出鎖後其他 goroutine 可以立即取得新值

func init() {
	where.SetSyncPrint(false)
}

func main() {
	// latch 用來控制 goroutine 在正確的狀態等待或執行, 以確保測試正確
	latch1 := make(chan struct{})
	latch2 := make(chan struct{})
	latch3 := make(chan struct{})
	latch4 := make(chan struct{})

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(4)
	routine := func(name string, action func(name string)) {
		defer waitGroup.Done()
		action(name)
	}

	// 驗證 rlock 可以同時進入, wlock 會 block 其他 goroutine
	// r1 ┐
	// r2 ┴─> w1 ─> r3
	rwMutex := sync.RWMutex{} // XXX test target
	go routine("r1", func(name string) {
		rwMutex.RLocker() // 取得讀鎖
		defer rwMutex.RUnlock()

		where.WrapPrint(name, func() { fmt.Println("lock") })
		latch1 <- struct{}{}
		<-latch2
		where.WrapPrint(name, func() { fmt.Println("unlock") })
	})
	go routine("r2", func(name string) {
		rwMutex.RLocker() // 取得讀鎖
		defer rwMutex.RUnlock()

		where.WrapPrint(name, func() { fmt.Println("lock") })
		latch1 <- struct{}{}
		<-latch2
		where.WrapPrint(name, func() { fmt.Println("unlock") })
	})
	go routine("w1", func(name string) {
		<-latch1
		<-latch1

		go func() { // 先讓下面進入"寫鎖邊界", 再來讓上面兩個 r1,r2 釋放鎖
			<-time.After(100 * time.Millisecond)
			latch2 <- struct{}{}
			latch2 <- struct{}{}
		}()

		rwMutex.Lock() // 取得寫鎖(互斥鎖的意思)
		defer rwMutex.Unlock()

		where.WrapPrint(name, func() { fmt.Println("lock") })
		latch3 <- struct{}{}
		<-latch4
		where.WrapPrint(name, func() { fmt.Println("unlock") })
	})
	go routine("r3", func(name string) {
		<-latch3

		go func() { // 先讓下面進入"讀鎖邊界", 再來讓上面的 w1 釋放鎖
			<-time.After(100 * time.Millisecond)
			latch4 <- struct{}{}
		}()

		rwMutex.RLocker() // 取得讀鎖
		defer rwMutex.RUnlock()

		where.WrapPrint(name, func() { fmt.Println("lock") })
		where.WrapPrint(name, func() { fmt.Println("unlock") })
	})

	waitGroup.Wait()
}
