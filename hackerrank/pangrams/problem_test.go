// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package pangrams

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
			"We promptly judged antique ivory buckles for the next prize",
			"pangram",
		},
		{
			"We promptly judged antique ivory buckles for the prize",
			"not pangram",
		},
	}

	for _, tcs := range tcs {
		t.Run(fmt.Sprintf("solve(%s)", tcs.input), func(t *testing.T) {
			got := pangrams(tcs.input)
			if tcs.want != got {
				t.Fatalf("fail: want: %s, got %s", tcs.want, got)
			}
		})
	}
}
