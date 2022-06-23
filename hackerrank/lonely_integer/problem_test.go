// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package lonely_integer

import (
	"fmt"
	"testing"
)

func TestProblem(t *testing.T) {
	tcs := []struct {
		input   []int32
		want 	int32
	}{
		{
			[]int32{1, 2, 3, 4, 3, 2, 1},
			4,
		},
	}

	for _, tcs := range tcs {
		t.Run(fmt.Sprintf("solve(%d)", tcs.input), func(t *testing.T) {
			got := lonelyinteger(tcs.input)
			if tcs.want != got {
				t.Fatalf("fail: want: %d, got %d", tcs.want, got)
			}
		})
	}
}
