// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package caesar_cipher

import (
	"fmt"
	"testing"
)

func TestProblem(t *testing.T) {
	tcs := []struct {
		n string
		k int32
		want  string
	}{
		{
			"middle-Outz",
			2,
			"okffng-Qwvb",
		},
		{
			"abcdefghijklmnopqrstuvwxyz",
			2,
			"cdefghijklmnopqrstuvwxyzab",
		},
	}

	for _, tcs := range tcs {
		t.Run(fmt.Sprintf("solve(%s, %d)", tcs.n, tcs.k), func(t *testing.T) {
			got := caesarCipher(tcs.n, tcs.k)
			if tcs.want != got {
				t.Fatalf("fail: want: %s, got %s", tcs.want, got)
			}
		})
	}
}
