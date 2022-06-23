// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package plus_minus

import (
	"fmt"
	"reflect"
	"testing"
)

func TestProblem(t *testing.T) {
	tcs := []struct {
		input []int32
		want  []float64
	}{
		{
			[]int32{1, 1, 0, -1, -1},
			[]float64{0.4, 0.4, 0.2},
		},
	}

	for _, tcs := range tcs {
		t.Run(fmt.Sprintf("solve(%d)", tcs.input), func(t *testing.T) {
			got := plusMinus(tcs.input)
			if !reflect.DeepEqual(tcs.want, got) {
				t.Fatalf("fail: want %v, got %v", tcs.want, got)
			}
		})
	}
}
