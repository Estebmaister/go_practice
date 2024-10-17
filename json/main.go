package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"runtime/debug"
	"strings"

	"github.com/Jeffail/gabs/v2"
)

var jsonPath = "./ex.json"

const (
	GREETINGS = "hi"
	JSON      = "json"
	EXIT      = "exit"
)

// Simple shell interface, json reader
func main() {

	handler := slog.NewJSONHandler(os.Stdout, nil)
	buildInfo, _ := debug.ReadBuildInfo()
	logger := slog.New(handler)
	logger = logger.With(
		slog.Group("program_info",
			slog.Int("pid", os.Getpid()),
			slog.String("go_version", buildInfo.GoVersion),
		),
	)
	slog.SetDefault(logger)
	slog.Info("Simple shell interface, json reader")

	fmt.Println("-----------------------------------")
	fmt.Println("Commands:", GREETINGS, JSON, EXIT)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		input, _ := reader.ReadString('\n')
		// convert CRLF to LF
		// input = strings.Replace(input, "\n", "", -1)
		input = strings.Trim(input, "\n ")

		if strings.EqualFold(GREETINGS, input) {
			fmt.Println("hello, Yourself")
		}

		if strings.Compare(JSON, input) == 0 {
			readJson(jsonPath)
		}

		if input == EXIT {
			break
		}
	}

}

func readJson(path string) {

	jsonFile, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error %s reading alias config from path %s", err, path)
		return
	}

	fmt.Println("Parsing example json elastic response with gabs")
	jsonDoc, err := gabs.ParseJSON(jsonFile)
	if err != nil {
		fmt.Printf("Error %s parsing alias config as gabs container", err)
		return
	}

	jsonTotal, ok := jsonDoc.Path("total").Data().(float64)
	if !ok {
		fmt.Println("No path 'total' found or incorrect type assertion")
		return
	}

	fmt.Println("Extracting outlets IDs values from hits")
	outIDs := make([]string, 0)
	for idx := 0; idx < int(jsonTotal); idx++ {
		jsonVal := jsonDoc.Path(fmt.Sprintf("hits.%d.Id", idx)).Data()
		jsonData, ok := jsonVal.(string)
		if ok {
			if strings.Contains(jsonData, "outlet") {
				outIDs = append(outIDs, jsonData)
			}
		} else {
			fmt.Printf("value: %+v, type: %T \n", jsonVal, jsonVal)
		}
	}

	fmt.Printf(
		"total_hits_reported: %f, outlets: %d, id_values: %s \n",
		jsonTotal,
		len(outIDs),
		strings.ReplaceAll(fmt.Sprintf("%+q", outIDs), `" "`, `","`),
	)
}
