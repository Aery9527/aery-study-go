package main

import (
	"aery-study-go/pkg/utils"
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"
)

// context 本質是一個 chain, 可以串起多個 context 來組裝多個功能, 主要用來:
// 攜帶截止時間（Deadline）或超時（Timeout）訊息
// 在多個 goroutine 間傳遞取消(cancel)訊號
// 傳遞一些系統用的資訊, 例如自定義的 log id, goroutines id, trace id 等等

func main() {
	var parentCtx context.Context = context.TODO() // "我之後再想要不要用 context"用的, 通常是占位用(不建議長期留在正式碼裡)
	parentCtx = context.Background()               // 一個空的 context, 基本是用來組裝 context chain 的頭

	utils.WrapPrint("withValue", withValue)
	utils.WrapPrint("withValue", withValue)
	utils.WrapPrint("withTimeout", withTimeout)
	utils.WrapPrint("withoutCancel", withoutCancel)

	utils.WrapPrint("withCancel", func() { withCancel(context.WithCancel(parentCtx)) })
	utils.WrapPrint("WithCancelCause", func() { // 跟 WithCancel 差別在於可以附 cancel cause
		ctx, cancel := context.WithCancelCause(parentCtx)
		withCancel(ctx, func() { cancel(errors.New("kerker")) })
	})

	delay := 100
	duration := func() time.Time { // 當前時間過 delay ms 的一個時刻
		return time.Now().Add(time.Duration(delay) * time.Millisecond)
	}
	utils.WrapPrint("withDeadline", func() {
		ctx, cancel := context.WithDeadline(parentCtx, duration())
		withDeadline(ctx, cancel, delay)
	})
	utils.WrapPrint("withDeadlineCause", func() { // 跟 WithCancel 差別在於可以附 cancel cause
		ctx, cancel := context.WithDeadlineCause(parentCtx, duration(), errors.New("kerker"))
		withDeadline(ctx, cancel, delay)
	})
}

// withValue 攜帶 kv 的 context
func withValue() {
	// 值只能進不能改
	ctx1 := context.WithValue(context.Background(), "name", "Aery")
	ctx2 := context.WithValue(ctx1, "age", 18)
	fmt.Println(ctx2)
	name := ctx2.Value("name").(string)
	age := ctx2.Value("age").(int)
	fmt.Println(name + ":" + strconv.Itoa(age))
}

// withTimeout 自動 timeout context
func withTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("finished work")
	case <-ctx.Done(): // 這個會被執行, 因為 1s 就超時
		fmt.Println("timeout:", ctx.Err())
	}
}

// withoutCancel 與 parent context 斷開 cancel 連結的 context
// 也就是說 parent context cancel 時, 子 context 不會 cancel
func withoutCancel() {
	ctx1, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	ctx2 := context.WithoutCancel(ctx1)

	go func() {
		<-ctx1.Done()
		fmt.Println("ctx1 finish:", ctx1.Err())
	}()

	go func() {
		select {
		case <-ctx2.Done(): // ctx1 1s timeout, 若傳遞給 ctx2 的話, 這邊會被執行
			fmt.Println("ctx2 cancel:", ctx2.Err())
		case <-time.After(200 * time.Millisecond): // 但因為 ctx2 WithoutCancel(ctx1), 所以這邊才會被執行
			fmt.Println("ctx2 timeout")
		}
	}()

	time.Sleep(300 * time.Millisecond) // 這裡等超過 ctx1 的 timeout 時間, 所以 ctx1 觸發 deadline 而不是 defer 的 cancel
}

func withCancel(ctx context.Context, cancel func()) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("ctx.Err(): %s, cause: %s\n", ctx.Err(), context.Cause(ctx))
				return
			default:
				// 這樣的寫法會有個小問題, 上面要執行可能會多等一個小於 500ms 的時間,
				// 例如這邊等了 200ms 之後 ctx cancel 了, 可是這邊還是在 sleep, 所以得繼續等 300ms 才會進到上面的 case
				fmt.Println("wait start...")
				time.Sleep(100 * time.Millisecond)
				fmt.Println("wait finish...")
			}
		}
	}()

	time.Sleep(500 * time.Millisecond)
	fmt.Println("canceling context")
	cancel()
}

func withDeadline(ctx context.Context, cancel func(), delay int) {
	defer cancel()

	select {
	case <-ctx.Done(): // 因為 ctx 的 delay, 而下面是 delay * 2, 所以這邊會先執行
		fmt.Printf("ctx.Err(): %s, cause: %s\n", ctx.Err(), context.Cause(ctx))
	case <-time.After(time.Duration(delay*2) * time.Millisecond):
		fmt.Println("timeout")
	}
}
