// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package bomberman

import (
	"fmt"
	"reflect"
	"testing"
)

func TestProblem(t *testing.T) {
	tcs := []struct {
		n    int32
		grid []string
		want []string
	}{
		{
			3,
			[]string{
				".......",
				"...O...",
				"....O..",
				".......",
				"OO.....",
				"OO.....",
			},
			[]string{
				"OOO.OOO",
				"OO...OO",
				"OOO...O",
				"..OO.OO",
				"...OOOO",
				"...OOOO",
			},
		},
		{
			5,
			[]string{
				".......",
				"...O.O.",
				"....O..",
				"..O....",
				"OO...OO",
				"OO.O...",
			},
			[]string{
				".......",
				"...O.O.",
				"...OO..",
				"..OOOO.",
				"OOOOOOO",
				"OOOOOOO",
			},
		},
	}

	for _, tcs := range tcs {
		t.Run(fmt.Sprintf("solve(%d)", tcs.n), func(t *testing.T) {
			got := bomberMan(tcs.n, tcs.grid)
			if !reflect.DeepEqual(got, tcs.want) {
				t.Fatalf("fail: want: %v, got %v", tcs.want, got)
			}
		})
	}
}
