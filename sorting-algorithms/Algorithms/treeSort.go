package Algorithms
// TreeSort Algorithm
func TreeSort(pArray []int) []int {
	lenght := len(pArray)
	arrayToUse := append([]int{0}, pArray...)

	for i := lenght / 2; i >= 2; i-- {
		siftup(i, lenght, arrayToUse)
	}
	for i := lenght; i >= 2; i-- {
		siftup(1, i, arrayToUse)
		Swap(1, i, arrayToUse)
	}

	pArray = arrayToUse[1:]
	return pArray
}
// Swap Utility
func Swap(x int, y int, array []int) {
	temp := array[x]
	array[x] = array[y]
	array[y] = temp
}

func siftup(index int, size int, array []int) {
	rootValue := array[index]
	var sibling int
	for childIndex := 2 * index; childIndex <= size; childIndex = 2 * index {
		if childIndex < size {
			sibling = childIndex + 1
			if array[sibling] > array[childIndex] {
				childIndex = sibling
			}
		}
		if array[childIndex] > rootValue {
			array[index] = array[childIndex]
			index = childIndex
			continue
		}
		break
	}
	array[index] = rootValue
}