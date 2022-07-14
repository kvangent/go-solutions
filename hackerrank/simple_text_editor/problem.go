// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package simple_text_editor

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'anagram' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

func solve(ops [][]string) []string {
	s := ""
	out := []string{}
	hist := []string{}
	for _, op := range ops {
		switch op[0] {
		case "1":
			hist = append(hist, s)
			s += op[1]
		case "2":
			hist = append(hist, s)
			k, err := strconv.Atoi(op[1])
			checkError(err)
			s = s[:len(s)-k]
		case "3":
			k, err := strconv.Atoi(op[1])
			checkError(err)
			out = append(out, string(s[k-1]))
		case "4":
			idx := len(hist) - 1
			s, hist = hist[idx], hist[:idx]
		}
	}
	return out
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	writer := bufio.NewWriterSize(os.Stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	var lines [][]string
	for qItr := 0; qItr < int(q); qItr++ {
		s := strings.Split(readLine(reader), " ")
		lines = append(lines, s)
	}

	for _, v := range solve(lines) {
		fmt.Println(v)
	}

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
