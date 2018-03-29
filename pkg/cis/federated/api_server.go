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

package federated

import (
	. "github.com/onsi/ginkgo"

	"github.com/mesosphere/kubernetes-security-benchmark/pkg/framework"
	. "github.com/mesosphere/kubernetes-security-benchmark/pkg/ginkgo/matchers"
)

const apiServerProcessName = "federation-apiserver"

func APIServer(index, subIndex int, missingProcessFunc framework.MissingProcessHandlerFunc) {
	f := framework.New(apiServerProcessName, missingProcessFunc)
	BeforeEach(f.BeforeEach)

	It("[3.1.1] Ensure that the --anonymous-auth argument is set to false [Scored]", func() {
		ExpectProcess(f).To(HaveFlagWithValue("--anonymous-auth", "false"))
	})

	It("[3.1.2] Ensure that the --basic-auth-file argument is not set [Scored]", func() {
		ExpectProcess(f).To(NotHaveFlag("--basic-auth-file"))
	})

	It("[3.1.3] Ensure that the --insecure-allow-any-token argument is not set [Scored]", func() {
		ExpectProcess(f).To(NotHaveFlag("--insecure-allow-any-token"))
	})

	It("[3.1.4] Ensure that the --insecure-bind-address argument is not set [Scored]", func() {
		ExpectProcess(f).To(NotHaveFlagOrHaveFlagWithValue("--insecure-bind-address", "127.0.0.1"))
	})

	It("[3.1.5] Ensure that the --insecure-port argument is set to 0 [Scored]", func() {
		ExpectProcess(f).To(HaveFlagWithValue("--insecure-port", "0"))
	})

	It("[3.1.6] Ensure that the --secure-port argument is not set to 0 [Scored]", func() {
		ExpectProcess(f).To(NotHaveFlagOrHaveFlagWithDifferentValue("--secure-port", "0"))
	})

	It("[3.1.7] Ensure that the --profiling argument is set to false [Scored]", func() {
		ExpectProcess(f).To(HaveFlagWithValue("--profiling", "false"))
	})

	It("[3.1.8] Ensure that the admission control policy is not set to AlwaysAdmit [Scored]", func() {
		ExpectProcess(f).To(HaveFlagThatDoesNotContainValue("--admission-control", "AlwaysAdmit"))
	})

	It("[3.1.9] Ensure that the admission control policy is set to NamespaceLifecycle [Scored]", func() {
		ExpectProcess(f).To(HaveFlagThatContainsValue("--admission-control", "NamespaceLifecycle"))
	})

	It("[3.1.10] Ensure that the --audit-log-path argument is set as appropriate [Scored]", func() {
		ExpectProcess(f).To(HaveFlagWithAnyValue("--audit-log-path"))
	})

	It("[3.1.11] Ensure that the --audit-log-maxage argument is set to 30 or as appropriate [Scored]", func() {
		ExpectProcess(f).To(HaveFlagThatMatchesNumerically("--audit-log-maxage", ">=", 30))
	})

	It("[3.1.12] Ensure that the --audit-log-maxbackup argument is set to 10 or as appropriate [Scored]", func() {
		ExpectProcess(f).To(HaveFlagThatMatchesNumerically("--audit-log-maxbackup", ">=", 10))
	})

	It("[3.1.13] Ensure that the --audit-log-maxsize argument is set to 100 or as appropriate [Scored]", func() {
		ExpectProcess(f).To(HaveFlagThatMatchesNumerically("--audit-log-maxsize", ">=", 100))
	})

	It("[3.1.14] Ensure that the --authorization-mode argument is not set to AlwaysAllow [Scored]", func() {
		ExpectProcess(f).To(HaveFlagWithDifferentValue("--authorization-mode", "AlwaysAllow"))
	})

	It("[3.1.15] Ensure that the --token-auth-file parameter is not set [Scored]", func() {
		ExpectProcess(f).To(NotHaveFlag("--token-auth-file"))
	})

	It("[3.1.16] Ensure that the --service-account-lookup argument is set to true [Scored]", func() {
		ExpectProcess(f).To(HaveFlagWithValue("--service-account-lookup", "true"))
	})

	It("[3.1.17] Ensure that the --service-account-key-file argument is set as appropriate [Scored]", func() {
		ExpectProcess(f).To(HaveFlagWithAnyValue("--service-account-key-file"))
	})

	It("[3.1.18] Ensure that the --etcd-certfile and --etcd-keyfile arguments are set as appropriate [Scored]", func() {
		ExpectProcess(f).To(HaveFlagWithAnyValue("--etcd-certfile"))
		ExpectProcess(f).To(HaveFlagWithAnyValue("--etcd-keyfile"))
	})

	It("[3.1.19] Ensure that the --tls-cert-file and --tls-private-key-file arguments are set as appropriate [Scored]", func() {
		ExpectProcess(f).To(HaveFlagWithAnyValue("--tls-cert-file"))
		ExpectProcess(f).To(HaveFlagWithAnyValue("--tls-private-key-file"))
	})
}
