// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.

package queue_using_two_stacks

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func solve(in [][]int) []int {
	q := []int{}
	out := []int{}
	for _, i := range in {
		switch i[0] {
		case 1: // enqueue
			q = append(q, i[1])
		case 2: // dequeue
			q = q[1:]
		case 3: // print
			out = append(out, q[0])
		}
	}
	return out
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	writer := bufio.NewWriterSize(os.Stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int(qTemp)

	input := make([][]int, q)
	for i := 0; i < q; i++ {
		inTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
		in := []int{}
		for _, s := range inTemp {
			v, err := strconv.Atoi(s)
			checkError(err)
			in = append(in, v)
		}
		input[i] = in

	}

	result := solve(input)

	for _, r := range result {
		fmt.Fprintf(writer, "%d\n", r)
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
