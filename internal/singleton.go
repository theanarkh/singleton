package singleton

import (
	"sync"
)

type F[T any] func() (*T, error)

type singleton[T any] struct {
	factory  F[T]
	instance *T
	mutex    sync.RWMutex
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
	s.mutex.RLock()
	if s.instance != nil {
		s.mutex.RUnlock()
		return s.instance, nil
	}
	s.mutex.RUnlock()
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.instance != nil {
		return s.instance, nil
	}
	result, err := s.factory()
	if err != nil {
		return nil, err
	}
	s.instance = result
	return result, nil
}
