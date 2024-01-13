package workers

import (
	"context"

	"github.com/seadiaz/go-flow/pkg/helpers"
)

var _ Worker = &simpleWorker{}

func NewSimpleWorker(in <-chan any) *simpleWorker {
	return &simpleWorker{
		baseWorker: newBaseWorker(),
		in:         in,
	}
}

type simpleWorker struct {
	baseWorker
	in <-chan any
}

func (w *simpleWorker) AddMiddleware(m middleware) {
	w.middlewares = append(w.middlewares, m)
}

func (w *simpleWorker) Run(h Handler) {
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

func (w *simpleWorker) RunWithContext(ctx context.Context, h Handler, done func()) {
	helpers.LogInfo("not yet implemented")
	done()
}
