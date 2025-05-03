package iterators

import "testing"

func BenchmarkGetTotalActiveWithIter(b *testing.B) {
	users := []User{
		{Username: "Alice", Active: true, Points: 120},
		{Username: "Bob", Active: false, Points: 85},
		{Username: "Charlie", Active: true, Points: 200},
		{Username: "Diana", Active: false, Points: 50},
		{Username: "Eve", Active: true, Points: 150},
	}

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		GetTotalActiveWithIter(users)
	}
}

func BenchmarkGetTotalActiveWithRange(b *testing.B) {
	users := []User{
		{Username: "Alice", Active: true, Points: 120},
		{Username: "Bob", Active: false, Points: 85},
		{Username: "Charlie", Active: true, Points: 200},
		{Username: "Diana", Active: false, Points: 50},
		{Username: "Eve", Active: true, Points: 150},
	}

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		GetTotalActiveWithRange(users)
	}
}

func BenchmarkGetTotalActiveWithSlices(b *testing.B) {
	users := []User{
		{Username: "Alice", Active: true, Points: 120},
		{Username: "Bob", Active: false, Points: 85},
		{Username: "Charlie", Active: true, Points: 200},
		{Username: "Diana", Active: false, Points: 50},
		{Username: "Eve", Active: true, Points: 150},
	}

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		GetTotalActiveWithSlices(users)
	}
}
