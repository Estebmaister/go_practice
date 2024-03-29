package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/estebmaister/go_practice/server/handlers"
)

const (
	Host = ""
	Port = 9090
)

// Run function starts a server in localhost:9090
func Run() {
	// http.HandleFunc("/test", func(w http.ResponseWriter, _ *http.Request) {
	// 	io.WriteString(w, "Hello from a testing func!\n")
	// })
	// http.ListenAndServe(host, nil)

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	hh := handlers.NewHello(l)
	dh := handlers.NewData(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/data", dh)

	host := fmt.Sprintf("%s:%v", Host, Port)
	s := &http.Server{
		Addr:         host,
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 2)
	signal.Notify(sigChan,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)
	//os.Notify

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := s.Shutdown(tc)
	if err != nil {
		l.Panic("Problem during shutdown", err)
	}
}
