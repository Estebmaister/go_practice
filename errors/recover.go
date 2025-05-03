package errors

import (
	"errors"
	"fmt"
	"runtime"
)

func init() {
	fmt.Println(safediv(4, 0))
}

func safediv(a, b int) int {
	defer func() {
		if err := recover(); err != nil {
			if err, ok := err.(error); ok {
				if errors.As(err, new(runtime.Error)) {
					fmt.Println("This is a runtime error:", err)
				} else {
					fmt.Println("This is a regular error:", err)
				}
			}
		}
	}()

	return a / b
}
