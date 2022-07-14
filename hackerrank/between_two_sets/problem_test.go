// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package between_two_sets

import (
	"fmt"
	"testing"
)

func TestProblem(t *testing.T) {
	tcs := []struct {
		a    []int32
		b    []int32
		want int32
	}{
		{
			[]int32{2, 4},
			[]int32{16, 32, 96},
			3,
		},
	}

	for idx, tcs := range tcs {
		t.Run(fmt.Sprintf("solve(%v)", idx), func(t *testing.T) {
			got := getTotalX(tcs.a, tcs.b)
			if tcs.want != got {
				t.Fatalf("fail: want: %v, got %v", tcs.want, got)
			}
		})
	}
}
