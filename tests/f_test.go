package tests

import (
	"strings"
	"testing"
)

func TestStringsIndex(t *testing.T) {
	f := func(s, substr string, nExpected int) {
		t.Helper()

		n := strings.Index(s, substr)
		if n != nExpected {
			t.Fatalf("unexpected n; got %d; want %d", n, nExpected)
		}
	}

	// first char match
	f("foobar", "foo", 0)

	// middle char match
	f("foobar", "bar", 3)

	// mismatch
	f("foobar", "baz", -1)
}

func TestSimpleErrFunc_Failure(t *testing.T) {
	f := func(input string) {
		t.Helper()

		_, err := SimpleErrFunc(input)
		if err == nil {
			t.Fatalf("expecting non-nil error")
		}
	}

	f("broken_input_1")
	f("broken_input_2")
}

func TestSimpleErrFunc_Success(t *testing.T) {
	f := func(input, resultExpected string) {
		t.Helper()

		result, err := SimpleErrFunc(input)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		if result != resultExpected {
			t.Fatalf("unexpected result; got %q; want %q", result, resultExpected)
		}
	}

	f("input_1", "result_1")
	f("input_2", "result_2")
}

func TestSimpleErrFuncWithSubtests(t *testing.T) {
	f := func(t *testing.T, input, outputExpected string) {
		t.Helper()

		output, err := SimpleErrFunc(input)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		if output != outputExpected {
			t.Fatalf("unexpected output; got %q; want %q", output, outputExpected)
		}
	}

	t.Run("first_subtest", func(t *testing.T) {
		f(t, "foo", "foo")
	})

	t.Run("second_subtest", func(t *testing.T) {
		f(t, "baz", "baz")
	})
}
