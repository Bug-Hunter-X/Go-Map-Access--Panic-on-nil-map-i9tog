```go
func main() {
    var m = make(map[string]int)
    m["a"] = 1
    fmt.Println(m["b"]) // Output: 0.  This is expected.

    var m2 map[string]int
    fmt.Println(m2["b"]) // This will panic! Because the map is nil.
}
```