package common

type Result[T any] struct {
	Value T
	Err   error
}

func (r Result[T]) IsSuccess() bool {
	return r.Err == nil
}

func (r Result[T]) IsFailure() bool {
	return r.Err != nil
}
