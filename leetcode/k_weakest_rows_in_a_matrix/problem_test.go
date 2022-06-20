// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package k_weakest_rows_in_a_matrix

import (
	"fmt"
	"reflect"
	"testing"
)

func TestProblem(t *testing.T) {
	tcs := []struct {
		mat  [][]int
		k    int
		want []int
	}{
		{
			[][]int{
				{1, 1, 0, 0, 0},
				{1, 1, 1, 1, 0},
				{1, 0, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{1, 1, 1, 1, 1},
			},
			3,
			[]int{2, 0, 3},
		},
		{
			[][]int{
				{1, 0, 0, 0},
				{1, 1, 1, 1},
				{1, 0, 0, 0},
				{1, 0, 0, 0},
			},
			2,
			[]int{0, 2},
		},
		{
			[][]int{
				{1, 1, 0},
				{1, 1, 0},
				{1, 1, 1},
				{1, 1, 1},
				{0, 0, 0},
				{1, 1, 1},
				{1, 0, 0},
			},
			6,
			[]int{4, 6, 0, 1, 2, 3},
		},
	}
	for i, tcs := range tcs {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got := kWeakestRows(tcs.mat, tcs.k)
			if !reflect.DeepEqual(tcs.want, got) {
				t.Fatalf("fail: want %v, got %v", tcs.want, got)
			}
		})
	}
}
