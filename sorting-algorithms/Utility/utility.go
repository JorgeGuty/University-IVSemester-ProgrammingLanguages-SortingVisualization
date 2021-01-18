package utility

// Pair es un struct para hacer un tipo tupla que se puede hacer con {}
type Pair struct {
    A int
    B int
}
// Stats is an struct used to store the number of iterations and the number of comparisons made by each algorithm
type Stats struct {
    Comparisons int
    Iterations int
}

// Swap Utility with channel 
func Swap(pIndex1 int, pIndex2 int, array []int, channel chan Pair) {
	temp := array[pIndex1]
	array[pIndex1] = array[pIndex2]
    array[pIndex2] = temp
    channel <- Pair{A: pIndex1, B: pIndex2}
}