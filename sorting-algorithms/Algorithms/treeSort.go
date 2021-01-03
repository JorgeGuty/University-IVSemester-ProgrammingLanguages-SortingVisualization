package Algorithms
//TreeSortVisualizer consumer function of TreeSort()
func TreeSortVisualizer(pArray []int){
	channel:= make(chan Pair);
	go TreeSort(pArray, channel)
}
// TreeSort Algorithm
func TreeSort(pArray []int, pChannel chan Pair) []int {
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
func siftup(index int, size int, array []int, channel chan Pair) {
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
			channel <- Pair{index, childIndex}
			index = childIndex
			continue
		}
		break
	}
	array[index] = rootValue
	// ? Hay que visualizar a ver si aqui es necesario poner uno
}
//Este es un struct para hacer un tipo tupla que se puede hacer con {}
//https://stackoverflow.com/questions/13670818/pair-tuple-data-type-in-go
//https://stackoverflow.com/questions/17825857/how-to-make-a-channel-that-receive-multiple-return-values-from-a-goroutine
type Pair struct {
    a, b interface{}
}