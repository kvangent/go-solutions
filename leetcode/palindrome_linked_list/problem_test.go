// Copyright Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package palindrome_linked_list

import (
	"fmt"
	"testing"
)

func TestProblem(t *testing.T) {
	tcs := []struct {
		input []int
		want  bool
	}{
		{[]int{1, 2, 2, 1}, true},
		{[]int{1, 2}, false},
		{[]int{1, 0, 1}, true},
	}

	for i, tcs := range tcs {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			head := &ListNode{Val: tcs.input[0]}
			tail := head
			for _, v := range tcs.input[1:] {
				tail.Next = &ListNode{Val: v}
				tail = tail.Next
			}
			got := isPalindrome(head)
			if tcs.want != got {
				t.Fatalf("fail: want %t, got %t", tcs.want, got)
			}
		})
	}
}
