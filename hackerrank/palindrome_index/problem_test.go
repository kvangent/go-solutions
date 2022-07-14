// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package palindrome_index

import (
	"fmt"
	"testing"
)

func TestProblem(t *testing.T) {
	tcs := []struct {
		input string
		want  int32
	}{
		{
			"aaab",
			3,
		},
		{
			"aab",
			2,
		},
		{
			"baa",
			0,
		},
		{
			"baaa",
			0,
		},
		{
			"aaa",
			-1,
		},
		{
			"aba",
			-1,
		},
		{
			"abba",
			-1,
		},

	}

	for _, tcs := range tcs {
		t.Run(fmt.Sprintf("solve(%v)", tcs.input), func(t *testing.T) {
			got := palindromeIndex(tcs.input)
			if tcs.want != got {
				t.Fatalf("fail: want: %v, got %v", tcs.want, got)
			}
		})
	}
}
