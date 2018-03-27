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

const apiServerProcessName = "kube-apiserver"

func APIServer(index, subIndex int, missingProcessFunc framework.MissingProcessHandlerFunc) {
	f := framework.New(apiServerProcessName, missingProcessFunc)
	BeforeEach(f.BeforeEach)

	It(fmt.Sprintf("[%d.%d.1] Ensure that the --anonymous-auth argument is set to false", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagWithValue("--anonymous-auth", "false"))
	})

	It(fmt.Sprintf("[%d.%d.2] Ensure that the --basic-auth-file argument is not set", index, subIndex), func() {
		ExpectProcess(f).To(NotHaveFlag("--basic-auth-file"))
	})

	It(fmt.Sprintf("[%d.%d.3] Ensure that the --insecure-allow-any-token argument is not set", index, subIndex), func() {
		ExpectProcess(f).To(NotHaveFlag("--insecure-allow-any-token"))
	})

	It(fmt.Sprintf("[%d.%d.4] Ensure that the --kubelet-https argument is set to true", index, subIndex), func() {
		ExpectProcess(f).To(NotHaveFlagOrHaveFlagWithValue("--kubelet-https", "true"))
	})

	It(fmt.Sprintf("[%d.%d.5] Ensure that the --insecure-bind-address argument is not set", index, subIndex), func() {
		ExpectProcess(f).To(NotHaveFlagOrHaveFlagWithValue("--insecure-bind-address", "127.0.0.1"))
	})

	It(fmt.Sprintf("[%d.%d.6] Ensure that the --insecure-port argument is set to 0", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagWithValue("--insecure-port", "0"))
	})

	It(fmt.Sprintf("[%d.%d.7] Ensure that the --secure-port argument is not set to 0", index, subIndex), func() {
		ExpectProcess(f).To(NotHaveFlagOrHaveFlagWithDifferentValue("--secure-port", "0"))
	})

	It(fmt.Sprintf("[%d.%d.8] Ensure that the --profiling argument is set to false", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagWithValue("--profiling", "false"))
	})

	It(fmt.Sprintf("[%d.%d.9] Ensure that the --repair-malformed-updates argument is set to false", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagWithValue("--repair-malformed-updates", "false"))
	})

	It(fmt.Sprintf("[%d.%d.10] Ensure that the admission control policy is not set to AlwaysAdmit", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagThatDoesNotContainValue("--admission-control", "AlwaysAdmit"))
	})

	It(fmt.Sprintf("[%d.%d.11] Ensure that the admission control policy is set to AlwaysPullImages", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagThatContainsValue("--admission-control", "AlwaysPullImages"))
	})

	It(fmt.Sprintf("[%d.%d.12] Ensure that the admission control policy is set to DenyEscalatingExec", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagThatContainsValue("--admission-control", "DenyEscalatingExec"))
	})

	It(fmt.Sprintf("[%d.%d.13] Ensure that the admission control policy is set to SecurityContextDeny", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagThatContainsValue("--admission-control", "SecurityContextDeny"))
	})

	It(fmt.Sprintf("[%d.%d.14] Ensure that the admission control policy is set to NamespaceLifecycle", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagThatContainsValue("--admission-control", "NamespaceLifecycle"))
	})

	It(fmt.Sprintf("[%d.%d.15] Ensure that the --audit-log-path argument is set as appropriate", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagWithAnyValue("--audit-log-path"))
	})

	It(fmt.Sprintf("[%d.%d.16] Ensure that the --audit-log-maxage argument is set to 30 or as appropriate", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagThatMatchesNumerically("--audit-log-maxage", ">=", 30))
	})

	It(fmt.Sprintf("[%d.%d.17] Ensure that the --audit-log-maxbackup argument is set to 10 or as appropriate", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagThatMatchesNumerically("--audit-log-maxbackup", ">=", 10))
	})

	It(fmt.Sprintf("[%d.%d.18] Ensure that the --audit-log-maxsize argument is set to 100 or as appropriate", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagThatMatchesNumerically("--audit-log-maxsize", ">=", 100))
	})

	It(fmt.Sprintf("[%d.%d.19] Ensure that the --authorization-mode argument is not set to AlwaysAllow", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagWithDifferentValue("--authorization-mode", "AlwaysAllow"))
	})

	It(fmt.Sprintf("[%d.%d.20] Ensure that the --token-auth-file parameter is not set", index, subIndex), func() {
		ExpectProcess(f).To(NotHaveFlag("--token-auth-file"))
	})

	It(fmt.Sprintf("[%d.%d.21] Ensure that the --kubelet-certificate-authority argument is set as appropriate", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagWithAnyValue("--kubelet-certificate-authority"))
	})

	It(fmt.Sprintf("[%d.%d.22] Ensure that the --kubelet-client-certificate and --kubelet-client-key arguments are set as appropriate", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagWithAnyValue("--kubelet-client-certificate"))
		ExpectProcess(f).To(HaveFlagWithAnyValue("--kubelet-client-key"))
	})

	It(fmt.Sprintf("[%d.%d.23] Ensure that the --service-account-lookup argument is set to true", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagWithValue("--service-account-lookup", "true"))
	})

	It(fmt.Sprintf("[%d.%d.24] Ensure that the admission control policy is set to PodSecurityPolicy", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagThatContainsValue("--admission-control", "PodSecurityPolicy"))
	})

	It(fmt.Sprintf("[%d.%d.25] Ensure that the --service-account-key-file argument is set as appropriate", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagWithAnyValue("--service-account-key-file"))
	})

	It(fmt.Sprintf("[%d.%d.26] Ensure that the --etcd-certfile and --etcd-keyfile arguments are set as appropriate", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagWithAnyValue("--etcd-certfile"))
		ExpectProcess(f).To(HaveFlagWithAnyValue("--etcd-keyfile"))
	})

	It(fmt.Sprintf("[%d.%d.27] Ensure that the admission control policy is set to ServiceAccount", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagThatContainsValue("--admission-control", "ServiceAccount"))
	})

	It(fmt.Sprintf("[%d.%d.28] Ensure that the --tls-cert-file and --tls-private-key-file arguments are set as appropriate", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagWithAnyValue("--tls-cert-file"))
		ExpectProcess(f).To(HaveFlagWithAnyValue("--tls-private-key-file"))
	})

	It(fmt.Sprintf("[%d.%d.29] Ensure that the --client-ca-file argument is set as appropriate", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagWithAnyValue("--client-ca-file"))
	})

	It(fmt.Sprintf("[%d.%d.30] Ensure that the --etcd-cafile argument is set as appropriate", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagWithAnyValue("--etcd-cafile"))
	})

	It(fmt.Sprintf("[%d.%d.31] Ensure that the --authorization-mode argument is set to Node", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagThatContainsValue("--authorization-mode", "Node"))
	})

	It(fmt.Sprintf("[%d.%d.32] Ensure that the admission control policy is set to NodeRestriction", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagThatContainsValue("--admission-control", "NodeRestriction"))
	})

	PIt(fmt.Sprintf("[%d.%d.34] Ensure that the encryption provider is set to aescbc", index, subIndex))

	It(fmt.Sprintf("[%d.%d.35] Ensure that the admission control policy is set to EventRateLimit", index, subIndex), func() {
		ExpectProcess(f).To(HaveFlagThatContainsValue("--admission-control", "EventRateLimit"))
		ExpectProcess(f).To(HaveFlagWithAnyValue("--admission-control-config-file"))
	})

	It(fmt.Sprintf("[%d.%d.36] Ensure that the AdvancedAuditing argument is not set to false", index, subIndex), func() {
		ExpectProcess(f).To(NotHaveFlagOrNotContainValue("--feature-gates", "AdvancedAuditing=false"))
	})

	PIt(fmt.Sprintf("[%d.%d.37] Ensure that the --request-timeout argument is set as appropriate", index, subIndex))
}
