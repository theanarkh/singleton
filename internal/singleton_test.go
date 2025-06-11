package singleton

import (
	"testing"
)

type Dummy struct {
	// Avoid the empty structure return the same address every time
	Ptr *string
}

func factory() (*Dummy, error) {
	ptr := "dummy"
	return &Dummy{Ptr: &ptr}, nil
}

func BenchmarkSingletonGet(b *testing.B) {
	singleton := New(factory)
	for n := 0; n < b.N; n++ {
		singleton.Get()
	}
}
