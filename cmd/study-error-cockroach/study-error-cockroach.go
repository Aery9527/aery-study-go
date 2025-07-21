package main

import (
	"fmt"
	"github.com/cockroachdb/errors"
	"io"
	"os"
)

// 基礎錯誤，可用於 Is/As 判斷
var errBase = errors.New("base error")

// ErrSentinel Sentinel error
var ErrSentinel = errors.New("sentinel error")

// MyError 自定義錯誤類型
type MyError struct {
	msg string
}

func (e *MyError) Error() string {
	return e.msg
}

func main() {
	fmt.Println("--- 1. New error ---")
	err1 := errors.New("this is a new error")
	fmt.Printf("Error: %v\n", err1)
	fmt.Printf("Detailed error:\n%+v\n", err1)

	fmt.Println("\n--- 2. Wrap error ---")
	err2 := errors.Wrap(err1, "wrapped error")
	fmt.Printf("Error: %v\n", err2)
	fmt.Printf("Detailed error:\n%+v\n", err2)

	fmt.Println("\n--- 3. Wrap with message ---")
	err3 := errors.Wrapf(err2, "wrapped with message: %s", "additional info")
	fmt.Printf("Error: %v\n", err3)
	fmt.Printf("Detailed error:\n%+v\n", err3)

	fmt.Println("\n--- 4. Get cause ---")
	cause := errors.Cause(err3)
	fmt.Printf("Cause of err3: %v\n", cause)

	fmt.Println("\n--- 5. Check error type (Is/As) ---")
	if errors.Is(err3, err1) {
		fmt.Println("err3 is err1")
	}
	if errors.Is(err3, errBase) {
		// This will not be printed as errBase is not in the chain
		fmt.Println("err3 is errBase")
	}

	// 使用一個新的包含 errBase 的錯誤鏈
	errWithBase := errors.Wrap(errBase, "wrapped base error")
	if errors.Is(errWithBase, errBase) {
		fmt.Println("errWithBase is errBase")
	}

	var myErr *MyError
	errWithMyError := errors.Wrap(&MyError{msg: "this is my error"}, "wrapped my error")
	if errors.As(errWithMyError, &myErr) {
		fmt.Printf("errWithMyError is of type MyError with message: %s\n", myErr.msg)
	}

	fmt.Println("\n--- 6. WithMessage ---")
	errWithMessage := errors.WithMessage(err1, "added message")
	fmt.Printf("Error: %v\n", errWithMessage)
	fmt.Printf("Detailed error:\n%+v\n", errWithMessage)

	fmt.Println("\n--- 7. WithHint ---")
	errWithHint := errors.WithHint(err1, "try to fix it by...")
	fmt.Printf("Error: %v\n", errWithHint)
	fmt.Printf("Detailed error:\n%+v\n", errWithHint)

	fmt.Println("\n--- 8. WithDetail ---")
	errWithDetail := errors.WithDetail(err1, "here is a detailed explanation.")
	fmt.Printf("Error: %v\n", errWithDetail)
	fmt.Printf("Detailed error:\n%+v\n", errWithDetail)

	fmt.Println("\n--- 9. Sentinel errors ---")
	errSentinel := errors.Wrap(ErrSentinel, "wrapped sentinel")
	if errors.Is(errSentinel, ErrSentinel) {
		fmt.Println("errSentinel is a wrapped ErrSentinel")
	}

	fmt.Println("\n--- 10. Barrier errors ---")
	// Barrier error 會阻止 Is/As 檢查其內部的錯誤
	barrierErr := errors.New("barrier error")
	errWithBarrier := errors.Wrap(barrierErr, "wrapped")
	errWithBarrier = errors.Mark(errWithBarrier, barrierErr) // 使用 Mark 替代 WithBarrier
	errWithBarrier = errors.Wrap(errWithBarrier, "outer wrap")

	if errors.Is(errWithBarrier, barrierErr) {
		fmt.Println("errWithBarrier is barrierErr (this should not be printed)")
	} else {
		fmt.Println("errWithBarrier is NOT barrierErr (due to barrier)")
	}

	fmt.Println("\n--- 11. Leaf errors ---")
	// Leaf error 會將自己作為錯誤鏈的葉子節點，即使它包裝了其他錯誤
	leafErr := errors.Wrap(io.EOF, "some details")
	leafErr = errors.Mark(leafErr, io.EOF) // 使用 Mark 替代 WithLeaf
	leafErr = errors.Wrap(leafErr, "outer wrap")

	if errors.Is(leafErr, io.EOF) {
		fmt.Println("leafErr is io.EOF (this should not be printed)")
	} else {
		fmt.Println("leafErr is NOT io.EOF (due to leaf)")
	}
	// Cause 會返回 Leaf error 本身
	fmt.Printf("Cause of leafErr: %v\n", errors.Cause(leafErr))

	fmt.Println("\n--- 12. Error properties (AssertionFailedf, Newf, etc.) ---")
	// 當斷言失敗時使用
	val := 10
	if val != 5 {
		errAssertion := errors.AssertionFailedf("value should be 5, but got %d", val)
		fmt.Printf("Detailed error:\n%+v\n", errAssertion)
	}

	// Newf
	errNewf := errors.Newf("new error with code %d", 123)
	fmt.Printf("Detailed error:\n%+v\n", errNewf)

	fmt.Println("\n--- 13. Combining multiple errors ---")
	errs := []error{
		errors.New("first error"),
		errors.New("second error"),
		nil, // nil errors should be skipped
		errors.New("third error"),
	}
	combinedErr := errors.Join(errs...)
	if combinedErr != nil {
		fmt.Printf("Combined error:\n%+v\n", combinedErr)
	}

	fmt.Println("\n--- 14. File and line info ---")
	errFromFile := readFile("non_existent_file.txt")
	fmt.Printf("Error with file info:\n%+v\n", errFromFile)
}

func readFile(filename string) error {
	_, err := os.Open(filename)
	if err != nil {
		// Wrap the error from os.Open, preserving the original error and adding context.
		// The stack trace will point here.
		return errors.Wrapf(err, "failed to open file %s", filename)
	}
	return nil
}
