// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package k_weakest_rows_in_a_matrix

import "sort"

type row struct {
	idx int
	sCt int
}

func newRow(idx int, s []int) row {
	// binary search to find first 0
	// 1 1 1 0
	l := 0
	for h := len(s); l < h; {
		m := (l + h) / 2
		if s[m] == 0 {
			h = m
		} else {
			l = m + 1
		}
	}
	return row{idx, l}
}

func kWeakestRows(mat [][]int, k int) []int {
	var rows []row
	for i, r := range mat {
		rows = append(rows, newRow(i, r))
	}
	sort.SliceStable(rows, func(i, j int) bool {
		return rows[i].sCt < rows[j].sCt
	})
	out := make([]int, k)
	for i := 0; i < k; i++ {
		out[i] = rows[i].idx
	}
	return out[:]
}
