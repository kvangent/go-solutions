// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package subarray_division_1

import (
	"fmt"
	"testing"
)

func TestProblem(t *testing.T) {
	tcs := []struct {
		s    []int
		d, m int
		want int
	}{
		{
			[]int{4},
			4, 1,
			1,
		},
		{
			[]int{2,2,1,3,2},
			4, 2,
			2,
		},
	}

	for i, tcs := range tcs {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := birthday(tcs.s, tcs.d, tcs.m)
			if got != tcs.want {
				t.Fatalf("fail: want: %d, got: %d", tcs.want, got)
			}
		})
	}
}
