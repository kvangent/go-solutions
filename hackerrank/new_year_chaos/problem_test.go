// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package new_year_chaos

import (
	"fmt"
	"testing"
)

func TestProblem(t *testing.T) {
	tcs := []struct {
		input []int
		want  string
	}{
		{
			[]int{2, 1, 5, 3, 4},
			"3",
		},
		{
			[]int{2, 5, 1, 3, 4},
			"Too chaotic",
		},
	}

	for i, tcs := range tcs {
		t.Run(fmt.Sprintf("solve %d", i), func(t *testing.T) {
			got := minimumBribes(tcs.input)
			if tcs.want != got {
				t.Fatalf("fail: want: %v, got %v", tcs.want, got)
			}
		})
	}
}
