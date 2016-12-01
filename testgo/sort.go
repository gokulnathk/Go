package main

import (
  "learning/testgo/sort"
  "fmt"
)

func main() {
    unsort := []int {5, 7, 3, 2 ,4 ,8, 9, 12, 1}
    r := sort.BubbleSort {unsort}
    sorted := r.Sort();
    fmt.Println(sorted);
}
