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
	. "github.com/mesosphere/kubernetes-security-benchmark/pkg/matcher"
)

const controllerManagerProcessName = "kube-controller-manager"

func ControllerManager(index, subIndex int) {
	f := framework.New(controllerManagerProcessName)

	PIt(fmt.Sprintf("[%d.%d.1] Ensure that the --terminated-pod-gc-threshold argument is set as appropriate", index, subIndex))

	It(fmt.Sprintf("[%d.%d.2] Ensure that the --profiling argument is set to false", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagWithValue("--profiling", "false"))
	})

	It(fmt.Sprintf("[%d.%d.3] Ensure that the --use-service-account-credentials argument is set to true", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagWithValue("--use-service-account-credentials", "true"))
	})

	It(fmt.Sprintf("[%d.%d.4] Ensure that the --service-account-private-key-file argument is set as appropriate", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagWithAnyValue("--service-account-private-key-file"))
	})

	It(fmt.Sprintf("[%d.%d.5] Ensure that the --root-ca-file argument is set as appropriate", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagWithAnyValue("--root-ca-file"))
	})

	PIt(fmt.Sprintf("[%d.%d.6] Apply Security Context to Your Pods and Containers", index, subIndex))

	It(fmt.Sprintf("[%d.%d.7] Ensure that the RotateKubeletServerCertificate argument is set to true", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagThatContainsValue("--feature-gates", "RotateKubeletServerCertificate=true"))
	})
}
