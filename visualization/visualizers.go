package visualizer

import(
	"sorting-visualization/sorting-algorithms/Utility"
	"sorting-visualization/sorting-algorithms/Algorithms"
)

// TreeSortVisualizer consumer function of TreeSort()
func TreeSortVisualizer(pArray []int, pSortID string){
	channel:= make(chan utility.Pair);
	go Algorithms.TreeSort(pArray, channel)
	var pairToSwap utility.Pair;
	for {
		pairToSwap = <- channel;
		VisualSwap(pairToSwap.A,pairToSwap.B,pSortID)
	}
}