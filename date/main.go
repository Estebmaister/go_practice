// to modified the marshing parsing check the following url
// https://blog.charmes.net/post/json-dates-go/
package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type myDate struct {
	DeletedAt time.Time  `json:"deleted_at,omitempty" example:"2019-11-19T12:58:42.28119Z"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" example:"2009-11-10T23:00:00Z"`
}

func main() {

	// Managing dates in golang default format
	dateNow := time.Now().UTC()
	fmt.Println(dateNow)
	dateStruct := myDate{dateNow, &dateNow}
	fmt.Printf("%+v \n", dateStruct)
	dateJSON, _ := json.Marshal(dateStruct)
	fmt.Printf("%s \n\n", dateJSON)

	// Parsing from Unix format
	if intDate, err := strconv.ParseInt("1641221067", 10, 64); err != nil {
		fmt.Println("Error while parsing unix date :", err)
	} else {
		unixDate := time.Unix(intDate, 0)
		fmt.Println("extractUnix:", unixDate)
	}

	// Other formats
	parsedDate, err := time.Parse(time.RFC3339, "2022-01-03T14:44:27Z") //"1641221067"||2022-01-03T14:44:27Z
	if err != nil {
		fmt.Println("Error while parsing date :", err)
	}
	// Display the time as RFC3339 (same as JSON) (not UTC)
	fmt.Println("RFC3339 dft:", parsedDate.Format(time.RFC3339))
	fmt.Println("RFC3339Nano:", parsedDate.Format(time.RFC3339Nano))
	fmt.Println("Unix:       ", parsedDate.Unix())                        //timestamp
	fmt.Println("Unix str:   ", strconv.FormatInt(parsedDate.Unix(), 10)) //timestamp
	fmt.Println("hour/minute:", parsedDate.Format("3:04PM"))
	fmt.Println("ANSIC:      ", parsedDate.Format(time.ANSIC))
	fmt.Println("UnixDate:   ", parsedDate.Format(time.UnixDate))
	fmt.Println("RFC1123:    ", parsedDate.Format(time.RFC1123))
	fmt.Println("RubyDate:   ", parsedDate.Format(time.RubyDate))
}
