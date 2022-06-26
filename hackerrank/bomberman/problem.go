// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package bomberman

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'bomberMan' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. STRING_ARRAY grid
 */

func bomberMan(n int32, grid []string) []string {
	if n < 2 {
		return grid
	}
	// After the first state, there will always be a cycle of 4 states:
	// 1 - inital bombs
	// 2 - all bombs (inital + "leftovers planted")
	// 3 - state after one explosion
	// 4 - all bombs (all execpt "leftovers+radius planted")
	// 5 = 1 - state after two epxlosions
	n = n % 4
	switch n {
	case 0, 2: // all bombs
		out := make([]string, len(grid))
		row := strings.Repeat("O", len(grid[0]))
		for i := range out {
			out[i] = row
		}
		return out
	case 1: // 2 explosions
		return explode(explode(grid))
	case 3: // 1 explosion
		return explode(grid)

	}
	panic("invalid n")
}

func explode(grid []string) []string {
	// create a grid full of boms
	newGrid := make([][]byte, len(grid))
	for i := range newGrid {
		newGrid[i] = bytes.Repeat([]byte("O"), len(grid[0]))
	}
	// explode any bombs in previous grid
	for r := range newGrid {
		for c := range newGrid[r] {
			if grid[r][c] == 'O' {
				newGrid[r][c] = '.'
				if r > 0 { // North
					newGrid[r-1][c] = '.'
				}
				if c+1 < len(newGrid[r]) { // East
					newGrid[r][c+1] = '.'
				}
				if r+1 < len(newGrid) { // South
					newGrid[r+1][c] = '.'
				}
				if c > 0 { // West
					newGrid[r][c-1] = '.'
				}
			}
		}
	}
	out := make([]string, len(newGrid))
	for i := range newGrid {
		out[i] = string(newGrid[i])
	}
	return out
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	rTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	r := int32(rTemp)

	cTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	_ = int32(cTemp)

	nTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)
	n := int32(nTemp)

	var grid []string

	for i := 0; i < int(r); i++ {
		gridItem := readLine(reader)
		grid = append(grid, gridItem)
	}

	result := bomberMan(n, grid)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
