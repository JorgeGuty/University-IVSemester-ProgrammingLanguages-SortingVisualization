package algorithms

import(
    utility "sorting-visualization/sorting-algorithms/Utility"
)

func SelectionSort( pArray []int, pChannel chan utility.Pair,pStats *utility.Stats) {
    var n = len(pArray)
    for i := 0; i < n; i++ {
        pStats.Iterations++
        var minIdx = i
        for j := i; j < n; j++ {
            pStats.Iterations++
            pStats.Comparisons++
            if pArray[j] < pArray[minIdx] {
                minIdx = j
            }
        }
        pStats.Iterations++
        utility.Swap(i, minIdx, pArray, pChannel )
    }
    pStats.Iterations++
}

