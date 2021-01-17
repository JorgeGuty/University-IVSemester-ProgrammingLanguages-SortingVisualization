package main

import (
	"sorting-visualization/sorting-algorithms/Algorithms"
	"sorting-visualization/sorting-algorithms/Utility"
	"sync"
	"time"
)


// AlgorithmVisualizer calls the specified function
func AlgorithmVisualizer(pArray []int, pSortID string, pWaitGroup *sync.WaitGroup){

	defer pWaitGroup.Done()
	defer VisualDone(pSortID)

	swaps := 0

	channel := make(chan utility.Pair)

	start := time.Now()

	go executeAlgorithm(pArray, pSortID, channel)

	for pairToSwap := range channel{	
		swaps++
		VisualSwap(pairToSwap.A,pairToSwap.B,pSortID)
	}
	elapsed := float64(time.Now().Sub(start)/time.Millisecond)
	ShowStats(elapsed,0,swaps,0,pSortID)
}

func executeAlgorithm(pArray []int, pSortID string, pChannel chan utility.Pair){

	switch pSortID {
		case bubbleSortID :
			algorithms.BubbleSort(pArray, pChannel)
			break
		case treeSortID:
			algorithms.TreeSort(pArray, pChannel)
			break
		case heapSortID:
			algorithms.HeapSort(pArray, pChannel)
			break
		case selectionSortID:
			algorithms.SelectionSort(pArray, pChannel)
			break
		case insertionSortID:
			algorithms.InsertionSort(pArray, pChannel)
			break
		case quickSortID:
			algorithms.QuickSort(pArray, 0, len(pArray)-1, pChannel)
			break
	}
	close(pChannel)
}