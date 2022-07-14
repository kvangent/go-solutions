// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package anagram

import (
	"fmt"
	"testing"
)

func TestProblem(t *testing.T) {
	tcs := []struct {
		input string
		want  int
	}{
		{
			"aaabbb",
			3,
		},
		{
			"ab",
			1,
		},
		{
			"abc",
			-1,
		},
		{
			"mnop",
			2,
		},
		{
			"xyyx",
			0,
		},
		{
			"xaxbbbxx",
			1,
		},
	}

	for _, tcs := range tcs {
		t.Run(fmt.Sprintf("solve(%v)", tcs.input), func(t *testing.T) {
			got := anagram(tcs.input)
			if tcs.want != got {
				t.Fatalf("fail: want: %v, got %v", tcs.want, got)
			}
		})
	}
}
