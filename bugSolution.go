```go
package main

import (

    "fmt"
)

func main() {
    var m = make(map[string]int)
    m["a"] = 1
    fmt.Println(m["b"]) // Output: 0.  This is expected.

    var m2 map[string]int
    if m2 != nil {
        fmt.Println(m2["b"])
    } else {
        fmt.Println("Map is nil. Cannot access key.") // Handle the nil map case.
    }
}
```