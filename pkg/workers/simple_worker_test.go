package workers_test

import (
	"errors"

	"github.com/seadiaz/go-flow/pkg/workers"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Simple Worker", func() {
	var sut workers.Worker[any]
	var handler workers.Handler[any]
	var in chan any
	var shutdownFn func()
	BeforeEach(func() {
		in = make(chan any)
		sut = workers.NewSimpleWorker(in)
		shutdownFn = func() {
			// empty function
		}
		sut.AddMiddleware(workers.NewLoggingMiddleware[any]())
		sut.AddMiddleware(workers.NewPerformanceMiddleware[any]())
		sut.OnShutdown(shutdownFn)
	})

	When("a new valid message arrives", func() {
		It("should execute the handler", func() {
			handler = func(params ...any) (workers.HandlerResult[any], error) {
				Expect(params).Should(HaveLen(1))
				return workers.HandlerResultEmpty[any](), nil
			}
			go sut.Run(handler)
			in <- struct{}{}
		})
	})

	When("a new invalid message arrives", func() {
		It("should execute the handler", func() {
			handler = func(params ...any) (workers.HandlerResult[any], error) {
				Expect(params).Should(HaveLen(1))
				return workers.HandlerResultEmpty[any](), errors.New("d03ff5bb-dab2-4c2e-9ffe-92216146cab5")
			}
			go sut.Run(handler)
			in <- struct{}{}
		})
	})

})
