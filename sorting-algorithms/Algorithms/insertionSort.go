package algorithms

import(
    //Import for Pair Struct
    //"sorting-visualization/sorting-algorithms/Utility"
    utility "sorting-visualization/sorting-algorithms/Utility"
)

func InsertionSort( pArray []int, pChannel chan utility.Pair) {
    var n = len(pArray)
    for i := 1; i < n; i++ {
        j := i
        for j > 0 {
            if pArray[j-1] > pArray[j] {
                utility.Swap( j - 1, j, pArray, pChannel )
            }
            j = j - 1
        }
    }
}
