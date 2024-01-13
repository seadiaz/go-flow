package workers

func newBaseWorker[T any]() baseWorker[T] {
	return baseWorker[T]{
		middlewares: make([]middleware[T], 0),
		shutdownFn: func() {
			// empty function
		},
	}
}

type baseWorker[T any] struct {
	middlewares []middleware[T]
	shutdownFn  func()
}

func (w *baseWorker[T]) combine(h Handler[T]) Handler[T] {
	if len(w.middlewares) == 0 {
		return h
	}

	result := w.middlewares[0](h)
	for _, m := range w.middlewares[1:] {
		result = m(result)
	}
	return result
}

func (w *baseWorker[T]) OnShutdown(cb func()) {
	w.shutdownFn = cb
}
