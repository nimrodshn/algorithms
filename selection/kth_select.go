package selection

import (
	"github.com/nimrodshn/algorithms/partition"
)

// PartitionSelect uses the Lamotu (random variant) partition scheme
// to select the k'th elemnt in the input array in avarage linear time.
func PartitionSelect(input []int, k int) int {
	pivotIdx := partition.RandomPartition(input, 0, len(input)-1)
	if pivotIdx == k {
		return input[pivotIdx]
	}
	if k < pivotIdx {
		return PartitionSelect(input[:pivotIdx], k)
	}
	return PartitionSelect(input[pivotIdx+1:], k-pivotIdx)
}
