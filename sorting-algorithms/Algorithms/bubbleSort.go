package Algorithms

import (
	"sync"
)

func bubbleSort ( pArray []int, pWaitGroup *sync.WaitGroup ) []int{

	isSorted := false
	unSortedQuantity := len(pArray)

	for isSorted == false {

		isSorted = true
		for j := 1 ; j < unSortedQuantity ; j++ {
			if pArray[j - 1] > pArray[j]{
				pivot := pArray[j]
				pArray[j] = pArray[j - 1]
				pArray[j - 1] = pivot
				isSorted = false
			}
		}

		unSortedQuantity--
	}

	return pArray
}
