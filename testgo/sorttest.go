package main

import (
  "fmt"
)

// Struct BubbleSort (similar to Class in OOP Languages)
type InsertionSort struct {
    // Slice (Dynamically increasing array alternative in Go)
    Unsorted []int;
}

func (is InsertionSort) greaterThan (a int, b int) bool {
    return a > b;
}

func (is InsertionSort) lesserThan (a int, b int) bool {
    return a < b;
}


// Sorting function belongs to struct BubbleSort
func (is InsertionSort) Sort(mode string) []int {
    length := len(is.Unsorted);

    funcs := map[string] func(int, int) bool {"method":is.greaterThan};
    if mode == "desc" {
        funcs = map[string] func(int, int) bool {"method":is.lesserThan};
    }

    for i := 0; i < length; i++ {
        temp := is.Unsorted[i];
        j := i - 1;
        for ; j >=0 && funcs["method"](is.Unsorted[j], temp); j-- {
            is.Unsorted[j + 1] = is.Unsorted[j];
        }
        is.Unsorted[j + 1] = temp;
    }

    return is.Unsorted;
}


func main() {
    unsort := []int {5, 7, 3, 2 ,4 ,8, 9, 12, 1};
    i := InsertionSort {unsort};
    sorted := i.Sort("asc");
    fmt.Println(sorted);
}
