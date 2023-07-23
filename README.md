
# Assert for Go 1.18+ test assertions

## Install

```sh
go get github.com/weiwenchen2022/assert
```

## Usage

```go
package main

import (
    "testing"
    "github.com/weiwenchen2022/assert"
)

func greet(name string) (string, int) {
    greeting := fmt.Sprintf("Hello %s", name)

    // Return the greeting and its length (in bytes).
    return greeting, len(greeting)
}

func TestGreet(t *testing.T) {
    greeting, greetingLength := greet("Alice")

    assert.Equal(t, "Hello Alice", greeting, "greeting:")
    assert.Equal(t, 11, greetingLength, "greetingLength:")
}
```

## Reference

GoDoc: [https://godoc.org/github.com/weiwenchen2022/assert](https://godoc.org/github.com/weiwenchen2022/assert)
