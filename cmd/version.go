// Copyright © 2018 Jimmi Dyson <jdyson@mesosphere.com>
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

	"github.com/spf13/cobra"

	"github.com/mesosphere/kubernetes-security-benchmark/pkg/version"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Kubernetes Security Benchmark",
	Long:  `Print the version number of Kubernetes Security Benchmark.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Kubernetes Security Benchmark %s BuildDate: %s\n", version.AppVersion, version.BuildDate)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
