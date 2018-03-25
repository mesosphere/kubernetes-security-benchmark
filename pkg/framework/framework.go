// Copyright © 2018 Jimmi Dyson <jimmidyson@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
