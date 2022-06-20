// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package middle_of_the_linked_list

import (
	"fmt"
	"reflect"
	"testing"
)

func NewListNode(n []int) *ListNode {
	head := &ListNode{Val: n[0]}
	tail := head
	for _, v := range n[1:] {
		tail.Next = &ListNode{Val: v}
		tail = tail.Next
	}
	return head
}

func (l *ListNode) toInt() (out []int) {
	for cur := l; cur != nil; cur = cur.Next {
		out = append(out, cur.Val)
	}
	return
}

func TestProblem(t *testing.T) {
	tcs := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5}},
		{[]int{1, 2, 3, 4, 5, 6}, []int{4, 5, 6}},
	}
	for _, tcs := range tcs {
		t.Run(fmt.Sprintf("solve(%d)", tcs.input), func(t *testing.T) {
			got := middleNode(NewListNode(tcs.input))
			if !reflect.DeepEqual(tcs.want, got.toInt()) {
				t.Fatalf("fail: want %v, got %v", tcs.want, got)
			}
		})
	}
}
