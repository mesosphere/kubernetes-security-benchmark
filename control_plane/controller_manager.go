package control_plane

import (
	"fmt"

	. "github.com/onsi/ginkgo"

	"github.com/mesosphere/kubernetes-security-benchmark/framework"
	. "github.com/mesosphere/kubernetes-security-benchmark/matcher"
)

const controllerManagerProcessName = "kube-controller-manager"

func ControllerManager(index, subIndex int) {
	f := framework.New(controllerManagerProcessName)

	PIt(fmt.Sprintf("[%d.%d.1] Ensure that the --terminated-pod-gc-threshold argument is set as appropriate", index, subIndex))

	It(fmt.Sprintf("[%d.%d.2] Ensure that the --profiling argument is set to false", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagWithValue("--profiling", "false"))
	})

	It(fmt.Sprintf("[%d.%d.3] Ensure that the --use-service-account-credentials argument is set to true", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagWithValue("--use-service-account-credentials", "true"))
	})

	It(fmt.Sprintf("[%d.%d.4] Ensure that the --service-account-private-key-file argument is set as appropriate", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagWithAnyValue("--service-account-private-key-file"))
	})

	It(fmt.Sprintf("[%d.%d.5] Ensure that the --root-ca-file argument is set as appropriate", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagWithAnyValue("--root-ca-file"))
	})

	PIt(fmt.Sprintf("[%d.%d.6] Apply Security Context to Your Pods and Containers", index, subIndex))

	It(fmt.Sprintf("[%d.%d.7] Ensure that the RotateKubeletServerCertificate argument is set to true", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagThatContainsValue("--feature-gates", "RotateKubeletServerCertificate=true"))
	})
}
