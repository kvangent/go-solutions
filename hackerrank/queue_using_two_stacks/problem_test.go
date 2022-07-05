// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package queue_using_two_stacks

import (
	"fmt"
	"reflect"
	"testing"
)

func TestProblem(t *testing.T) {
	tcs := []struct {
		input [][]int
		want  []int
	}{
		{
			[][]int{
				{1, 42},
				{2},
				{1, 14},
				{3},
				{1, 28},
				{3},
				{1, 60},
				{1, 70},
				{2},
				{2},
			},
			[]int{14, 14},
		},
	}

	for idx, tcs := range tcs {
		t.Run(fmt.Sprintf("solve(%v)", idx), func(t *testing.T) {
			got := solve(tcs.input)
			if !reflect.DeepEqual(tcs.want, got) {
				t.Fatalf("fail: want: %v, got %v", tcs.want, got)
			}
		})
	}
}
