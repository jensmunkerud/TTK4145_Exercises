package main
import (
    . "fmt"
    "runtime"
)

func numberServer(increment chan bool, decrement chan bool, get chan int) {
    i := 0
    for {
        select {
        case <-increment:
            i++
        case <-decrement:
            i--
        case get <- i:
            return  // Exit after sending the final value
        }
    }
}

func incrementing(increment chan bool, done chan bool) {
    for j := 0; j < 1000000; j++ {
        increment <- true
    }
    done <- true  // Signal that we're finished
}

func decrementing(decrement chan bool, done chan bool) {
    for j := 0; j < 1000000; j++ {
        decrement <- true
    }
    done <- true  // Signal that we're finished
}

func main() {
    runtime.GOMAXPROCS(2)
    
    // Create channels
    increment := make(chan bool)
    decrement := make(chan bool)
    get := make(chan int)
    done := make(chan bool)
    
    // Start the number server
    go numberServer(increment, decrement, get)
    
    // Start the worker goroutines
    go incrementing(increment, done)
    go decrementing(decrement, done)
    
    // Wait for both workers to finish
    <-done  // Wait for first worker
    <-done  // Wait for second worker
    
    // Get the final value from the server
    finalValue := <-get
    
    Println("The magic number is:", finalValue)
}