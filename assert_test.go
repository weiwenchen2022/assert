package assert_test

import (
	"strings"
	"testing"

	"github.com/weiwenchen2022/assert"
)

func TestEqual(t *testing.T) {
	tests := []struct {
		name     string
		actual   func() string
		expected string
	}{
		{
			name: "Empty Builder",
			actual: func() string {
				var b strings.Builder
				return b.String()
			},
			expected: "",
		},
		{
			name: "WriteString",
			actual: func() string {
				var b strings.Builder
				b.WriteString("hello")
				return b.String()
			},
			expected: "hello",
		},
		{
			name: "WriteByte",
			actual: func() string {
				var b strings.Builder
				b.WriteString("hello")
				b.WriteByte(' ')
				return b.String()
			},
			expected: "hello ",
		},
		{
			name: "WriteString",
			actual: func() string {
				var b strings.Builder
				b.WriteString("hello")
				b.WriteByte(' ')
				b.WriteString("world")
				return b.String()
			},
			expected: "hello world",
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, tt.actual(), tt.name)
	}
}

func TestPanic(t *testing.T) {
	tests := []struct {
		name      string
		fn        func()
		wantPanic bool
	}{
		{
			name:      "panic",
			fn:        func() { panic("panic") },
			wantPanic: true,
		},
		{
			name:      "no painc",
			fn:        func() {},
			wantPanic: false,
		},
	}

	for _, tt := range tests {
		assert.Panic(t, tt.fn, tt.wantPanic, tt.name)
	}
}
