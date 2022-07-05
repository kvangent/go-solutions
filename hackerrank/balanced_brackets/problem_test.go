// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package balanced_brackets

import (
	"fmt"
	"reflect"
	"testing"
)

func TestProblem(t *testing.T) {
	tcs := []struct {
		input string
		want  string
	}{
		{
			"{[()]}",
			"YES",
		},
		{
			"{[(])}",
			"NO",
		},
		{
			"{{[[(())]]}}",
			"YES",
		},
		{
			"}][}}(}][))]",
			"NO",
		},
		{
			"[](){()}",
			"YES",
		},
		{
			"()",
			"YES",
		},
		{
			"({}([][]))[]()",
			"YES",
		},
		{
			"{)[](}]}]}))}(())(",
			"NO",
		},
		{
			"([[)",
			"NO",
		},
	}

	for idx, tcs := range tcs {
		t.Run(fmt.Sprintf("solve(%v)", idx), func(t *testing.T) {
			got := isBalanced(tcs.input)
			if !reflect.DeepEqual(tcs.want, got) {
				t.Fatalf("fail: want: %v, got %v", tcs.want, got)
			}
		})
	}
}
