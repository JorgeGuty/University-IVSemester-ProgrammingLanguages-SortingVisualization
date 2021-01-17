package algorithms

import (
    utility "sorting-visualization/sorting-algorithms/Utility"
)

func QuickSort(arr []int, start int, end int, pChannel chan utility.Pair) {
    if start >= end {
        return
    }

    index := partition(arr, start, end, pChannel)

    QuickSort(arr, start, index - 1, pChannel)
    QuickSort(arr, index + 1, end, pChannel)
    return
}

func partition(arr []int, start int, end int, pChannel chan utility.Pair) int {
    pivotValue := arr[end]
    pivotIndex := start

    for i := start; i < end; i++ {
        if arr[i] < pivotValue {
            utility.Swap( i, pivotIndex, arr, pChannel)
            pivotIndex++
        }
    }
    utility.Swap( pivotIndex, end, arr, pChannel)
    return pivotIndex
}