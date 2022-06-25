// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package max_min

import (
	"fmt"
	"testing"
)

func TestProblem(t *testing.T) {
	tcs := []struct {
		k    int
		arr  []int
		want int
	}{
		{
			4,
			[]int{1, 2, 3, 4, 10, 20, 30, 40, 100, 200},
			3,
		},
		{
			3,
			[]int{10, 100, 300, 200, 1000, 20, 30 },
			20,
		},
	}

	for _, tcs := range tcs {
		t.Run(fmt.Sprintf("solve(%d, %v)", tcs.k, tcs.arr), func(t *testing.T) {
			got := maxMin(tcs.k, tcs.arr)
			if tcs.want != got {
				t.Fatalf("fail: want: %d, got %d", tcs.want, got)
			}
		})
	}
}
