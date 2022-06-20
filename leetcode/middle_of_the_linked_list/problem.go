// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package middle_of_the_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	mid, tail := head, head
	for tail.Next != nil && tail.Next.Next != nil {
		mid, tail = mid.Next, tail.Next.Next
	}
	if tail.Next != nil { // even length, get second middle node
		mid = mid.Next
	}
	return mid
}
