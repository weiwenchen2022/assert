/*
Package assert defines helper functions for test assertions.

Using helpers for your test assertions can help to:

  - Make your test functions clean and clear;
  - Keep test failure messages consistent;
  - And reduce the potential for errors in your code due to typos.
*/
package assert

import (
	"testing"
)

func Equal[T comparable](tb testing.TB, expected, actual T, format string, args ...any) {
	tb.Helper()

	if expected != actual {
		tb.Errorf(format+" got: %#v; want: %#v", append(args, any(actual), any(expected))...)
	}
}

func Panic(tb testing.TB, fn func(), wantPanic bool, format string, args ...any) {
	tb.Helper()

	didPanic := make(chan bool, 1)
	go func() {
		defer func() { didPanic <- recover() != nil }()
		fn()
	}()
	if got := <-didPanic; wantPanic != got {
		if wantPanic {
			tb.Errorf(format+" got no panic, want panic", args...)
		} else {
			tb.Errorf(format+" got panic, want no panic", args...)
		}
	}
}
