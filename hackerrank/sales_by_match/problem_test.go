// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package sales_by_match

import (
	"fmt"
	"testing"
)

func TestProblem(t *testing.T) {
	tcs := []struct {
		input []int
		want  int
	}{
		{
			[]int{10,20,20,10,10,30,50,10,20},
			3,
		},
	}

	for _, tcs := range tcs {
		t.Run(fmt.Sprintf("solve(%v)", tcs.input), func(t *testing.T) {
			got := sockMerchant(len(tcs.input), tcs.input)
			if tcs.want != got {
				t.Fatalf("fail: want: %d, got: %d", tcs.want, got)
			}
		})
	}
}
