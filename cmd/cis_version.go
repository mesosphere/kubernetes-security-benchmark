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
	"fmt"
	"github.com/mesosphere/kubernetes-security-benchmark/pkg/cis"

	"github.com/spf13/cobra"
)

// cisVersionCmd represents the cisversion command
var cisVersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version of the Kubernetes CIS Benchmark",
	Long:  `Prints the version of the Kubernetes CIS Benchmark.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Kubernetes CIS Benchmark v%s\n", cis.CISVersion)
	},
}

func init() {
	cisCmd.AddCommand(cisVersionCmd)
}
