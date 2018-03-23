package cis

import (
	. "github.com/onsi/ginkgo"

	"github.com/mesosphere/kubernetes-security-benchmark/pkg/cis/controlplane"
)

const CISVersion = "1.2.0"

func CISDescribe(text string, body func()) bool {
	return Describe("[CIS 1.2.0] "+text, body)
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
