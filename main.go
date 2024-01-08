package main

import (
	"github.com/estebmaister/go_practice/concurrency"
	// "github.com/estebmaister/go_practice/channels"
	// "github.com/estebmaister/go_practice/server"
)

func main() {
	print("Starting from main\n\n")
	// greeting()
	// nilDiffs()

	concurrency.PrimesPipeline(10)
	// channels.LogChannel()

	// Server functions
	// fmt.Printf(
	// "Server listening on http://localhost%s:%v\n",
	// server.Host, server.Port
	// )
	// server.Run()
}
