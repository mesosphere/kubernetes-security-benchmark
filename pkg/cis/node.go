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

package cis

import (
	. "github.com/onsi/ginkgo"

	"github.com/mesosphere/kubernetes-security-benchmark/pkg/cis/node"
	"github.com/mesosphere/kubernetes-security-benchmark/pkg/framework"
)

func describeNode(missingProcFunc framework.MissingProcessHandlerFunc) {
	CISDescribe("[2] Node", func() {
		Context("[2.1] Kubelet", func() {
			node.Kubelet(2, 1, missingProcFunc)
		})
	})
}
