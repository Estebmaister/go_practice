// to modified the marshing parsing check the following url
// https://blog.charmes.net/post/json-dates-go/
package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type myDate struct {
	DeletedAt time.Time  `json:"deleted_at,omitempty" example:"2019-11-19T12:58:42.28119Z"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" example:"2009-11-10T23:00:00Z"`
}

func main() {

	// Managing dates in golang default format
	dateNow := time.Now().UTC()
	fmt.Println(dateNow, "\n")

	dateStruct := myDate{dateNow, &dateNow}
	fmt.Printf("%+v \n\n", dateStruct)

	dateJSON, _ := json.Marshal(dateStruct)
	fmt.Printf("%s \n\n", dateJSON)

	// Other formats
	now := time.Now()
	// Display the time as RFC3339 (same as JSON) (not UTC)
	fmt.Printf("%s\n", now.Format(time.RFC3339))
	// Display the timestamp
	fmt.Printf("%+v\n", now.Unix())
	// Display only the hour/minute
	fmt.Printf("%s\n", now.Format("3:04PM"))
}
