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

package controlplane

import (
	. "github.com/onsi/ginkgo"

	"github.com/mesosphere/kubernetes-security-benchmark/pkg/framework"
)

func PodSecurityPolicies(missingProcessFunc framework.MissingProcessHandlerFunc) {
	PIt("[1.7.1] Do not admit privileged containers [Not Scored]")

	PIt("[1.7.2] Do not admit containers wishing to share the host process ID namespace [Scored]")

	PIt("[1.7.3] Do not admit containers wishing to share the host IPC namespace [Scored]")

	PIt("[1.7.4] Do not admit containers wishing to share the host network namespace [Scored]")

	PIt("[1.7.5] Do not admit containers with allowPrivilegeEscalation [Scored]")

	PIt("[1.7.6] Do not admit root containers [Not Scored]")

	PIt("[1.7.7] Do not admit containers with dangerous capabilities [Not Scored]")
}
