package benchmark_test

import (
	. "github.com/onsi/ginkgo"

	controlplane "github.com/mesosphere/kubernetes-security-benchmark/control_plane"
)

func CISDescribe(text string, body func()) bool {
	return Describe("[CIS] "+text, body)
}

var _ = CISDescribe("[1] Control plane", func() {
	Context("[1.1] API Server", func() {
		controlplane.APIServer(1, 1)
	})
	Context("[1.2] Scheduler", func() {
		controlplane.Scheduler(1, 2)
	})
	Context("[1.3] Controller Manager", func() {
		controlplane.ControllerManager(1, 3)
	})
})
