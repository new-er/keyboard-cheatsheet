package observables

import "github.com/imkira/go-observer"

type Observable[T any] struct {
	prop observer.Property
}

func NewObservable[T any](value T) *Observable[T] {
	return &Observable[T]{observer.NewProperty(value)}
}

func (o *Observable[T]) Set(value T) {
	o.prop.Update(value)
}

func (o *Observable[T]) Get() T {
	return o.prop.Value().(T)
}

func (o *Observable[T]) Observe() *ObserverStream[T] {
	return NewObserverStream[T](o.prop.Observe())
}
