package workers

import (
	"context"
)

var _ Worker[struct{}] = &simpleWorker[struct{}]{}

func NewSimpleWorker[T any](in <-chan T) *simpleWorker[T] {
	return &simpleWorker[T]{
		baseWorker: newBaseWorker[T](),
		in:         in,
	}
}

type simpleWorker[T any] struct {
	baseWorker[T]
	in <-chan T
}

func (w *simpleWorker[T]) AddMiddleware(m middleware[T]) {
	w.middlewares = append(w.middlewares, m)
}

func (w *simpleWorker[T]) Run(h Handler[T]) {
	finalHandler := w.combine(h)
	for {
		msg, ok := <-w.in
		if !ok {
			w.in = nil
			continue
		}
		finalHandler(msg)
	}
}

func (w *simpleWorker[T]) RunWithContext(ctx context.Context, h Handler[T], done func()) {
	finalHandler := w.combine(h)
	for {
		select {
		case <-ctx.Done():
			w.shutdownFn()
			done()
			return
		case msg, ok := <-w.in:
			if !ok {
				w.in = nil
				continue
			}
			finalHandler(msg)
		}
	}
}
