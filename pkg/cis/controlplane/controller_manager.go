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
	. "github.com/mesosphere/kubernetes-security-benchmark/pkg/ginkgo/matchers"
)

const controllerManagerProcessName = "kube-controller-manager"

func ControllerManager(missingProcessFunc framework.MissingProcessHandlerFunc) {
	f := framework.New(controllerManagerProcessName, missingProcessFunc)
	BeforeEach(f.BeforeEach)

	PIt("[1.3.1] Ensure that the --terminated-pod-gc-threshold argument is set as appropriate [Scored]")

	It("[1.3.2] Ensure that the --profiling argument is set to false [Scored]", func() {
		ExpectProcess(f).To(HaveFlagWithValue("--profiling", "false"))
	})

	It("[1.3.3] Ensure that the --use-service-account-credentials argument is set to true [Scored]", func() {
		ExpectProcess(f).To(HaveFlagWithValue("--use-service-account-credentials", "true"))
	})

	It("[1.3.4] Ensure that the --service-account-private-key-file argument is set as appropriate [Scored]", func() {
		ExpectProcess(f).To(HaveFlagWithAnyValue("--service-account-private-key-file"))
	})

	It("[1.3.5] Ensure that the --root-ca-file argument is set as appropriate [Scored]", func() {
		ExpectProcess(f).To(HaveFlagWithAnyValue("--root-ca-file"))
	})

	It("[1.3.6] Ensure that the RotateKubeletServerCertificate argument is set to true [Scored]", func() {
		ExpectProcess(f).To(HaveFlagThatDoesNotContainValue("--feature-gates", "RotateKubeletServerCertificate=false"))
	})

	It("[1.3.7] Ensure that the --address argument is set to 127.0.0.1 [Scored]", func() {
		ExpectProcess(f).To(NotHaveFlagOrHaveFlagWithValue("--address", "127.0.0.1"))
		ExpectProcess(f).To(NotHaveFlagOrHaveFlagWithValue("--bind-address", "127.0.0.1"))
	})
}
