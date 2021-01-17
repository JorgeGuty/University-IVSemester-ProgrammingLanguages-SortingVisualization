package algorithms

import(
    utility "sorting-visualization/sorting-algorithms/Utility"
)

func SelectionSort( pArray []int, pChannel chan utility.Pair) {
    var n = len(pArray)
    for i := 0; i < n; i++ {
        var minIdx = i
        for j := i; j < n; j++ {
            if pArray[j] < pArray[minIdx] {
                minIdx = j
            }
        }
        utility.Swap(i, minIdx, pArray, pChannel )
    }
}

