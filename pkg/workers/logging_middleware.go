package workers

import (
	"fmt"

	"github.com/seadiaz/go-flow/pkg/helpers"
)

func NewLoggingMiddleware[T any]() middleware[T] {
	return loggingMiddleware
}

func loggingMiddleware[T any](next Handler[T]) Handler[T] {
	return func(args ...any) (HandlerResult[T], error) {
		helpers.LogInfo("hanlder initiatized")

		result, err := next(args...)
		if err != nil {
			helpers.LogInfo("hanlder failed: %s", err.Error())
			return HandlerResult[T]{}, fmt.Errorf("error executing child: %w", err)
		}

		helpers.LogInfo("hanlder finished")
		return result, nil
	}
}
