// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package sherlock_and_array

import (
	"fmt"
	"testing"
)

func TestProblem(t *testing.T) {
	tcs := []struct {
		arr []int32
		want  string
	}{
		{
			[]int32{1, 2, 3},
			"NO",
		},
		{
			[]int32{1, 2, 3, 3},
			"YES",
		},
	}

	for _, tcs := range tcs {
		t.Run(fmt.Sprintf("solve(%v)", tcs.arr), func(t *testing.T) {
			got := balancedSums(tcs.arr)
			if tcs.want != got {
				t.Fatalf("fail: want: %s, got %s", tcs.want, got)
			}
		})
	}
}
