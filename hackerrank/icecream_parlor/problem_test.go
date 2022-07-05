// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package icecream_parlor

import (
	"fmt"
	"reflect"
	"testing"
)

func TestProblem(t *testing.T) {
	tcs := []struct {
		m     int32
		input []int32
		want  []int32
	}{
		{
			4,
			[]int32{1, 4, 5, 3, 2},
			[]int32{1, 4},
		},
		{
			4,
			[]int32{2, 2, 4, 3},
			[]int32{1, 2},
		},
	}

	for idx, tcs := range tcs {
		t.Run(fmt.Sprintf("solve(%v)", idx), func(t *testing.T) {
			got := icecreamParlor(tcs.m, tcs.input)
			if !reflect.DeepEqual(tcs.want, got) {
				t.Fatalf("fail: want: %v, got %v", tcs.want, got)
			}
		})
	}
}
