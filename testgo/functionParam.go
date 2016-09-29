package main

import (
  "./greeting"
  "fmt"
)
func main() {
  s := greeting.Name{"Gokulnath", "K"}
  // greet(s, print)
  fmt.Println()
  // greet(s, println)
  greeting.Greet(s, greeting.CustomMessageFunction("Hello"), false)
}
