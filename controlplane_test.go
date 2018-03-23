package benchmark_test

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("1 Control plane", func() {
	Context("1.1 API Server", apiServer)
	Context("1.2 Scheduler", scheduler)
	Context("1.3 Controller Manager", controllerManager)
})
