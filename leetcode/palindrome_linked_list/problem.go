// Copyright 2021 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package palindrome_linked_list

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	// edge case: len of 1
	if head.Next == nil {
		return true
	}
	// find the middle of the list
	mid := head
	for tail := head; tail.Next != nil && tail.Next.Next != nil; {
		mid = mid.Next
		tail = tail.Next.Next
	}

	// reverse the back half
	rev := reverse(mid.Next)
	mid.Next = nil

	for head != nil && rev != nil {
		if head.Val != rev.Val {
			return false
		}
		head = head.Next
		rev = rev.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	var prev, next *ListNode
	for head != nil {
		next = head.Next
		head.Next = prev
		prev, head = head, next
	}
	return prev
}
