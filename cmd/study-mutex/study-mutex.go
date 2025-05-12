package main

import (
	"aery-study-go/pkg/where"
	"fmt"
	"sync"
)

// 互斥鎖: 僅有一個 goroutine 可以取得鎖, 其他 goroutine 會在邊界 block 等待獲得鎖

// 在鎖的範圍內對所有變數的操作都有立即可見性,
// 不會有 cpu cache 的問題, 所以在退出鎖後其他 goroutine 可以立即取得新值

func init() {
	where.SetSyncPrint(false)
}

func main() {
	// latch 用來控制 goroutine 在正確的狀態等待或執行, 以確保測試正確
	latch1 := make(chan struct{})
	latch2 := make(chan struct{})

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)

	mutex := sync.Mutex{} // XXX test target
	go func() {
		name := "g1"

		defer waitGroup.Done()

		mutex.Lock()
		defer mutex.Unlock()

		where.WrapPrint(name, func() { fmt.Println("lock") })
		latch1 <- struct{}{}
		<-latch2
		where.WrapPrint(name, func() { fmt.Println("unlock") })
	}()
	go func() {
		name := "g2"

		defer waitGroup.Done()

		<-latch1

		gotLock := mutex.TryLock()
		if gotLock {
			defer mutex.Unlock()
			panic("should not get lock")
		} else {
			where.WrapPrint(name, func() { fmt.Println("try lock failed") })
			latch2 <- struct{}{}
		}
	}()

	waitGroup.Wait()
}
