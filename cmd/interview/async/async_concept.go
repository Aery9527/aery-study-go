package main

import (
	"fmt"
	"sync"
)

func main() {
	target := new(int) // 有更好的

	waitGroup := sync.WaitGroup{} // WaitGroup 功能為何?
	waitGroup.Add(2)

	mutex := sync.Mutex{} // XXX 這個與 sync.RWMutex 有何差別?
	go increment(&waitGroup, &mutex, target)
	go increment(&waitGroup, &mutex, target)

	waitGroup.Wait() // XXX 如果沒有這行會發生甚麼事?

	fmt.Printf("*target: %d\n", *target)
}

// increment XXX 要達成這個功能有更好的方法嗎?
func increment(waitGroup *sync.WaitGroup, mutex *sync.Mutex, target *int) int {
	defer waitGroup.Done()

	mutex.Lock()
	defer mutex.Unlock()

	*target++
	return *target
}
