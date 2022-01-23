package catlib

import "fmt"

type Result[T any] struct {
	t   T
	err error
}

func Ok[T any](t T) Result[T] {
	return Result[T]{t: t, err: nil}
}

func Err[T any](err error) Result[T] {
	return Result[T]{err: err}
}

func (t Result[T]) String() string {
	if t.err == nil {
		return fmt.Sprintf("Ok(%v)", t.t)
	}

	return fmt.Sprintf("Err(%s)", t.err.Error())
}

var _ fmt.Stringer = Result[struct{}]{}

func (t Result[T]) Unwrap() T {
	if t.err == nil {
		return t.t
	}

	panic(fmt.Sprintf(`unwrap "%s"`, t))
}

func (t Result[T]) UnwrapOr(s T) T {
	if t.err == nil {
		return t.t
	}

	return s
}

func (t Result[T]) UnwrapOrElse(f func() T) T {
	if t.err == nil {
		return t.t
	}

	return f()
}

func (t Result[T]) UnwrapOrZero() T {
	if t.err == nil {
		return t.t
	}

	var s T
	return s
}

func (t Result[T]) Value() (T, error) {
	return t.t, t.err
}

func (t Result[T]) IsOk() bool {
	return t.err == nil
}

func (t Result[T]) IsErr() bool {
	return t.err != nil
}

func (t Result[T]) Expect(message string) T {
	if t.err == nil {
		return t.t
	}

	panic(message)
}
