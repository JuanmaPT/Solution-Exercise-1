// Use `go run foo.go` to run your program

package main

import (
    "fmt"
    "runtime"
    "time"
)

var i = 0

func incrementing() {
    //TODO: increment i 1000000 times
    for a := 0; a < 1000000; i++ {
        i = i + 1
    }
}

func decrementing() {
    //TODO: decrement i 1000000 times
    for a := 0; a < 1000000; i++ {
        i = i + 1
    }
}

func getGOMAXPROCS() int {
    return runtime.GOMAXPROCS(0)
}

func main() {
    //fmt.Printf("GOMAXPROCS before assigning nCPUs is %d\n", getGOMAXPROCS())
    runtime.GOMAXPROCS(runtime.NumCPU())    // I guess this is a hint to what GOMAXPROCS does...
	                                    // Try doing the exercise both with and without it!
    //fmt.Printf("GOMAXPROCS after assigning nCPUs is %d\n", getGOMAXPROCS())
    // TODO: Spawn both functions as goroutines
	go incrementing()
    go decrementing()
    // We have no way to wait for the completion of a goroutine (without additional syncronization of some sort)
    // We'll come back to using channels in Exercise 2. For now: Sleep.
    time.Sleep(100*time.Millisecond)
    fmt.println("The magic number is:", i)
}
