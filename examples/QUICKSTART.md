# 🎓 Quick Start
```go
package main

import (
	"github.com/emacsified/r2d2"
	"fmt"
)

func check(err error) {
	if err != nil {
		panic(err)
    }
}

func main() {
    spaceTrader := r2d2.New("<token>", "<username>")
    status, err := spaceTrader.ApiStatus()
    check(err)
    fmt.Printf("API Status: %s", status)
}
```
