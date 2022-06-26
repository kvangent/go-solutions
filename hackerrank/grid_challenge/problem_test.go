// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package grid_challenge

import (
	"fmt"
	"testing"
)

func TestProblem(t *testing.T) {
	tcs := []struct {
		grid []string
		want string
	}{
		{
			[]string{
				"abc",
				"ade",
				"efg",
			},
			"YES",
		},
		{
			[]string{
				"mpxz",
				"abcd",
				"wlmf",
			},
			"NO",
		},
		{
			[]string{
				"eabcd",
				"fghij",
				"olkmn",
				"trpqs",
				"xywuv",
			},
			"YES",
		},
	}

	for _, tcs := range tcs {
		t.Run(fmt.Sprintf("solve(%v)", tcs.grid), func(t *testing.T) {
			got := gridChallenge(tcs.grid)
			if tcs.want != got {
				t.Fatalf("fail: want: %s, got %s", tcs.want, got)
			}
		})
	}
}
