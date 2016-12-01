package sort

// Struct BubbleSort (similar to Class in OOP Languages)
type BubbleSort struct {
    // Slice (Dynamically increasing array alternative in Go)
    Unsorted []int;
}

// Sorting function belongs to struct BubbleSort
func (bs BubbleSort) Sort() []int {
    length := len(bs.Unsorted) - 1;

    // loop through the passes until we get the sorted response
    for {
        sorted := false;
        // sorting for the each pass
        for i := 0; i < length; i++ {
            if bs.Unsorted[i] > bs.Unsorted[i + 1] {
                temp := bs.Unsorted[i];
                bs.Unsorted[i] = bs.Unsorted[i + 1];
                bs.Unsorted[i + 1] = temp;
                sorted = true;
            }
        }

        // If there is no swap happened, it is sorted result
        if sorted == false {
            break;
        }
    }
    return bs.Unsorted;
}
