package main

import (
	"aery-study-go/pkg/where"
	"fmt"
)

func main() {
	// 用於建立型別 slice/map/channel 的記憶體分配,
	// 它們實際上是一個 struct, 內有一個指標指向"真實資料"位置,
	// 所以它們是種 "reference-like 的 struct",
	// 傳遞時(pass by value) 會把指向"真實資料"的指標傳過去,
	// 因此對這三種型別操作資料時, 實際上是 reference 的效果.
	// 所以需要透過 make() 來初始化這三種型別才能正常

	/*
		type slice struct {
		    array	unsafe.Pointer	// 指向底層 array 的指標
		    len		int				// 長度
		    cap		int				// 容量
		}
	*/
	where.WrapPrint("make(slice)", func() {
		s := make([]int, 3, 5)
		s[0] = 1
		s[1] = 2
		s[2] = 3
		fmt.Println("slice s:", s)
	})

	/*
		type hmap struct {
			count		int 			// 實際元素數量
			flags		uint8
			B 			uint8			// buckets 的大小 log2
			noverflow	uint16
			hash0		uint32			// hash seed
			buckets		unsafe.Pointer	// *bmap
			oldbuckets	unsafe.Pointer
			nevacuate	uintptr
			extra		*mapextra
		}
	*/
	where.WrapPrint("make(map)", func() {
		m := make(map[string]int)
		m["apple"] = 5
		m["banana"] = 10
		fmt.Println("map m:", m)
	})

	/**
	type hchan struct {
		qcount   uint           // channel 中目前元素數量
		dataqsiz uint           // channel buffer 大小
		buf      unsafe.Pointer // 指向環狀 buffer
		elemsize uint16
		closed   uint32
		sendx    uint           // 下一個要送的 index
		recvx    uint           // 下一個要收的 index
		recvq    waitq          // 等待 recv 的 goroutine queue
		sendq    waitq          // 等待 send 的 goroutine queue
		lock     mutex
	}
	*/
	where.WrapPrint("make(chan)", func() {
		ch := make(chan int, 2)
		ch <- 42
		ch <- 100
		fmt.Println("channel ch:", <-ch, <-ch)
	})
}
