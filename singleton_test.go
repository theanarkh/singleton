package singleton

import (
	"testing"

	"golang.org/x/sync/errgroup"
)

type Dummy struct {
	// Avoid the empty structure return the same address every time
	Ptr *string
}

func factory() (*Dummy, error) {
	ptr := "dummy"
	return &Dummy{Ptr: &ptr}, nil
}

func TestNew(t *testing.T) {
	singleton := New(factory)
	instance, err := singleton.Get()
	if err != nil {
		t.Fatal(err)
	}
	if instance == nil {
		t.Fatal("singleton should not be nil")
	}
}

func TestCallTwice(t *testing.T) {
	singleton := New(factory)
	len := 2
	instances := make([]*Dummy, len)
	for i := 0; i < len; i++ {
		instance, err := singleton.Get()
		if err != nil {
			t.Fatal(err)
		}
		if instance == nil {
			t.Fatal("singleton should not be nil")
		}
		instances = append(instances, instance)
	}
	var last *Dummy
	for i := 0; i < len; i++ {
		if last == nil {
			last = instances[i]
		} else if instances[i] != last {
			t.Fatal("instance should be the same")
		}
	}
}

func TestConcurrent(t *testing.T) {
	singleton := New(factory)
	len := 100
	instances := make([]*Dummy, len)
	var eg errgroup.Group
	for i := 0; i < len; i++ {
		i := i
		eg.Go(func() error {
			instance, err := singleton.Get()
			if err != nil {
				return err
			}
			instances[i] = instance
			return nil
		})
	}
	err := eg.Wait()
	if err != nil {
		t.Fatal(err)
	}
	var last *Dummy
	for i := 0; i < len; i++ {
		if instances[i] == nil {
			t.Fatal("singleton should not be nil")
		}
		if last == nil {
			last = instances[i]
		} else if instances[i] != last {
			t.Fatal("instance should be the same")
		}
	}
}
