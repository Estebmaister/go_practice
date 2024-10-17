package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'cavityMap' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts STRING_ARRAY grid as parameter.
 */

func cavityMap(grid []string) []string {
	n := len(grid)
	newGrid := make([]string, n)
	copy(newGrid, grid)
	if n < 3 {
		return newGrid
	}
	for i := 1; i < n-1; i++ {
		pervLine := grid[i-1]
		currentLine := grid[i]
		nextLine := grid[i+1]
		for j := 1; j < n-1; j++ {
			north, _ := strconv.ParseInt(string(pervLine[j]), 10, 64)
			south, _ := strconv.ParseInt(string(nextLine[j]), 10, 64)
			east, _ := strconv.ParseInt(string(currentLine[j-1]), 10, 64)
			west, _ := strconv.ParseInt(string(currentLine[j+1]), 10, 64)
			cavity, _ := strconv.ParseInt(string(currentLine[j]), 10, 64)
			if cavity > north && cavity > south && cavity > east && cavity > west {
				newGrid[i] = newGrid[i][:j] + "X" + newGrid[i][j+1:]
			}
		}
	}
	return newGrid
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var grid []string

	for i := 0; i < int(n); i++ {
		gridItem := readLine(reader)
		grid = append(grid, gridItem)
	}

	result := cavityMap(grid)

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
