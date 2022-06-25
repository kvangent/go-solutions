// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package dynamic_array

import (
	"fmt"
	"reflect"
	"testing"
)

func TestProblem(t *testing.T) {
	tcs := []struct {
		n       int32
		queries [][]int32
		want    []int32
	}{
		{
			2,
			[][]int32{
				{1, 0, 5},
				{1, 1, 7},
				{1, 0, 3},
				{2, 1, 0},
				{2, 1, 1},
			},
			[]int32{7, 3},
		},
	}

	for _, tcs := range tcs {
		t.Run(fmt.Sprintf("solve(%d, %v)", tcs.n, tcs.queries), func(t *testing.T) {
			got := dynamicArray(tcs.n, tcs.queries)
			if !reflect.DeepEqual(got, tcs.want) {
				t.Fatalf("fail: want: %v, got %v", tcs.want, got)
			}
		})
	}
}
