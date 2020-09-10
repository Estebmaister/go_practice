package channels

import (
	"fmt"
	"sync"
)

type Elemento struct {
	Tipo string
}

func Channels() {

	e1 := Elemento{
		Tipo: "Tipo A",
	}

	e2 := Elemento{
		Tipo: "Tipo B",
	}

	arreglo := []Elemento{e1, e2, e1, e2, e1, e2, e1, e2}

	channelA := make(chan Elemento)
	channelB := make(chan Elemento)

	wg := &sync.WaitGroup{}
	wg.Add(len(arreglo))

	go func(channel <-chan Elemento) {
		for e := range channel {
			fmt.Println("Recibido A", e)
			for i := 0; i < 1000000000; i++ {
			}
		}
	}(channelA)

	go func(channel <-chan Elemento) {
		for e := range channel {
			fmt.Println("Recibido B", e)
			// time.Sleep(4 * time.Second)
			for i := 0; i < 100000000; i++ {
			}
		}
	}(channelB)

	for _, e := range arreglo {
		go func(e Elemento) {
			if e.Tipo == "Tipo A" {
				fmt.Println("Antes de mandar A")
				channelA <- e
				fmt.Println("Después de mandar A")
			}
			if e.Tipo == "Tipo B" {
				fmt.Println("Antes de mandar B")
				channelB <- e
				fmt.Println("Después de mandar B")
			}
			wg.Done()
		}(e)
	}

	wg.Wait()
}
