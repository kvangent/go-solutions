// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package sparse_arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestProblem(t *testing.T) {
	tcs := []struct {
		strings []string
		queries []string
		want    []int32
	}{
		{
			[]string{"aba", "baba", "aba", "xzxb"},
			[]string{"aba", "xzxb", "ab"},
			[]int32{2, 1, 0},
		},
		{
			[]string{"def", "de", "fgh"},
			[]string{"de", "lmn", "fgh"},
			[]int32{1, 0, 1},
		},
		{
			[]string{
				"abcde", "sdaklfj", "asdjf", "na", "basdn", "sdaklfj", "asdjf",
				"na", "asdjf", "na", "basdn", "sdaklfj", "asdjf",
			},
			[]string{"abcde", "sdaklfj", "asdjf", "na", "basdn"},
			[]int32{1, 3, 4, 3, 2},
		},
	}

	for _, tcs := range tcs {
		t.Run(fmt.Sprintf("solve(%s)", tcs.strings), func(t *testing.T) {
			got := matchingStrings(tcs.strings, tcs.queries)
			if !reflect.DeepEqual(tcs.want, got) {
				t.Fatalf("fail: want: %v, got: %v", tcs.want, got)
			}
		})
	}
}
