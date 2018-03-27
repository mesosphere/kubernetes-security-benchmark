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
	"github.com/onsi/ginkgo/config"
	"github.com/spf13/cobra"
)

// cisAPIServerCmd represents the cis control-plane api-server command
var cisAPIServerCmd = &cobra.Command{
	Use:   "api-server",
	Short: "Run the API server specific benchmarks",
	Long:  `Run the API server specific benchmarks.`,
	Run: func(cmd *cobra.Command, args []string) {
		config.GinkgoConfig.FocusString = `\[1\.1\]`
		cisCmd.Run(cmd, args)
	},
}

func init() {
	cisControlPlaneCmd.AddCommand(cisAPIServerCmd)
}
