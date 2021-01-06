package main

import(
	"sorting-visualization/sorting-algorithms/Utility"
	"sorting-visualization/sorting-algorithms/Algorithms"
	"sync"
)

func AlgorithmVisualizer(pArray []int, pSortID string, pWaitGroup *sync.WaitGroup){
	defer pWaitGroup.Done()
	channel := make(chan utility.Pair)
	go executeAlgorithm(pArray, pSortID, channel)
	var pairToSwap utility.Pair
	for {
		pairToSwap = <- channel
		VisualSwap(pairToSwap.A,pairToSwap.B,pSortID)
	}
}

func executeAlgorithm(pArray []int, pSortID string, pChannel chan utility.Pair){
	switch pSortID {
		case bubbleSortID :
			algorithms.BubbleSort(pArray, pChannel)
			break
		case treeSortID:
			algorithms.TreeSort(pArray, pChannel)
		case heapSortID:
			algorithms.HeapSort(pArray, pChannel)
		//TODO: poner los cases que faltan para cada algoritmo.
	}
}
/*
// TreeSortVisualizer consumer function of TreeSort()
func TreeSortVisualizer(pArray []int, pSortID string, pWaitGroup *sync.WaitGroup){
	defer pWaitGroup.Done()
	channel := make(chan utility.Pair)
	go algorithms.TreeSort(pArray, channel)
	var pairToSwap utility.Pair
	for {
		pairToSwap = <- channel
		VisualSwap(pairToSwap.A,pairToSwap.B,pSortID)
	}
}

// HeapSortVisualizer consumer function of HeapSort()
func HeapSortVisualizer(pArray []int, pSortID string, pWaitGroup *sync.WaitGroup){
	defer pWaitGroup.Done()
	channel := make(chan utility.Pair)
	go algorithms.HeapSort(pArray, channel)
	var pairToSwap utility.Pair
	for {
		pairToSwap = <- channel
		VisualSwap(pairToSwap.A,pairToSwap.B,pSortID)
	}
}

func BubbleSortVisualizer(pArray []int, pSortID string, pWaitGroup *sync.WaitGroup){
	defer pWaitGroup.Done()
	channel := make(chan utility.Pair)
	go algorithms.BubbleSort(pArray, channel)
	var pairToSwap utility.Pair
	for {
		pairToSwap = <- channel
		VisualSwap(pairToSwap.A,pairToSwap.B,pSortID)
	}
}
*/