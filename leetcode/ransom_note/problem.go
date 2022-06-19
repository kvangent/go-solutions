// Copyright 2021 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package ransom_note

func canConstruct(ransomNote string, magazine string) bool {
	ltr := [26]int{}
	for _, l := range magazine {
		ltr[int(l-'a')]++
	}
	for _, l := range ransomNote {
		c := int(l - 'a')
		ltr[c]--
		if ltr[c] < 0 {
			return false
		}
	}
	return true
}
