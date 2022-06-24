// Copyright 2022 Kurtis Van Gent
//
// Use of this source code is governed by an MIT-style license that can be found
// in the LICENSE file or at https://opensource.org/licenses/MIT.
package subarray_division_1

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'birthday' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY s
 *  2. INTEGER d
 *  3. INTEGER m
 */

func birthday(s []int, d int, m int) int {
    sum := 0
    for _, v := range s[:m-1] {
        sum += v
    }
    ct := 0
    for i := m-1; i < len(s); i++ {
        sum += s[i]
        if sum == d {
            ct++
        }
        sum -= s[i-(m-1)]
    }
    return ct
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    sTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var s []int

    for i := 0; i < int(n); i++ {
        sItem, err := strconv.ParseInt(sTemp[i], 10, 64)
        checkError(err)
        s = append(s, int(sItem))
    }

    firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    d, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
    checkError(err)

    m, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
    checkError(err)

    result := birthday(s, int(d), int(m))

    fmt.Fprintf(writer, "%d\n", result)

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
