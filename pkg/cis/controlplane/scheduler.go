package controlplane

import (
	"fmt"

	. "github.com/onsi/ginkgo"

	"github.com/mesosphere/kubernetes-security-benchmark/pkg/framework"
	. "github.com/mesosphere/kubernetes-security-benchmark/pkg/matcher"
)

const schedulerProcessName = "kube-scheduler"

func Scheduler(index, subIndex int) {
	f := framework.New(schedulerProcessName)

	It(fmt.Sprintf("[%d.%d.1] Ensure that the --profiling argument is set to false", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagWithValue("--profiling", "false"))
	})
}
