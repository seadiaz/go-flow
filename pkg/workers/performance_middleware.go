package workers

import (
	"fmt"
	"time"

	"github.com/seadiaz/go-flow/pkg/helpers"
)

const (
	HeaderPerformance headerKey = "performance"
)

func NewPerformanceMiddleware[T any]() middleware[T] {
	return performanceMiddleware
}

func performanceMiddleware[T any](next Handler[T]) Handler[T] {
	return func(args ...any) (HandlerResult[T], error) {
		initialTime := time.Now()

		result, err := next(args...)
		if err != nil {
			return HandlerResult[T]{}, fmt.Errorf("error executing child: %w", err)
		}

		duration := time.Since(initialTime)
		result.Headers[HeaderPerformance] = headerValue(duration.String())
		helpers.LogInfo("hanlder durations: %s", duration.String())
		return result, nil
	}
}
