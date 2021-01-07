package main

import(
	"github.com/pusher/pusher-http-go"
	"sorting-visualization/sorting-algorithms/Utility"
	"sorting-visualization/sorting-algorithms/Algorithms"
	"sync"
)

func AlgorithmVisualizer(pArray []int, pSortID string, pWaitGroup *sync.WaitGroup, pPusherClient pusher.Client){

	defer pWaitGroup.Done()
	defer pPusherClient.Trigger("arrayVisualization", "solved", solvedEventElement{pSortID})

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
		//TODO: meter el case del quicksort.
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