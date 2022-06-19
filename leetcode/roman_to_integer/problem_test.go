// Copyright 2021 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package roman_to_integer

import "testing"

func TestProblem(t *testing.T) {
	tcs := []struct {
		input string
		want  int
	}{
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for _, tcs := range tcs {
		t.Run(tcs.input, func(t *testing.T) {
			got := romanToInt(tcs.input)
			if tcs.want != got {
				t.Fatalf("failed(%s): want %d, got %d", tcs.input, tcs.want, got)
			}
		})
	}
}
