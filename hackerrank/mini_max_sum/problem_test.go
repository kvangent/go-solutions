// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package min_max_sum

import (
	"fmt"
	"testing"
)

func TestProblem(t *testing.T) {
	tcs := []struct {
		input   []int64
		wantMin int64
		wantMax int64
	}{
		{
			[]int64{1, 2, 3, 4, 5},
			10, 14,
		},
	}

	for _, tcs := range tcs {
		t.Run(fmt.Sprintf("solve(%d)", tcs.input), func(t *testing.T) {
			gMin, gMax := miniMaxSum(tcs.input)
			if tcs.wantMin != gMin || tcs.wantMax != gMax {
				t.Fatalf("fail: want: (%d, %d), got (%d, %d)", tcs.wantMin, tcs.wantMax, gMin, gMax)
			}
		})
	}
}
