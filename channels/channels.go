package channels

import (
	"fmt"
	"sync"
)

type Element struct {
	Type string
}

var e1 = Element{
	Type: "Type A",
}
var e2 = Element{
	Type: "Type B",
}

func channelSender(channel <-chan Element, channelName string) {
	for e := range channel {
		fmt.Printf("Received element %v in %s", e, channelName)
		// time.Sleep(4 * time.Second)
		for i := 0; i < 1000000000; i++ {
			// nothing but still takes time
		}
	}
}

func Channels() {

	array := []Element{e1, e2, e1, e2, e1, e2, e1, e2}

	channelA := make(chan Element)
	channelB := make(chan Element)
	defer close(channelA)
	defer close(channelB)

	// Constructing the w-group with the # elements to pass
	wg := &sync.WaitGroup{}
	wg.Add(len(array))

	go channelSender(channelA, "Channel A")
	go channelSender(channelB, "Channel B")

	for _, e := range array {
		go func(e Element) {
			if e.Type == "Type A" {
				fmt.Println("Before sending A")
				channelA <- e
				fmt.Println("After sending A")
			}
			if e.Type == "Type B" {
				fmt.Println("Before sending B")
				channelB <- e
				fmt.Println("After sending B")
			}
			wg.Done()
		}(e)
	}

	wg.Wait()
}
