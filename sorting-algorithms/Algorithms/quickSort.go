package algorithms

import (
    utility "sorting-visualization/sorting-algorithms/Utility"
)

func QuickSort(arr []int, start int, end int, pChannel chan utility.Pair, pStats *utility.Stats) {
    pStats.Iterations++
    if start >= end {
        return
    }

    index := partition(arr, start, end, pChannel, pStats)

    QuickSort(arr, start, index - 1, pChannel,pStats)
    QuickSort(arr, index + 1, end, pChannel,pStats)
    return
}

func partition(arr []int, start int, end int, pChannel chan utility.Pair,pStats *utility.Stats) int {
    pivotValue := arr[end]
    pivotIndex := start

    for i := start; i < end; i++ {
        pStats.Iterations++
        pStats.Comparisons++
        if arr[i] < pivotValue {
            utility.Swap( i, pivotIndex, arr, pChannel)
            pivotIndex++
        }
    }
    pStats.Iterations++
    utility.Swap( pivotIndex, end, arr, pChannel)
    return pivotIndex
}