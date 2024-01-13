package workers

import "context"

type headerKey string
type headerValue string
type bodyValue any

const (
	HeaderA headerKey = "A"
)

type Worker[T any] interface {
	AddMiddleware(m middleware[T])
	OnShutdown(cb func())
	Run(Handler[T])
	RunWithContext(ctx context.Context, h Handler[T], done func())
}

type Handler[T any] func(...any) (HandlerResult[T], error)

type HandlerResult[T any] struct {
	Headers map[headerKey]headerValue
	Body    T
}

type middleware[T any] func(Handler[T]) Handler[T]

func HandlerResultEmpty[T any]() HandlerResult[T] {
	var bodyValue T
	return HandlerResult[T]{
		Headers: make(map[headerKey]headerValue),
		Body:    bodyValue,
	}
}
