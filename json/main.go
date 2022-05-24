package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/Jeffail/gabs/v2"
)

var jsonPath = "./ex.json"

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
		}

		if strings.Compare("json", text) == 0 {
			readJson(jsonPath)
		}
	}

}

func readJson(path string) {

	jsonFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Error %s reading alias config from path %s", err, path)
		return
	}

	jsonDoc, err := gabs.ParseJSON(jsonFile)
	if err != nil {
		fmt.Printf("Error %s parsing alias config as gabs container", err)
		return
	}

	jsonTotal, ok := jsonDoc.Path("total").Data().(float64)
	if !ok {
		fmt.Printf("ok: %v,total: %v \n", ok, jsonTotal)
		return
	}

	idsArray := make([]string, 0)
	for index := 0; index < int(jsonTotal); index++ {
		jsonVal := jsonDoc.Path(fmt.Sprintf("hits.%d.Id", index)).Data()
		jsonData, ok := jsonVal.(string)
		if ok {
			if strings.Contains(jsonData, "outlet") {
				idsArray = append(idsArray, jsonData)
			}
		} else {
			fmt.Printf("value: %+v, type: %T \n", jsonVal, jsonVal)
		}
	}

	fmt.Printf("len: %d, values: %s \n", len(idsArray), strings.ReplaceAll(fmt.Sprintf("%+q", idsArray), `" "`, `","`))
}
