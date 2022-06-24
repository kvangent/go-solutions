// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.
package permuting_two_arrays

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'twoArrays' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY A
 *  3. INTEGER_ARRAY B
 */
func twoArrays(k int, A []int, B []int) string {
    // Write your code here
    sort.Sort(sort.IntSlice(A))
    sort.Sort(sort.Reverse(sort.IntSlice(B)))
    for i := range A {
        if A[i] + B[i] < k {
            return "NO"
        }
    }
    return "YES"
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    q := int32(qTemp)

    for qItr := 0; qItr < int(q); qItr++ {
        firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

        nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
        checkError(err)
        n := int32(nTemp)

        kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
        checkError(err)
        k := int(kTemp)

        ATemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

        var A []int

        for i := 0; i < int(n); i++ {
            AItem, err := strconv.ParseInt(ATemp[i], 10, 64)
            checkError(err)
            A = append(A, int(AItem))
        }

        BTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

        var B []int

        for i := 0; i < int(n); i++ {
            BItem, err := strconv.ParseInt(BTemp[i], 10, 64)
            checkError(err)
            B = append(B, int(BItem))
        }

        result := twoArrays(k, A, B)

        fmt.Fprintf(writer, "%s\n", result)
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
