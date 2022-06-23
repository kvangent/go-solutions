// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package time_conversion

import (
	"fmt"
	"testing"
)

func TestProblem(t *testing.T) {
	tcs := []struct {
		input string
		want  string
	}{
		{
			"07:05:45PM",
			"19:05:45",
		},
	}

	for _, tcs := range tcs {
		t.Run(fmt.Sprintf("solve(%s)", tcs.input), func(t *testing.T) {
			got := timeConversion(tcs.input)
			if tcs.want != got {
				t.Fatalf("fail: want: %s, got: %s", tcs.want, got)
			}
		})
	}
}
