package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func dynamicArray(n int32, queries [][]int32) []int32 {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic: %+v\n", r)
		}
	}()
	var results []int32
	var lastAnswer int32 = 0
	var arr [][]int32
	for i := 0; i < int(n); i++ {
		arr = append(arr, []int32{})
	}

	for _, query := range queries {
		x, y := query[1], query[2]
		arrN := (x ^ lastAnswer) % n
		if query[0] == 1 {
			arr[arrN] = append(arr[arrN], y)
		} else {
			lastAnswer = arr[arrN][int(y)%len(arr[arrN])]
			results = append(results, lastAnswer)
			fmt.Println(lastAnswer)
		}
	}

	return results
}

func main() {

	file, _ := os.Open("dyn_test_1")
	reader := bufio.NewReaderSize(file, 16*1024*1024)
	// reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	qTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	q := int32(qTemp)

	var queries [][]int32
	for i := 0; i < int(q); i++ {
		queriesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}
		if len(queriesRow) != 3 {
			panic("Bad input")
		}
		queries = append(queries, queriesRow)
	}

	result := dynamicArray(n, queries)

	cDir, _ := os.Getwd()
	stdout, err := os.CreateTemp(cDir, "dyn_result_*")
	checkError(err)
	defer stdout.Close()
	writer := bufio.NewWriterSize(stdout, 16*1024*1024)
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
