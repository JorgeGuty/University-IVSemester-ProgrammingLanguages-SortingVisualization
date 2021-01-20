package algorithms

import(
	"sorting-visualization/sorting-algorithms/Utility"
)

// TreeSort Algorithm
func TreeSort(pArray []int, pChannel chan utility.Pair, pStats *utility.Stats) []int {
	lenght := len(pArray) 
	for i := (lenght / 2) - 1 ; i > 0; i-- {
		pStats.Iterations++
		siftup(i, (lenght - 1), pArray, pChannel,pStats)
	}
	pStats.Iterations++
	for i := (lenght - 1); i > 0; i-- {
		pStats.Iterations++
		siftup(0, i, pArray, pChannel,pStats)
		utility.Swap(0, i, pArray,pChannel)
	}
	pStats.Iterations++
	return pArray
}
func siftup(index int, size int, array []int, channel chan utility.Pair,pStats *utility.Stats) {
	rootValue := array[index]
	var sibling int
	for childIndex := (2 * index) + 1; childIndex <= size; childIndex = (2 * index) + 1 {
		pStats.Iterations++
		
		pStats.Comparisons++
		if childIndex < size {
			sibling = childIndex + 1
			pStats.Comparisons++
			if array[sibling] > array[childIndex] {
				childIndex = sibling
			}
		}
		pStats.Comparisons++
		if array[childIndex] > rootValue {
			array[index] = array[childIndex]
			channel <- utility.Pair{A: index, B: childIndex}
			index = childIndex
			continue
		}
		break
	}
	pStats.Iterations++
	array[index] = rootValue
}