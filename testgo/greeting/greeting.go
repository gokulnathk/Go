package greeting

import (
  "fmt"
)

type Name struct {
  Firstname string
  Lastname string
}

type Printer func(string) ()

func Greet(name Name, do Printer, displayLast bool) {
// func greet(name Name, do func(string)) {
  if prefix := "Mr "; displayLast {
    do(name.Lastname)
  } else {
    do(prefix + name.Firstname)
  }
}

func print(s string) {
  fmt.Print(s)
}

func println(s string) {
  fmt.Println(s)
}

func CustomMessageFunction (custom string) Printer {
  return func (s string) {
    fmt.Println(custom + " " + s)
  }
}
