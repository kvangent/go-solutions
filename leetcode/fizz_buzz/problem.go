// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package fizz_buzz

import "strconv"

var results [10000]string

func fizzBuzz(n int) []string {
	for i := n; i > 0 && results[i-1] == ""; i-- {
		three := i%3 == 0
		five := i%5 == 0
		switch {
		case three && five:
			results[i-1] = "FizzBuzz"
		case three:
			results[i-1] = "Fizz"
		case five:
			results[i-1] = "Buzz"
		default:
			results[i-1] = strconv.Itoa(i)
		}
	}
	return results[:n]
}
