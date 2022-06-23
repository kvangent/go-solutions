// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package diagonal_difference

import (
	"fmt"
	"testing"
)

func TestProblem(t *testing.T) {
	tcs := []struct {
		input [][]int
		want  int
	}{
		{
			[][]int{
				{11, 2, 4},
				{4, 5, 6},
				{10, 8, -12},
			},
			15,
		},
	}

	for _, tcs := range tcs {
		t.Run(fmt.Sprintf("solve(%d)", tcs.input), func(t *testing.T) {
			got := diagonalDifference(tcs.input)
			if tcs.want != got {
				t.Fatalf("fail: want: %d, got %d", tcs.want, got)
			}
		})
	}
}
