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
	"fmt"

	. "github.com/onsi/ginkgo"

	"github.com/mesosphere/kubernetes-security-benchmark/pkg/framework"
)

func GeneralSecurityPrimitives(index, subIndex int, missingProcessFunc framework.MissingProcessHandlerFunc) {
	PIt(fmt.Sprintf("[%d.%d.1] Ensure that the cluster-admin role is only used where required", index, subIndex))

	PIt(fmt.Sprintf("[%d.%d.2] Create Pod Security Policies for your cluster", index, subIndex))

	PIt(fmt.Sprintf("[%d.%d.3] Create administrative boundaries between resources using namespaces", index, subIndex))

	PIt(fmt.Sprintf("[%d.%d.4] Create network segmentation using Network Policies", index, subIndex))

	PIt(fmt.Sprintf("[%d.%d.5] Ensure that the seccomp profile is set to docker/default in your pod definitions", index, subIndex))

	PIt(fmt.Sprintf("[%d.%d.6] Apply Security Context to Your Pods and Containers", index, subIndex))

	PIt(fmt.Sprintf("[%d.%d.7] Configure Image Provenance using ImagePolicyWebhook admission controller", index, subIndex))

	PIt(fmt.Sprintf("[%d.%d.8] Configure Network policies as appropriate", index, subIndex))

	PIt(fmt.Sprintf("[%d.%d.9] Place compensating controls in the form of PSP and RBAC for privileged containers usage", index, subIndex))
}
