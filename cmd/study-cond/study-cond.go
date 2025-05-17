package main

import (
	"fmt"
	"sync"
	"time"
)

// 使用在生產者/消費者模式中, 生產者可用來通知消費者有新數據可用, 否則消費者會 block
// cond.Wait() 消費者用來等待資料
// cond.Signal() 生產者用來通知"1個"消費者有新資料可用
// cond.Broadcast() 生產者用來通知"所有"消費者有新資料可用

// 注意: 這些都操作當是"當下", 不朔及既往,
// 若生產者先執行了 cond.Signal(),
// 消費者才執行 cond.Wait() 就有可能會導致消費者永遠 block,
// 所以相關周邊配套要處理好, 如這邊先判斷 queue 為空才會進入 cond.Wait()

type Queue[T any] struct {
	items  []T
	lock   sync.Mutex
	cond   *sync.Cond
	cancel bool
}

func NewQueue[T any]() *Queue[T] {
	q := &Queue[T]{}
	q.cond = sync.NewCond(&q.lock)
	return q
}

func (q *Queue[T]) Enqueue(item T) bool {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.cancel {
		return false
	}

	q.items = append(q.items, item)
	q.cond.Signal() // 喚醒一個消費者
	return true
}

func (q *Queue[T]) Dequeue() (T, bool) {
	q.lock.Lock()
	defer q.lock.Unlock()

	for len(q.items) == 0 {
		fmt.Println("佇列為空，等待中...")
		q.cond.Wait() // 會自動 unlock mutex, 喚醒後會重新 lock

		if q.cancel { // 如果取消了，則退出
			return *new(T), false
		}
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

func (q *Queue[T]) Cancel() []T {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.cancel = true
	q.cond.Broadcast()

	return q.items
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	queue := NewQueue[int]()

	// 生產者 goroutine
	go func() {
		defer wg.Done()
		for i := 1; ; i++ {
			ok := queue.Enqueue(i)
			if ok {
				fmt.Println("生產：", i)
			} else {
				fmt.Println("生產者退出")
				break
			}
			time.Sleep(400 * time.Millisecond) // 模擬生產延遲
		}
	}()

	// 消費者 goroutine
	go func() {
		defer wg.Done()
		for {
			element, ok := queue.Dequeue()
			if ok {
				fmt.Println("消費：", element)
			} else {
				fmt.Println("消費者退出")
				break
			}
		}
	}()

	time.Sleep(2 * time.Second)
	queue.Cancel()
	wg.Wait()
}
