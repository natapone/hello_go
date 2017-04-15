// $ cd /Users/npc/src/npc/go/src/hello_go
// $ go install hello_go/hello
// $ hello

package main

import (
    "fmt"
    "time"
    "hello_go/stringutil"
    "math/rand"
    "math"
    "math/cmplx"
    )

var (
    c, python, java bool
    ToBe    bool       = false
    MaxInt  uint64     = 1<<64 - 1
    z       complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
    fmt.Println("Hello, 世界, สวัสดี")
    fmt.Println("The time is", time.Now())
    fmt.Println(stringutil.Reverse("!oG ,olleH"))
    
    rand.Seed(time.Now().UTC().UnixNano())
    fmt.Println("My favorite number is", rand.Intn(10))
    
    fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
    
    fmt.Println(math.Pi)
    fmt.Println("555 + 666 = ", add(555, 666))
    
    a, b := swap("first", "second")
    fmt.Println(a, b)
    
    // var i int
    // fmt.Println(i, c, python, java)
    
    var i, j int = 1, 2
    k := 3
    c, python, java := true, false, "no!"
    fmt.Println(i, j, k, c, python, java)
    
    fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
    fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
    fmt.Printf("Type: %T Value: %v\n", z, z)
    
    // big := 1 << 100
    // fmt.Println("big === ", big)
}

func add(x, y int) int {
    return x + y
}

func swap(a, b string) (string, string) {
    return b, a
}

