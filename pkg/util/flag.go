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

package util

import (
	"fmt"
	"strings"

	"github.com/shirou/gopsutil/process"
)

func FlagValueFromProcess(p *process.Process, flagName string) (interface{}, error) {
	cmdline, err := p.CmdlineSlice()
	if err != nil {
		return "", err
	}

	flagPrefix := fmt.Sprintf("--%s=", flagName)
	for _, f := range cmdline {
		if strings.HasPrefix(f, flagPrefix) {
			return f[len(flagPrefix)], nil
		}
	}

	return "", nil
}
