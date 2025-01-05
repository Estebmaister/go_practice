package concurrency

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"time"
)

type Logger interface {
	Error(string, ...any)
	Info(string, ...any)
}

// The iolayer type represents an input/output layer with a semaphore and logger.
// @property sema - The `sema` property is a pointer to a `sema` struct. It is likely used for managing
// concurrency and synchronization in the I/O layer.
// @property {logger} logger - The `logger` property is a variable of type `logger`. It is used to log
// messages or events in the code.
type iolayer struct {
	sema   *sema
	logger Logger
	rang   *rand.Rand
}

func NewIO(size int, logger Logger) *iolayer {
	return &iolayer{
		rang:   rand.New(rand.NewSource(time.Now().UnixMilli())),
		logger: logger,
		sema:   NewSema(size),
	}
}

// The `LimitedProcess` function is a method of the `iolayer` struct. It takes a variadic parameter
// `data` of type `string`.
func (s *iolayer) LimitedProcess(data ...string) {
	//-----------begin preprocessing----------
	var buff = bytes.NewBuffer([]byte{})
	for _, d := range data {
		buff.Write([]byte(d))
		buff.Write([]byte("\n"))
	}
	//-----------end preprocessing------------

	s.sema.Enter()
	//-----------begin critical section----------
	s.logger.Info("begin uploading", "file", data[0])
	err := s.upload(buff, data[0])
	s.logger.Info("end uploading:", "file", data[0])
	//-----------end critical section------------
	s.sema.Exit()

	//-----------begin postprocessing----------
	if err != nil {
		s.logger.Error("exception while upload", err.Error())
	}
	//-----------end  postprocessing-----------
}

// The `Upload` function is a method of the `service` struct. It takes an `io.Reader` as an argument,
// which represents a source of data. This function is responsible for uploading the data from the
// `io.Reader` to some destination.
func (s *iolayer) upload(reader io.Reader, filename string) error {
	// writing the reader data to some destination
	ms := time.Millisecond * time.Duration(s.rang.Intn(5000))
	time.Sleep(ms)
	s.logger.Info(fmt.Sprintf("it took %s to upload the file: %s", ms, filename))
	return nil
}
