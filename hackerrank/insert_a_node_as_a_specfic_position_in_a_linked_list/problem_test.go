// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package insert_a_node_as_a_specfic_position_in_a_linked_list

import (
	"fmt"
	"reflect"
	"testing"
)

func TestProblem(t *testing.T) {
	tcs := []struct {
		input []int32
		data, pos int32
		want  []int32
	}{
		{
			[]int32{16, 13, 7},
			1,
			2,
			[]int32{16, 13, 1, 7},
		},
	}

	for idx, tcs := range tcs {
		t.Run(fmt.Sprintf("solve(%v)", idx), func(t *testing.T) {
			in := SinglyLinkedList{}
			for _, v := range tcs.input {
				in.insertNodeIntoSinglyLinkedList(v)
			}
			got := insertNodeAtPosition(in.head, tcs.data, tcs.pos).toSlice()
			if !reflect.DeepEqual(tcs.want, got) {
				t.Fatalf("fail: want: %v, got %v", tcs.want, got)
			}
		})
	}
}
