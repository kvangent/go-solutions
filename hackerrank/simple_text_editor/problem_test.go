// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package simple_text_editor

import (
	"fmt"
	"reflect"
	"testing"
)

func TestProblem(t *testing.T) {
	tcs := []struct {
		input [][]string
		want  []string
	}{
		{
			[][]string{
				{"1", "abc"},
				{"3", "3"},
				{"2", "3"},
				{"1", "xy"},
				{"3", "2"},
				{"4"},
				{"4"},
				{"3", "1"},
			},
			[]string{"c", "y", "a"},
		},
		{
			[][]string{
				{"1", "abcde"},
				{"1", "fg"},
				{"3", "6"},
				{"2", "5"},
				{"4"},
				{"3", "7"},
				{"4"},
				{"3", "4"},
			},
			[]string{"f", "g", "d"},
		},
	}

	for _, tcs := range tcs {
		t.Run(fmt.Sprintf("solve(%v)", tcs.input), func(t *testing.T) {
			got := solve(tcs.input)
			if !reflect.DeepEqual(tcs.want, got) {
				t.Fatalf("fail: want: %v, got %v", tcs.want, got)
			}
		})
	}
}
