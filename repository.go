package main

import (
	"sync"
)

type Observable[T any] struct {
	value     T
	mu        sync.RWMutex
	observers []chan T
}

func NewObservable[T any](value T) *Observable[T] {
	return &Observable[T]{
		value: value,
	}
}

func (o *Observable[T]) Subscribe() <-chan T {
	o.mu.Lock()
	defer o.mu.Unlock()
	ch := make(chan T)
	o.observers = append(o.observers, ch)
	return ch
}

func (o *Observable[T]) Unsubscribe(ch <-chan T) {
	o.mu.Lock()
	defer o.mu.Unlock()
	for i, observer := range o.observers {
		if observer == ch {
			o.observers = append(o.observers[:i], o.observers[i+1:]...)
			close(observer)
			break
		}
	}
}

func (o *Observable[T]) SetValue(value T) {
	o.mu.Lock()
	o.value = value
	for _, observer := range o.observers {
		observer <- value
	}
	o.mu.Unlock()
}

func (o *Observable[T]) GetValue() T {
	o.mu.RLock()
	defer o.mu.RUnlock()
	return o.value
}

type Repository struct {
	keyCombinations    Observable[[]KeyCombination]
	currentApplication Observable[string]
	pressedKeys        Observable[[]KeyCode]
}

func NewRepository() *Repository {
	return &Repository{
		keyCombinations:    *NewObservable([]KeyCombination{}),
		currentApplication: *NewObservable(""),
		pressedKeys:        *NewObservable([]KeyCode{}),
	}
}

var (
	instance *Repository
	once     sync.Once
)

func GetRepository() *Repository {
	once.Do(func() {
		instance = NewRepository()
	})
	return instance
}
