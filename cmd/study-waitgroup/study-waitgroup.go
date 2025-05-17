package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	start := time.Now() // 開始計時

	sleep1 := int(rand.Float32() * 1000)
	sleep2 := int(rand.Float32() * 1000)

	go func() {
		defer wg.Done()
		time.Sleep(time.Duration(sleep1) * time.Millisecond)
	}()

	go func() {
		defer wg.Done()
		time.Sleep(time.Duration(sleep2) * time.Millisecond)
	}()

	wg.Wait() // 等待全部完成

	elapsed := time.Since(start) // 計算經過的時間

	fmt.Println("sleep1: ", sleep1, "ms")
	fmt.Println("sleep2: ", sleep2, "ms")
	fmt.Println("elapsed:", elapsed.Milliseconds(), "ms")
}
