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
	stats := utility.Stats{}
	 

	channel := make(chan utility.Pair)

	start := time.Now()

	go executeAlgorithm(pArray, pSortID, channel, &stats)

	for pairToSwap := range channel{	
		swaps++
		VisualSwap(pairToSwap.A,pairToSwap.B,pSortID)
	}
	elapsed := float64(time.Now().Sub(start)/time.Millisecond)
	ShowStats(elapsed,stats.Iterations,swaps,stats.Comparisons,pSortID)
}

func executeAlgorithm(pArray []int, pSortID string, pChannel chan utility.Pair, pStats *utility.Stats){

	switch pSortID {
		case bubbleSortID :
			algorithms.BubbleSort(pArray, pChannel,pStats)
			break
		case treeSortID:
			algorithms.TreeSort(pArray, pChannel,pStats)
			break
		case heapSortID:
			algorithms.HeapSort(pArray, pChannel,pStats)
			break
		case selectionSortID:
			algorithms.SelectionSort(pArray, pChannel,pStats)
			break
		case insertionSortID:
			algorithms.InsertionSort(pArray, pChannel,pStats)
			break
		case quickSortID:
			algorithms.QuickSort(pArray, 0, len(pArray)-1, pChannel,pStats)
			break
	}
	close(pChannel)
}