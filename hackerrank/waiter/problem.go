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

var primes = generatePrimes(10000)

func generatePrimes(n int) []int {
	notPrime := make([]bool, n+1)
	notPrime[0], notPrime[1] = true, true
	for i := 2; i*i <= n; i++ {
		if !notPrime[i] { // if prime
			for j := i * 2; j <= n; j += i {
				notPrime[j] = true
			}
		}
	}
	var primes []int
	for i, nP := range notPrime {
		if !nP {
			primes = append(primes, i)
		}
	}
	return primes
}

func waiter(plates []int32, q int32) []int32 {
	answers := []int32{}
	for i := int32(0); i < q; i++ {
		a, b := []int32{}, []int32{}
		v := primes[len(primes)-1]
		fmt.Println(v)
		for j := len(plates) - 1; j >= 0; j-- {
			p := plates[j]
			if int(p)%primes[i] == 0 {
				b = append(b, p)
			} else {
				a = append(a, p)
			}
		}
		// move b into answers
		for i := len(b) - 1; i >= 0; i-- {
			answers = append(answers, b[i])
		}
		plates = a
	}
	for i := len(plates) - 1; i >= 0; i-- {
		answers = append(answers, plates[i])
	}
	return answers
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	qTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	q := int32(qTemp)

	numberTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var number []int32

	for i := 0; i < int(n); i++ {
		numberItemTemp, err := strconv.ParseInt(numberTemp[i], 10, 64)
		checkError(err)
		numberItem := int32(numberItemTemp)
		number = append(number, numberItem)
	}

	result := waiter(number, q)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

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
