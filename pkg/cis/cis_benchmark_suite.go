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
	"testing"

	. "github.com/onsi/ginkgo"
	corereporters "github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"

	"github.com/mesosphere/kubernetes-security-benchmark/pkg/framework"
	"github.com/mesosphere/kubernetes-security-benchmark/pkg/ginkgo/reporters"
)

const CISVersion = "1.2.0"

func CISBenchmark(missingProcFunc framework.MissingProcessHandlerFunc) func(*testing.T) {
	describeControlPlane(missingProcFunc)
	describeNode(missingProcFunc)
	describeFederatedDeployment(missingProcFunc)

	return func(t *testing.T) {
		RegisterFailHandler(Fail)
		junitReporter := corereporters.NewJUnitReporter("junit.xml")
		jsonReporter := reporters.NewJSONReporter("cis.json")
		RunSpecsWithDefaultAndCustomReporters(
			t,
			"Kubernetes CIS benchmark "+CISVersion,
			[]Reporter{junitReporter, jsonReporter},
		)
	}
}
