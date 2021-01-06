package algorithms

import(
    //Import for Pair Struct
    //"sorting-visualization/sorting-algorithms/Utility"
)

func selectionsort(items []int) {
    var n = len(items)
    for i := 0; i < n; i++ {
        var minIdx = i
        for j := i; j < n; j++ {
            if items[j] < items[minIdx] {
                minIdx = j
            }
        }
        items[i], items[minIdx] = items[minIdx], items[i]
    }
}

