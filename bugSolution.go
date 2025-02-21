```go
package main

import "fmt"

func main() {
    var m map[string]int
    m = make(map[string]int)
    fmt.Println(m["a"]) // Accessing a key in an initialized map. This will now print 0 as intended, not throw any error
    val, ok := m["b"] // checking for existence of key
    if ok{
        fmt.Printf("Value of b: %v\n",val)
    }else{
        fmt.Printf("Key b does not exists\n")
    }
}
```