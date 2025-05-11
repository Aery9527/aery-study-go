package main

import (
	"aery-study-go/pkg/utils"
	"fmt"
	"sync/atomic"
)

func main() {
	var counter int64

	utils.WrapPrint("Add", func() {
		atomic.AddInt64(&counter, 1)
		fmt.Println(atomic.LoadInt64(&counter))
	})

	utils.WrapPrint("Store", func() {
		atomic.StoreInt64(&counter, 100)
		fmt.Println(atomic.LoadInt64(&counter))
	})

	utils.WrapPrint("CompareAndSwap", func() {
		swapped := atomic.CompareAndSwapInt64(&counter, 100, 200)
		fmt.Println(atomic.LoadInt64(&counter))
		fmt.Println(swapped)
	})
}
