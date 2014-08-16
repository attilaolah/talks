# math/big

---

## Fibonacci

Sample from //golang.org/doc/play/fib.go:

```go
package main

import "fmt"

// fib returns a function that returns
// successive Fibonacci numbers.
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fib()
	// Function calls are evaluated left-to-right.
	fmt.Println(f(), f(), f(), f(), f())
}
```


---

## More Fibonacci

```go
func main() {
	f := fib()
	for {
		fmt.Println(f())
	}
}
```

`fib()` #93 overflows: -6246583658587674878!

---

## Big Fibonacci

```go
func fib() func() *big.Int {
	a, b := big.NewInt(0), big.NewInt(1)
	return func() *big.Int {
		a, b = b, a.Add(a, b)
		return a
	}
}

func main() {
	f := fib()
	for {
		fmt.Println(f().String())
	}
}
```

* `big.Int` is actually a struct:

```go
// An Int represents a signed multi-precision integer.
// The zero value for an Int represents the value 0.
type Int struct {
	neg bool // sign
	abs nat  // absolute value of the integer
}
```

* methods are implemented on pointers
* as such, we usually work with `*big.Int`
  * faster, no need to copy the data
  * **mutable**, must be careful!

---

## Printing an encoding

* `big.Int` is a struct, cannot be printed natively
* `*big.Int` [implements][1] [`fmt.Formatter`][2], can be printed:

```go
fmt.Println(big.NewInt(37))
```

* `*big.Int` also implements
  * [`json.Marshaler`][4] and [`json.Unmarshaler`][5]
  * [`encoding.TextMarshaler`][6] and [`encoding.TextUnmarshaler`][7]
  * [`gob.GobEncoder`][8] and [`gob.GobDecoder`][9]

[1]: http://golang.org/pkg/math/big/#Int.Format
[2]: http://golang.org/pkg/fmt/#Formatter
[4]: http://golang.org/pkg/encoding/json/#Marshaler
[5]: http://golang.org/pkg/encoding/json/#Unmarshaler
[6]: http://golang.org/pkg/encoding/#TextMarshaler
[7]: http://golang.org/pkg/encoding/#TextUnmarshaler
[8]: http://golang.org/pkg/encoding/gob/#GobEncoder
[9]: http://golang.org/pkg/encoding/gob/#GobDecoder

---

## Integration with other systems

* JSON encode/decode works very well (more explicit but type safe):

### Python:

```python
import json

print json.dumps({"result": 99*99})
```

### Ruby:

```ruby
require 'json'

puts JSON.dump result: 99**99
```

### Go:

```go
package main

import (
	"encoding/json"
	"math/big"
	"os"
)

type Response struct {
	Result *big.Int `json:"result,omitempty"`
}

func main() {
	json.NewEncoder(os.Stdout).Encode(struct {
		Result *big.Int `json:"result"`
	}{
		Result: big.NewInt(0).Exp(big.NewInt(99), big.NewInt(99), nil),
	})
}
```
