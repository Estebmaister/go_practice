package errors

import (
	"errors"
	"fmt"

	pkgErr "github.com/pkg/errors"
)

// ReturningError returns a specific error based on the number provided.
// The possible values are 1, 2 and 3. Anything else returns nil.
func ReturningError(errNumber int) error {
	err1 := fmt.Errorf("first error occurred: fmt")
	err2 := errors.New("second  error occurred: errors")
	err3 := pkgErr.New("third error occurred: pkg/errors")
	switch errNumber {
	case 1:
		return err1
	case 2:
		return err2
	case 3:
		return err3
	default:
		return nil
	}
}

func Testing() {
	fmt.Println(ReturningError(1))         // first error occurred
	fmt.Printf("%+v\n", ReturningError(1)) // first error occurred
	fmt.Println(ReturningError(2))         // second  error occurred
	fmt.Printf("%+v\n", ReturningError(2)) // second  error occurred
	fmt.Println(ReturningError(3))         // third error occurred
	fmt.Printf("%+v\n", ReturningError(3)) // third error occurred with complete stacktrace
}
