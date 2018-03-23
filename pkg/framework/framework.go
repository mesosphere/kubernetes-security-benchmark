package framework

import (
	"fmt"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	"github.com/shirou/gopsutil/process"
)

type Framework struct {
	ProcessName string
	Process     *process.Process
}

func New(processName string) *Framework {
	var proc *process.Process
	processes, _ := process.Processes()
	for _, p := range processes {
		exe, _ := p.Exe()
		if filepath.Base(exe) == processName {
			proc = p
			break
		}
	}

	f := &Framework{
		ProcessName: processName,
		Process:     proc,
	}

	BeforeEach(f.beforeEach)

	return f
}

func (f *Framework) beforeEach() {
	if f.Process == nil {
		Fail(fmt.Sprintf("%s is not running", f.ProcessName))
	}
}
