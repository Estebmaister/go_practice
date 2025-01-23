package server

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("New server with opts", func(t *testing.T) {
		want := &Server{}
		want.MaxConn = 111
		want.Host = "local"
		want.Port = 999
		got := New(WithTLS(false),
			WithHost("local"),
			WithPort(999),
			WithMaxConn(111))
		got.Run()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("New() = %#v, want %v", got, want)
		}
	})
}
