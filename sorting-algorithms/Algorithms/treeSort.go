package Algorithms

import(
	"sorting-visualization/sorting-algorithms/Utility"
	"sorting-visualization/visualization"
)

//TreeSortVisualizer consumer function of TreeSort()
func TreeSortVisualizer(pArray []int, pSortID string){
	channel:= make(chan utility.Pair);
	go TreeSort(pArray, channel)
	var pairToSwap utility.Pair;
	for {
		pairToSwap = <- channel;
		main.VisualSwap(pairToSwap.A,pairToSwap.B,pSortID)
	}
}
// TreeSort Algorithm
func TreeSort(pArray []int, pChannel chan utility.Pair) []int {
	lenght := len(pArray) 
	for i := (lenght / 2) - 1 ; i > 0; i-- {
		siftup(i, (lenght - 1), pArray, pChannel)
	}
	for i := (lenght - 1); i > 0; i-- {
		siftup(0, i, pArray, pChannel)
		Swap(0, i, pArray)
	}
	return pArray
}
// Swap Utility
func Swap(x int, y int, array []int) {
	temp := array[x]
	array[x] = array[y]
	array[y] = temp
}
func siftup(index int, size int, array []int, channel chan utility.Pair) {
	rootValue := array[index]
	var sibling int
	for childIndex := (2 * index) + 1; childIndex <= size; childIndex = (2 * index) + 1 {
		if childIndex < size {
			sibling = childIndex + 1
			if array[sibling] > array[childIndex] {
				childIndex = sibling
			}
		}
		if array[childIndex] > rootValue {
			//No es exactamente un swap pero se puede mostrar asi
			array[index] = array[childIndex]
			//swap
			channel <- utility.Pair{A: index, B: childIndex}
			index = childIndex
			continue
		}
		break
	}
	array[index] = rootValue
	// ? Hay que visualizar a ver si aqui es necesario poner uno
}