package algorithms

import(
	"sorting-visualization/sorting-algorithms/Utility"
)

//HeapSort algorithm
func HeapSort(pArray []int, pChannel chan utility.Pair, pStats *utility.Stats) {
	for i := len(pArray) / 2 - 1; i >= 0; i-- {
		pStats.Iterations++
		heapify(pArray,len(pArray),i,pChannel,pStats)
	}
	pStats.Iterations++
	for i := len(pArray)-1; i > 0; i-- {
		pStats.Iterations++
		utility.Swap(0, i, pArray,pChannel);
		heapify(pArray,i,0,pChannel,pStats)
	}
	pStats.Iterations++
}
func heapify(pArray []int, pSize int, pRoot int, pChannel chan utility.Pair, pStats *utility.Stats){
	max := pRoot;
	leftChild := 2 * pRoot + 1;
	rightChild := 2 * pRoot + 2;

	pStats.Comparisons += 2

	if(leftChild < pSize && pArray[leftChild] > pArray[max]){
		max = leftChild;
	}
	if(rightChild < pSize && pArray[rightChild] > pArray[max]){
		max = rightChild;
	}

	pStats.Iterations++

	if(max != pRoot){
		utility.Swap(pRoot,max,pArray,pChannel);
		heapify(pArray,pSize,max, pChannel,pStats);
	}
}