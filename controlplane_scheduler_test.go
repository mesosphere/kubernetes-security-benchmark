package benchmark_test

import (
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/shirou/gopsutil/process"
)

func scheduler() {
	var schedulerProcess *process.Process

	processes, _ := process.Processes()
	for _, p := range processes {
		exe, _ := p.Exe()
		if filepath.Base(exe) == "kube-scheduler" {
			schedulerProcess = p
			break
		}
	}

	BeforeEach(func() {
		if schedulerProcess == nil {
			Fail("Scheduler is not running")
		}
	})

	It("1.2.1 Ensure that the --profiling argument is set to false", func() {
		Expect(schedulerProcess.CmdlineSlice()).To(HaveFlagWithValue("--profiling", "false"))
	})
}
