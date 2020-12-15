package Algorithms

import (
	"sync"
	"testing"
)

func TestBubbleSort(t *testing.T) {

	var waitGroup sync.WaitGroup

	go bubbleSort([]int{0,1,0,1,0,1,0,1,0,1,0,1,0,1,0,1,0,1})
	go bubbleSort([]int{2,3,2,3,2,3,2,3,2,3,2,3,2,3,2,3,2,3})
	go bubbleSort([]int{5,6,5,6,5,6,5,6,5,6,5,6,5,6,5,6,5,6})

	waitGroup.Add(3)
	waitGroup.Wait()

}
