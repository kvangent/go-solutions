// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package recursive_digit_sum

import (
	"fmt"
	"testing"
)

func TestProblem(t *testing.T) {
	tcs := []struct {
		n string
		k int
		want  int
	}{
		{
			"9875",
			4, 
			8,
		},
	}

	for _, tcs := range tcs {
		t.Run(fmt.Sprintf("solve(%s, %d)", tcs.n, tcs.k), func(t *testing.T) {
			got := superDigit(tcs.n, tcs.k)
			if tcs.want != got {
				t.Fatalf("fail: want: %d, got %d", tcs.want, got)
			}
		})
	}
}
