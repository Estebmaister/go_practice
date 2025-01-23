package errors

import "fmt"

const (
	Pink = "pink"
	Red  = "red"
)

type ErrorPink struct{}
type ErrorRed struct{}

func (e ErrorPink) Error() string { return "can't believe it's pink" }
func (e ErrorRed) Error() string  { return "can't see red" }

func ErrorTrigger() error {
	err := ColorBlind(Pink)

	if err != nil {
		switch err.(type) {
		case ErrorPink:
			return fmt.Errorf("unidentified color: %w", err)
		case ErrorRed:
			return fmt.Errorf("maybe green?: %w", err)
		}
		return nil
	}

	return nil
}

func ColorBlind(color string) error {
	if color == Pink {
		return ErrorPink{}
	}
	if color == Red {
		return ErrorRed{}
	}
	return nil
}
