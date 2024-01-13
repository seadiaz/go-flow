package workers

func newBaseWorker() baseWorker {
	return baseWorker{
		middlewares: make([]middleware, 0),
	}
}

type baseWorker struct {
	middlewares []middleware
}

func (w *baseWorker) combine(h Handler) Handler {
	if len(w.middlewares) == 0 {
		return h
	}

	result := w.middlewares[0](h)
	for _, m := range w.middlewares[1:] {
		result = m(result)
	}
	return result
}
