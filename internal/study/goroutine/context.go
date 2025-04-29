package goroutine

// context 主要用來:
// 攜帶截止時間（Deadline）或超時（Timeout）訊息
// 在多個 goroutine 間傳遞取消（Cancel）訊號
// 傳遞一些跨 API 邊界的 request-scoped 資料（像是用戶 ID、認證 token）

func ShowContext() {
	//var ctx context.Context
	//var cancel context.CancelFunc
	//ctx = context.Background()                          // 最基本的空 context, 通常在 main、init 或測試時使用
	//ctx = context.TODO()                                // "我之後再想要不要用 context"用的, 通常是占位用(不建議長期留在正式碼裡)
	//ctx, cancel = context.WithCancel(nil)               // 產生一個"可以被手動取消"的 context
	//ctx, cancel = context.WithTimeout(nil, time.Second) // 產生一個"自動超時"的 context, 超時就會自動取消
	//ctx, cancel = context.WithDeadline(nil, time.Now()) // 產生一個"在指定時間點"自動取消的 context
	//ctx = context.WithValue(nil, "a", 1)                // 在 context 中夾帶資料

}
