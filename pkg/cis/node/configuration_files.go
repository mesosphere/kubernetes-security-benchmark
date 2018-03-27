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

package node

import (
	"fmt"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/mesosphere/kubernetes-security-benchmark/pkg/framework"
	. "github.com/mesosphere/kubernetes-security-benchmark/pkg/matcher"
	"github.com/mesosphere/kubernetes-security-benchmark/pkg/util"
)

func ConfigurationFiles(index, subIndex int, missingProcessFunc framework.MissingProcessHandlerFunc) {
	Context("", func() {
		kubelet := framework.New("kubelet", missingProcessFunc)
		BeforeEach(kubelet.BeforeEach)

		Context("", func() {
			var kubeconfigFilePath string

			BeforeEach(func() {
				kubeConfigFile, fileExists, err := util.FilePathFromFlag(kubelet.Process, "kubeconfig")
				Expect(err).NotTo(HaveOccurred())
				if !fileExists {
					Skip(fmt.Sprintf("%s does not exist", kubeConfigFile))
				}
				kubeconfigFilePath = kubeConfigFile
			})

			It(
				fmt.Sprintf("[%d.%d.1] Ensure that the kubelet.conf file permissions are set to 644 or more restrictive", index, subIndex),
				func() {
					Expect(kubeconfigFilePath).To(HavePermissionsNumerically("<=", os.FileMode(0644)))
				},
			)

			It(
				fmt.Sprintf("[%d.%d.2] Ensure that the kubelet.conf file ownership is set to root:root", index, subIndex),
				func() {
					Expect(kubeconfigFilePath).To(BeOwnedBy("root", "root"))
				},
			)
		})

		Context("", func() {
			const kubeletServiceFileDir = "/etc/systemd/system/kubelet.service.d"

			BeforeEach(func() {
				_, err := os.Stat(kubeletServiceFileDir)
				if os.IsNotExist(err) {
					Skip(kubeletServiceFileDir + " does not exist")
				}
				Expect(err).NotTo(HaveOccurred())
			})

			It(fmt.Sprintf("[%d.%d.3] Ensure that the kubelet service file permissions are set to 644 or more restrictive", index, subIndex), func() {
				err := filepath.Walk(kubeletServiceFileDir, func(path string, info os.FileInfo, err error) error {
					ExpectWithOffset(1, err).NotTo(HaveOccurred())
					if !info.IsDir() {
						Expect(path).To(HavePermissionsNumerically("<=", os.FileMode(0644)))
					}
					return nil
				})
				Expect(err).NotTo(HaveOccurred())
			})

			It(fmt.Sprintf("[%d.%d.4] Ensure that the kubelet service file ownership is set to root:root", index, subIndex), func() {
				err := filepath.Walk(kubeletServiceFileDir, func(path string, info os.FileInfo, err error) error {
					ExpectWithOffset(1, err).NotTo(HaveOccurred())
					if !info.IsDir() {
						Expect(path).To(BeOwnedBy("root", "root"))
					}
					return nil
				})
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})

	Context("", func() {
		kubeProxy := framework.New("kube-proxy", missingProcessFunc)
		BeforeEach(kubeProxy.BeforeEach)

		Context("", func() {
			var kubeconfigFilePath string

			BeforeEach(func() {
				kubeConfigFile, fileExists, err := util.FilePathFromFlag(kubeProxy.Process, "kubeconfig")
				Expect(err).NotTo(HaveOccurred())
				if !fileExists {
					Skip(fmt.Sprintf("%s does not exist", kubeConfigFile))
				}
				kubeconfigFilePath = kubeConfigFile
			})

			It(
				fmt.Sprintf("[%d.%d.5] Ensure that the proxy kubeconfig file permissions are set to 644 or more restrictive", index, subIndex),
				func() {
					Expect(kubeconfigFilePath).To(HavePermissionsNumerically("<=", os.FileMode(0644)))
				},
			)

			It(
				fmt.Sprintf("[%d.%d.6] Ensure that the proxy kubeconfig file ownership is set to root:root", index, subIndex),
				func() {
					Expect(kubeconfigFilePath).To(BeOwnedBy("root", "root"))
				},
			)
		})
	})

	Context("", func() {
		kubelet := framework.New("kubelet", missingProcessFunc)
		BeforeEach(kubelet.BeforeEach)

		Context("", func() {
			var clientCAFilePath string

			BeforeEach(func() {
				clientCAFile, fileExists, err := util.FilePathFromFlag(kubelet.Process, "client-ca-file")
				Expect(err).NotTo(HaveOccurred())
				if !fileExists {
					Skip(fmt.Sprintf("%s does not exist", clientCAFile))
				}
				clientCAFilePath = clientCAFile
			})

			It(
				fmt.Sprintf("[%d.%d.7] Ensure that the certificate authorities file permissions are set to 644 or more restrictive", index, subIndex),
				func() {
					Expect(clientCAFilePath).To(HavePermissionsNumerically("<=", os.FileMode(0644)))
				},
			)

			It(
				fmt.Sprintf("[%d.%d.8] Ensure that the client certificate authorities file ownership is set to root:root", index, subIndex),
				func() {
					Expect(clientCAFilePath).To(BeOwnedBy("root", "root"))
				},
			)
		})
	})
}
