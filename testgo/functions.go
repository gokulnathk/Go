package main

import (
  "fmt"
)

func sayHello(terms ...int) (numTerms int, sum int) {
  for _, i := range terms {
    sum += i
  }
  numTerms = len(terms)
  return
}


func main() {
  nTerms, summation := sayHello(1,3,6,9,12,15,18,65,24)
  fmt.Println("Added ", nTerms, "Elements and the sum is ", summation)
}
