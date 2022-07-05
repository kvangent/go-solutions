// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package sherlock_and_valid_string

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'isValid' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func isValid(s string) string {
	// find frequenies
	freq := make(map[rune]int)
	for _, r := range s {
		freq[r]++
	}
	// find unique frequencies
	unique := make(map[int]int)
	for _, ct := range freq {
		unique[ct]++
	}
	if len(unique) > 2 {
		return "NO"
	}
	if len(unique) == 2 {
		// we can only remove one character, so only valid if one frequency is off by 1
		uF := make([]int, 0)  // freq
		uCt := make([]int, 0) // freq of freq
		for k, v := range unique {
			uF = append(uF, k)
			uCt = append(uCt, v)
		}
		// sort smallest to biggest
		if uCt[0] > uCt[1] {
			uF[0], uF[1] = uF[1], uF[0]
			uCt[0], uCt[1] = uCt[1], uCt[0]
		}
		// only allow 1 execption if it's off by 1 or if it's the only character
		if uCt[0] > 1 || (uF[0] != 1 && AbsDiff(uF[0], uF[1]) > 1) {
			return "NO"
		}
	}
	return "YES"
}

func AbsDiff(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := isValid(s)

	fmt.Fprintf(writer, "%s\n", result)

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
