package algorithms
import (
    "math/rand"
    utility "sorting-visualization/sorting-algorithms/Utility"

    //Import for Pair Struct
    //"sorting-visualization/sorting-algorithms/Utility"
)
func QuickSort( pArray []int, pChannel chan utility.Pair) []int {
    if len(pArray) < 2 {
        return pArray
    }
     
    left, right := 0, len(pArray)-1
     
    pivot := rand.Int() % len(pArray)

    pArray[pivot], pArray[right] = pArray[right], pArray[pivot]
     
    for i, _ := range pArray {
        if pArray[i] < pArray[right] {
            pArray[left], pArray[i] = pArray[i], pArray[left]
            left++
        }
    }

    pArray[left], pArray[right] = pArray[right], pArray[left]

    QuickSort(pArray[:left], pChannel)
    QuickSort(pArray[left+1:], pChannel)

    return pArray
}
