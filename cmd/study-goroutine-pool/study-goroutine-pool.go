package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// 實現一個簡單的 goroutine pool 當作練習

func main() {
	poolSize := 3   // 幾個 goroutine 在 pool 中執行
	queueNum := 100 // 任務佇列大小

	jobQueue := make(chan func() int, queueNum)
	resultQueue := make(chan string, queueNum)
	latch := make(chan struct{}) // channel 不需要傳值時, 使用 struct{} 是零大小(zero-sized)不佔記憶體, 最佳同步手段

	// 啟動 poolSize 數量的 goroutine
	for workerId := 1; workerId <= poolSize; workerId++ {
		go worker(workerId, jobQueue, resultQueue, latch)
	}

	// 當前 goroutine 投放任務去執行
	for jobId := 1; jobId <= 11; jobId++ {
		jobQueue <- func() int {
			rest := 100 + rand.Intn(201)                       // 沒有設定隨機種子, 可能導致每次執行會產生相同的隨機數序列
			time.Sleep(time.Duration(rest) * time.Millisecond) // 模擬任務耗時
			return jobId
		}
	}
	close(jobQueue)

	// 用以監控 worker 是否全部結束, 全部結束關閉 resultQueue 讓主 goroutine 結束
	go func() {
		finishWorkerCount := 0
		for range latch {
			finishWorkerCount++
			if finishWorkerCount == poolSize {
				// 因為 job 與 result queue 是 1:1, 所以當 worker 全部結束時就代表 jobQueue 也已全部執行完, resultQueue 也有相對應的數量
				close(resultQueue)
				break
			}
		}
	}()

	// 如果 buffer channel 內還有東西但已經關閉, range 會全部取完再退出
	for result := range resultQueue {
		fmt.Printf("result: %s\n", result)
	}

	fmt.Println("finish")
}

func worker(id int, jobQueue chan func() int, resultQueue chan string, latch chan struct{}) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("worker(%d) panic: %v\n", id, r)
		}
		latch <- struct{}{} // 確保即使 panic 也會通知主 goroutine
	}()

	for job := range jobQueue {
		jobId := job() // 執行任務
		fmt.Printf("worker(%d) 處理 job: %d\n", id, jobId)
		resultQueue <- strconv.Itoa(jobId)
	}

	latch <- struct{}{} // 通知主 goroutine 自己已經結束了
}
