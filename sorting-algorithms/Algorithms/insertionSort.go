package algorithms

import(
    utility "sorting-visualization/sorting-algorithms/Utility"
)

func InsertionSort( pArray []int, pChannel chan utility.Pair, pStats *utility.Stats) {
    var n = len(pArray)
    
    for i := 1; i < n; i++ {
        pStats.Iterations++
        j := i
        for j > 0 {
            pStats.Iterations++
            pStats.Comparisons++
            if pArray[j-1] > pArray[j] {
                utility.Swap( j - 1, j, pArray, pChannel )
            }
            j = j - 1
        }
        pStats.Iterations++
    }
    pStats.Iterations++
}
