// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package drawing_book

import (
	"fmt"
	"testing"
)

func TestProblem(t *testing.T) {
	tcs := []struct {
		n, p int32
		want int32
	}{
		{
			5, 3,
			1,
		},
	}

	for _, tcs := range tcs {
		t.Run(fmt.Sprintf("solve(%d,%d)", tcs.n, tcs.p), func(t *testing.T) {
			got := pageCount(tcs.n, tcs.p)
			if tcs.want != got {
				t.Fatalf("failed: want %d, got %d", tcs.want, got)
			}

		})
	}
}
