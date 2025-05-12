package main

import (
	"aery-study-go/pkg/where"
	"fmt"
	"sync/atomic"
)

func main() {
	var counter int64

	where.WrapPrint("Load", func() {
		n := atomic.LoadInt64(&counter)
		fmt.Println(n)
	})

	where.WrapPrint("Add", func() {
		new := atomic.AddInt64(&counter, 1)
		fmt.Println(new)
	})

	where.WrapPrint("Store", func() {
		atomic.StoreInt64(&counter, 100)
		fmt.Println(atomic.LoadInt64(&counter))
	})

	where.WrapPrint("CompareAndSwap", func() {
		swapped := atomic.CompareAndSwapInt64(&counter, 100, 200)
		fmt.Println(swapped, atomic.LoadInt64(&counter))
	})
}
