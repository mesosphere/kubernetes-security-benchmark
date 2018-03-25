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

package cmd

import (
	"flag"

	"github.com/onsi/ginkgo/config"
	"github.com/spf13/cobra"

	"github.com/mesosphere/kubernetes-security-benchmark/pkg/cis"
	"github.com/mesosphere/kubernetes-security-benchmark/pkg/util"
)

// cisCmd represents the cis command
var cisCmd = &cobra.Command{
	Use:   "cis",
	Short: "Run Kubernetes CIS Benchmark tests",
	Long:  `Run Kubernetes CIS Benchmark tests.`,
	Run: func(cmd *cobra.Command, args []string) {
		util.RunTests("Kubernetes CIS Benchmark", cis.CISBenchmark)
	},
}

func init() {
	ginkgoFlagSet := flag.NewFlagSet("spec", flag.ContinueOnError)
	config.Flags(ginkgoFlagSet, "spec", false)
	config.DefaultReporterConfig.NoColor = true
	config.DefaultReporterConfig.Succinct = true
	config.DefaultReporterConfig.NoisySkippings = false
	config.DefaultReporterConfig.NoisyPendings = false
	ginkgoFlagSet.Lookup("spec.succinct").DefValue = "true"
	ginkgoFlagSet.Lookup("spec.noisySkippings").DefValue = "false"
	ginkgoFlagSet.Lookup("spec.noisyPendings").DefValue = "false"
	ginkgoFlagSet.Lookup("spec.noColor").DefValue = "true"
	cisCmd.Flags().AddGoFlagSet(ginkgoFlagSet)

	rootCmd.AddCommand(cisCmd)
}
