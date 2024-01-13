package workers

import "fmt"

func NewPerformanceMiddleware() middleware {
	return performanceMiddleware
}

func performanceMiddleware(next Handler) Handler {
	return func(args ...any) (any, error) {
		// do something before

		result, err := next(args...)
		if err != nil {
			return nil, fmt.Errorf("error executing child: %w", err)
		}

		// do something after
		return result, nil
	}
}
