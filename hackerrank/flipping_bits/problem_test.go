// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package flipping_bits

import (
	"fmt"
	"testing"
)

func TestProblem(t *testing.T) {
	tcs := []struct {
		input uint32
		want  uint32
	}{
		{
			2147483647,
			2147483648,
		},
		{
			1,
			4294967294,
		},
		{
			0, 
			4294967295,
		},
	}

	for _, tcs := range tcs {
		t.Run(fmt.Sprintf("solve(%d)", tcs.input), func(t *testing.T) {
			got := flippingBits(tcs.input)
			if tcs.want != got {
				t.Fatalf("fail: want: %d, got %d", tcs.want, got)
			}
		})
	}
}
