package main

import (
  "fmt"
)

func main() {
  foo := new(myStruct)
  bar := &myStruct{}
  baz := myStruct{"Test", 3}
  foo.myField = "Hello"
  bar.myField ="World"
  foo.myInt = 1
  bar.myInt = 2
  fmt.Println(foo.myField)
  fmt.Println(foo.myInt)
  fmt.Println(foo)
  fmt.Println(bar.myField)
  fmt.Println(bar.myInt)
  fmt.Println(bar)
  fmt.Println(baz)
}

type myStruct struct {
  myField string
  myInt int
}
