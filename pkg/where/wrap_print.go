package where

import (
	"fmt"
	"sync"
)

var setSyncPrintLock = sync.Mutex{}
var printer = wrapPrint

func SetSyncPrint(async bool) {
	setSyncPrintLock.Lock()
	defer setSyncPrintLock.Unlock()

	if async {
		printer = wrapPrint
	} else {
		mutex := sync.Mutex{}
		printer = func(scope string, action func(), at int) {
			mutex.Lock()
			defer mutex.Unlock()
			wrapPrint(scope, action, at+1)
		}
	}
}

func WrapPrint(scope string, action func()) {
	printer(scope, action, 1)
}

func wrapPrint(scope string, action func(), at int) {
	fmt.Println()
	fmt.Printf("[%s] ", scope)
	PrintWhereAt(at + 1)
	action()
}
