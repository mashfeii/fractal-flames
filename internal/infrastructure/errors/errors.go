package errors

import "fmt"

type ErrInvalidIntegerFlag struct{}

func NewErrInvalidIntegerFlag() error {
	return ErrInvalidIntegerFlag{}
}

func (e ErrInvalidIntegerFlag) Error() string {
	return "Negative values is not allowed"
}

type ErrEmptyTransitions struct {
	Value []int
}

func NewErrEmptyTransitions(val []int) error {
	return ErrEmptyTransitions{Value: val}
}

func (e ErrEmptyTransitions) Error() string {
	return fmt.Sprintf("No possible transitions for indexes: %v", e.Value)
}