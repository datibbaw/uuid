# uuid
uuid library for go

## Example
```golang
package main

import (
    "crypto/rand"
    "fmt"
    "github.com/datibbaw/uuid"
)

func main() {
    s, _ := uuid.CreateV4(rand.Reader)
    fmt.Println(s)
}
```
