package workers

import (
	"fmt"

	"github.com/seadiaz/go-flow/pkg/helpers"
)

func NewLoggingMiddleware() middleware {
	return loggingMiddleware
}

func loggingMiddleware(next Handler) Handler {
	return func(args ...any) (any, error) {
		helpers.LogInfo("hanlder initiatized")

		result, err := next(args...)
		if err != nil {
			helpers.LogInfo("hanlder failed: %s", err.Error())
			return nil, fmt.Errorf("error executing child: %w", err)
		}

		helpers.LogInfo("hanlder finished")
		return result, nil
	}
}
