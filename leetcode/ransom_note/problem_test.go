// Copyright Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package ransom_note

import (
	"fmt"
	"testing"
)

func TestProblem(t *testing.T) {
	tcs := []struct {
		ransom   string
		magazine string
		want     bool
	}{
		{"a", "b", false},
		{"aa", "ab", false},
		{"aa", "aab", true},
	}

	for _, tcs := range tcs {
		t.Run(fmt.Sprintf("solve(%s, %s)", tcs.ransom, tcs.magazine), func(t *testing.T) {
			got := canConstruct(tcs.ransom, tcs.magazine)
			if tcs.want != got {
				t.Fatalf("fail: want %t, got %t", tcs.want, got)
			}
		})
	}
}
