package visualizer

import(
	"sorting-visualization/sorting-algorithms/Utility"
	"sorting-visualization/sorting-algorithms/Algorithms"
	"sync"
)

// TreeSortVisualizer consumer function of TreeSort()
func TreeSortVisualizer(pArray []int, pSortID string, pWaitGroup sync.WaitGroup){
	channel:= make(chan utility.Pair);
	go Algorithms.TreeSort(pArray, channel)
	var pairToSwap utility.Pair;
	for {
		pairToSwap = <- channel;
		VisualSwap(pairToSwap.A,pairToSwap.B,pSortID)
	}

	//Dice que unreachable code. Hay que revisar el for a ver si es que nunca termina o porque dice eso
	pWaitGroup.Done()
}
// HeapSortVisualizer consumer function of HeapSort()
func HeapSortVisualizer(pArray []int, pSortID string, pWaitGroup sync.WaitGroup){
	channel:= make(chan utility.Pair);
	go Algorithms.HeapSort(pArray, channel)
	var pairToSwap utility.Pair;
	for {
		pairToSwap = <- channel;
		VisualSwap(pairToSwap.A,pairToSwap.B,pSortID)
	}
}