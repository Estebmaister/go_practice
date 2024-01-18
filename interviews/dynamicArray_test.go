package interviews

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

const inputTest = "dyn_test_%d"
const resultTest = "dyn_result_%s"

// Pending comparison with result files
func TestDynamicArray(t *testing.T) {

	tests := []struct {
		name        string
		number      int
		writeToFile bool
	}{
		{name: "2&5 expecting 7&3", number: 0},
		{name: "100&100 expecting more", number: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Open the file with test data
			file, err := os.Open(fmt.Sprintf(inputTest, tt.number))
			checkError(err)
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

			if tt.writeToFile {
				writeResultToFile(result)
			}

		})
	}
}

func writeResultToFile(result []int32) error {
	cDir, _ := os.Getwd()
	stdout, err := os.CreateTemp(cDir, fmt.Sprintf(resultTest, "*"))
	if err != nil {
		return err
	}
	defer stdout.Close()
	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	for idx, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)
		if idx != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")
	writer.Flush()

	return nil
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
