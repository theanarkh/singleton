package singleton

import (
	"sync"
	"sync/atomic"
)

type F[T any] func() (*T, error)

type singleton[T any] struct {
	factory  F[T]
	instance atomic.Pointer[T]
	mutex    sync.Mutex
}

type ISingleton[T any] interface {
	Get() (*T, error)
}

func New[T any](factory F[T]) ISingleton[T] {
	return &singleton[T]{
		factory: factory,
	}
}

func (s *singleton[T]) Get() (*T, error) {
	if s.instance.Load() != nil {
		return s.instance.Load(), nil
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.instance.Load() != nil {
		return s.instance.Load(), nil
	}
	result, err := s.factory()
	if err != nil {
		return nil, err
	}
	s.instance.Store(result)
	return result, nil
}
