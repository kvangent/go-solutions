// Copyright 2021 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package roman_to_integer

func romanToInt(s string) int {
	var v int
	var p rune
	for _, r := range s {
		switch r {
		case 'M':
			if p == 'C' {
				v += 800
			} else {
				v += 1000
			}
		case 'D':
			if p == 'C' {
				v += 300
			} else {
				v += 500
			}
		case 'C':
			if p == 'X' {
				v += 80
			} else {
				v += 100
			}
		case 'L':
			if p == 'X' {
				v += 30
			} else {
				v += 50
			}
		case 'X':
			if p == 'I' {
				v += 8
			} else {
				v += 10
			}
		case 'V':
			if p == 'I' {
				v += 3
			} else {
				v += 5
			}
		case 'I':
			v += 1
		}
		p = r
	}
	return v
}
