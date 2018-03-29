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

const etcdProcessName = "etcd"

func Etcd(missingProcessFunc framework.MissingProcessHandlerFunc) {
	f := framework.New(etcdProcessName, missingProcessFunc)
	BeforeEach(f.BeforeEach)

	It("[1.5.1] Ensure that the --cert-file and --key-file arguments are set as appropriate [Scored]", func() {
		ExpectProcess(f).To(HaveFlagWithAnyValue("--cert-file"))
		ExpectProcess(f).To(HaveFlagWithAnyValue("--key-file"))
	})

	It("[1.5.2] Ensure that the --client-cert-auth argument is set to true [Scored]", func() {
		ExpectProcess(f).To(HaveFlagWithOptionalValue("--client-cert-auth", "true"))
	})

	It("[1.5.3] Ensure that the --auto-tls argument is not set to true [Scored]", func() {
		ExpectProcess(f).To(NotHaveFlagOrHaveFlagWithValue("--auto-tls", "false"))
	})

	It("[1.5.4] Ensure that the --peer-cert-file and --peer-key-file arguments are set as appropriate [Scored]", func() {
		ExpectProcess(f).To(HaveFlagWithAnyValue("--peer-cert-file"))
		ExpectProcess(f).To(HaveFlagWithAnyValue("--peer-key-file"))
	})

	It("[1.5.5] Ensure that the --peer-client-cert-auth argument is set to true [Scored]", func() {
		ExpectProcess(f).To(HaveFlagWithOptionalValue("--peer-client-cert-auth", "true"))
	})

	It("[1.5.6] Ensure that the --peer-auto-tls argument is not set to true [Scored]", func() {
		ExpectProcess(f).To(NotHaveFlagOrHaveFlagWithValue("--peer-auto-tls", "false"))
	})

	It("[1.5.7] Ensure that the --wal-dir argument is set as appropriate [Scored]", func() {
		ExpectProcess(f).To(HaveFlagWithAnyValue("--wal-dir"))
	})

	It("[1.5.8] Ensure that the --max-wals argument is set to 0 [Scored]", func() {
		ExpectProcess(f).To(HaveFlagWithValue("--max-wals", "0"))
	})

	PIt("[1.5.9] Ensure that a unique Certificate Authority is used for etcd [Not Scored]")
}
