// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package permuting_two_arrays

import (
	"fmt"
	"testing"
)

func TestProblem(t *testing.T) {
	tcs := []struct {
		k    int
		a, b []int
		want string
	}{
		{
			10,
			[]int{2, 1, 3},
			[]int{7, 8, 9},
			"YES",
		},
		{
			5,
			[]int{1, 2, 2, 1},
			[]int{3, 3, 3, 4},
			"NO",
		},
	}

	for i, tcs := range tcs {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := twoArrays(tcs.k, tcs.a, tcs.b)
			if got != tcs.want {
				t.Fatalf("fail: want: %s, got: %s", tcs.want, got)
			}
		})
	}
}
