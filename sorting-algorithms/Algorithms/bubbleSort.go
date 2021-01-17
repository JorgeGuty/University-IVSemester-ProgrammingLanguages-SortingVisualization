package algorithms

import(
	utility "sorting-visualization/sorting-algorithms/Utility"
)

func BubbleSort( pArray []int, pChannel chan utility.Pair){

	isSorted := false
	unSortedQuantity := len(pArray)

	for isSorted == false {

		isSorted = true
		for j := 1 ; j < unSortedQuantity ; j++ {
			if pArray[j - 1] > pArray[j]{
				utility.Swap( j - 1, j, pArray, pChannel )
				isSorted = false
			}
		}

		unSortedQuantity--
	}
}
