package algorithms

import(
	utility "sorting-visualization/sorting-algorithms/Utility"
)

func BubbleSort( pArray []int, pChannel chan utility.Pair, pStats *utility.Stats){

	isSorted := false
	unSortedQuantity := len(pArray)

	for isSorted == false {
		pStats.Iterations++
		isSorted = true
		for j := 1 ; j < unSortedQuantity ; j++ {
			pStats.Iterations++
			pStats.Comparisons++
			if pArray[j - 1] > pArray[j]{
				utility.Swap( j - 1, j, pArray, pChannel )
				isSorted = false
			}
		}
		pStats.Iterations++
		unSortedQuantity--
	}
	pStats.Iterations++
}
