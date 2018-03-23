package benchmark_test

import (
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/shirou/gopsutil/process"
)

func controllerManager() {
	var controllerManagerProcess *process.Process

	processes, _ := process.Processes()
	for _, p := range processes {
		exe, _ := p.Exe()
		if filepath.Base(exe) == "kube-controller-manager" {
			controllerManagerProcess = p
			break
		}
	}

	BeforeEach(func() {
		if controllerManagerProcess == nil {
			Fail("Controller manager is not running")
		}
	})

	PIt("1.3.1 Ensure that the --terminated-pod-gc-threshold argument is set as appropriate")

	It("1.3.2 Ensure that the --profiling argument is set to false", func() {
		Expect(controllerManagerProcess.CmdlineSlice()).To(HaveFlagWithValue("--profiling", "false"))
	})

	It("1.3.3 Ensure that the --use-service-account-credentials argument is set to true", func() {
		Expect(controllerManagerProcess.CmdlineSlice()).To(HaveFlagWithValue("--use-service-account-credentials", "true"))
	})

	It("1.3.4 Ensure that the --service-account-private-key-file argument is set as appropriate", func() {
		Expect(controllerManagerProcess.CmdlineSlice()).To(HaveFlagWithAnyValue("--service-account-private-key-file"))
	})

	It("1.3.5 Ensure that the --root-ca-file argument is set as appropriate", func() {
		Expect(controllerManagerProcess.CmdlineSlice()).To(HaveFlagWithAnyValue("--root-ca-file"))
	})

	PIt("1.3.6 Apply Security Context to Your Pods and Containers")

	It("1.3.7 Ensure that the RotateKubeletServerCertificate argument is set to true", func() {
		Expect(controllerManagerProcess.CmdlineSlice()).To(HaveFlagThatContainsValue("--feature-gates", "RotateKubeletServerCertificate=true"))
	})
}
