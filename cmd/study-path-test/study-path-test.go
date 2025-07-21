package main

import (
	"aery-study-go/internal/user"
	"aery-study-go/pkg/errortest"
	"fmt"
	"github.com/cockroachdb/errors"
)

func main() {
	fmt.Println("=== Testing different path locations ===")

	fmt.Println("\n--- 1. Error from pkg directory ---")
	pkgErr := errortest.PkgError()
	fmt.Printf("Error: %v\n", pkgErr)
	fmt.Printf("Detailed:\n%+v\n", pkgErr)

	fmt.Println("\n--- 2. Wrapped error from pkg directory ---")
	wrappedPkgErr := errortest.PkgWrapError(errors.New("base error"))
	fmt.Printf("Error: %v\n", wrappedPkgErr)
	fmt.Printf("Detailed:\n%+v\n", wrappedPkgErr)

	fmt.Println("\n--- 3. Error from internal directory ---")
	internalErr := user.InternalError()
	fmt.Printf("Error: %v\n", internalErr)
	fmt.Printf("Detailed:\n%+v\n", internalErr)

	fmt.Println("\n--- 4. Wrapped error from internal directory ---")
	wrappedInternalErr := user.InternalWrapError(errors.New("base error"))
	fmt.Printf("Error: %v\n", wrappedInternalErr)
	fmt.Printf("Detailed:\n%+v\n", wrappedInternalErr)

	fmt.Println("\n--- 5. Third-party library error (cockroachdb/errors) ---")
	// 這會顯示第三方函式庫內部的 stack trace
	err := errors.Errorf("error created with Errorf: %d", 42)
	fmt.Printf("Error: %v\n", err)
	fmt.Printf("Detailed:\n%+v\n", err)
}
