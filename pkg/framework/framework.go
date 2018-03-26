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

	"github.com/shirou/gopsutil/process"
)

type MissingProcessHandlerFunc func(string, ...int)

type Framework struct {
	ProcessName        string
	Process            *process.Process
	missingProcessFunc MissingProcessHandlerFunc
}

func New(processName string, missingProcessFunc MissingProcessHandlerFunc) *Framework {
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
		ProcessName:        processName,
		Process:            proc,
		missingProcessFunc: missingProcessFunc,
	}

	return f
}

func (f *Framework) BeforeEach() {
	if f.Process == nil {
		f.missingProcessFunc(fmt.Sprintf("%s is not running", f.ProcessName))
	}
}
