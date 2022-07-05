// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package reverse_a_doubly_linked_list

import (
	"fmt"
	"reflect"
	"testing"
)

func TestProblem(t *testing.T) {
	tcs := []struct {
		input []int32
		want  []int32
	}{
		{
			[]int32{1, 2, 3, 4},
			[]int32{4, 3, 2, 1},
		},
	}

	for _, tcs := range tcs {
		t.Run(fmt.Sprintf("solve(%v)", tcs.input), func(t *testing.T) {
			in := DoublyLinkedList{}
			for _, v := range tcs.input {
				in.insertNodeIntoDoublyLinkedList(v)
			}
			got := reverse(in.head).toSlice()
			if !reflect.DeepEqual(tcs.want, got) {
				t.Fatalf("fail: want: %v, got %v", tcs.want, got)
			}
		})
	}
}
