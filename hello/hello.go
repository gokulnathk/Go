package main

import "fmt"

func main() {	
    fmt.Printf("Calculating Sum\n")
    sum := 0
    size := 2147483647
    for i := 0; i < size; i++ {
        sum += i
    }
    fmt.Println(sum)
}