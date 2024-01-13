package workers

import (
	"context"

	"github.com/seadiaz/go-flow/pkg/helpers"
)

var _ Worker = &simpleWorker[struct{}]{}

func NewSimpleWorker[T any](in <-chan T) *simpleWorker[T] {
	return &simpleWorker[T]{
		baseWorker: newBaseWorker(),
		in:         in,
	}
}

type simpleWorker[T any] struct {
	baseWorker
	in <-chan T
}

func (w *simpleWorker[T]) AddMiddleware(m middleware) {
	w.middlewares = append(w.middlewares, m)
}

func (w *simpleWorker[T]) Run(h Handler) {
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

func (w *simpleWorker[T]) RunWithContext(ctx context.Context, h Handler, done func()) {
	helpers.LogInfo("not yet implemented")
	done()
}
