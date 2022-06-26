// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package counter_game

import (
	"fmt"
	"testing"
)

func TestProblem(t *testing.T) {
	tcs := []struct {
		n    int64
		want string
	}{
		{
			132,
			"Louise",
		},
		{
			6,
			"Richard",
		},
		{
			1,
			"Richard",
		},
	}

	for _, tcs := range tcs {
		t.Run(fmt.Sprintf("solve(%d)", tcs.n), func(t *testing.T) {
			got := counterGame(tcs.n)
			if tcs.want != got {
				t.Fatalf("fail: want: %v, got: %v", tcs.want, got)
			}
		})
	}
}
