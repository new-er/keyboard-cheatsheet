package observables

import "github.com/imkira/go-observer"

type ObserverStream[T any] struct {
	stream observer.Stream
}

func NewObserverStream[T any](stream observer.Stream) *ObserverStream[T] {
	return &ObserverStream[T]{
		stream: stream,
	}
}
func (o *ObserverStream[T]) Changes() chan T {
	toReturn := make(chan T)
	go func() {
		for {
			toReturn <- o.stream.Value().(T)
		}
	}()
	return toReturn
}
