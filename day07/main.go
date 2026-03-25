package main

import (
	"github.com/sotiri-geo/aod-2025/common"
)

// SplitBeam returns a new set with index positions if a beam index hits a
// splitter index. This algorithm runs with time complexity O(N).
func SplitBeam(beamIndex, splitterIndex *common.Set[int], size int) *common.Set[int] {
	output := common.NewSet([]int{})

	for i := range beamIndex.Set {
		if splitterIndex.Has(i) {
			left, right := i-1, i+1
			if left >= 0 {
				output.Add(left)
			}
			if right < size {
				output.Add(right)
			}
		}
	}
	return output
}
