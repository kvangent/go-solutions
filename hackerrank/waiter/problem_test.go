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
		q     int32
		input []int32
		want  []int32
	}{
		{
			3,
			[]int32{2, 3, 4, 5, 6, 7},
			[]int32{2, 4, 6, 3, 5, 7},
		},
		{
			1,
			[]int32{3, 4, 7, 6, 5},
			[]int32{4, 6, 3, 7, 5},
		},
		{
			3,
			[]int32{3, 4, 7, 6, 5},
			[]int32{4, 6, 3, 5, 7},
		},
	}

	for idx, tcs := range tcs {
		t.Run(fmt.Sprintf("solve(%v)", idx), func(t *testing.T) {
			got := waiter(tcs.input, tcs.q)
			if !reflect.DeepEqual(tcs.want, got) {
				t.Fatalf("fail: want: %v, got %v", tcs.want, got)
			}
		})
	}
}
