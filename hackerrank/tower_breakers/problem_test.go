// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package tower_breakers

import (
	"fmt"
	"testing"
)

func TestProblem(t *testing.T) {
	tcs := []struct {
		n, m int32
		want  int32
	}{
		{
			2, 2,
			2,
		},
		{
			1, 4,
			1,
		},

	}

	for _, tcs := range tcs {
		t.Run(fmt.Sprintf("solve(%d, %d)", tcs.n, tcs.m), func(t *testing.T) {
			got := towerBreakers(tcs.n, tcs.m)
			if tcs.want != got {
				t.Fatalf("fail: want: %d, got %d", tcs.want, got)
			}
		})
	}
}
