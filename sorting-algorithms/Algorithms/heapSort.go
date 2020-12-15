package Algorithms

//HeapSort algorithm
func HeapSort(pArray []int) {
	for i := len(pArray) / 2 - 1; i >= 0; i-- {
		heapify(pArray,len(pArray),i)
	}
	for i := len(pArray)-1; i > 0; i-- {
		Swap(0, i, pArray);
		heapify(pArray,i,0)
	}
}
func heapify(pArray []int, pSize int, pRoot int){
	max := pRoot;
	leftChild := 2 * pRoot + 1;
	rightChild := 2 * pRoot + 2;

	if(leftChild < pSize && pArray[leftChild] > pArray[max]){
		max = leftChild;
	}
	if(rightChild < pSize && pArray[rightChild] > pArray[max]){
		max = rightChild;
	}
	if(max != pRoot){
		Swap(pRoot,max,pArray);
		heapify(pArray,pSize,max);
	}
}