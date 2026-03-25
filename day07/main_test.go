package main_test

import (
	"testing"

	main "github.com/sotiri-geo/aod-2025/day07"

	"github.com/sotiri-geo/aod-2025/common"
)

func TestSplitBeam(t *testing.T) {
	testCases := map[string]struct {
		beamIndex  *common.Set[int]
		splitIndex *common.Set[int]
		size       int
		want       *common.Set[int]
	}{
		"SingleBeamWithSingleSplit": {
			beamIndex:  common.NewSet([]int{1}),
			splitIndex: common.NewSet([]int{1}),
			size:       3,
			want:       common.NewSet([]int{0, 2}),
		},
		"DoubleBeamWithSingleSplit": {
			beamIndex:  common.NewSet([]int{1, 2}),
			splitIndex: common.NewSet([]int{1}),
			size:       3,
			want:       common.NewSet([]int{0, 2}),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// GIVEN beam indexes AND split indexs
			// WHEN beam index hits a split index
			got := main.SplitBeam(tc.beamIndex, tc.splitIndex, tc.size)

			// THEN beam will split
			if !got.Equal(tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
