package utility

// Pair es un struct para hacer un tipo tupla que se puede hacer con {}
//https://stackoverflow.com/questions/13670818/pair-tuple-data-type-in-go
//https://stackoverflow.com/questions/17825857/how-to-make-a-channel-that-receive-multiple-return-values-from-a-goroutine
type Pair struct {
    A int
    B int
}

// Swap Utility with channel 
func Swap(x int, y int, array []int, channel chan Pair) {
	temp := array[x]
	array[x] = array[y]
    array[y] = temp
    channel <- Pair{A: x, B: y}
}