package workers

import "context"

type Worker interface {
	AddMiddleware(m middleware)
	Run(Handler)
	RunWithContext(ctx context.Context, h Handler, done func())
}

type Handler func(...any) (any, error)

type middleware func(Handler) Handler
