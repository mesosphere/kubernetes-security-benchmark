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

func GeneralSecurityPrimitives(missingProcessFunc framework.MissingProcessHandlerFunc) {
	PIt("[1.6.1] Ensure that the cluster-admin role is only used where required [Not Scored]")

	PIt("[1.6.2] Create Pod Security Policies for your cluster [Not Scored]")

	PIt("[1.6.3] Create administrative boundaries between resources using namespaces [Not Scored]")

	PIt("[1.6.4] Create network segmentation using Network Policies [Not Scored]")

	PIt("[1.6.5] Ensure that the seccomp profile is set to docker/default in your pod definitions [Not Scored]")

	PIt("[1.6.6] Apply Security Context to Your Pods and Containers [Not Scored]")

	PIt("[1.6.7] Configure Image Provenance using ImagePolicyWebhook admission controller [Not Scored]")

	PIt("[1.6.8] Configure Network policies as appropriate [Not Scored]")

	PIt("[1.6.9] Place compensating controls in the form of PSP and RBAC for privileged containers usage [Not Scored]")
}
