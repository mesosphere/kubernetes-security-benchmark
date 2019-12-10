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
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/mesosphere/kubernetes-security-benchmark/pkg/framework"
	. "github.com/mesosphere/kubernetes-security-benchmark/pkg/ginkgo/matchers"
	"github.com/mesosphere/kubernetes-security-benchmark/pkg/util"
)

func configFilePermissionsContext(directory, fileName string, specFunc func(filePath string)) {
	Context("", func() {
		filePath := filepath.Join(directory, fileName)

		BeforeEach(func() {
			_, err := os.Stat(filePath)
			if err != nil {
				if os.IsNotExist(err) {
					Skip(fmt.Sprintf("%s does not exist", filePath))
				}
				Expect(err).NotTo(HaveOccurred())
			}
			Expect(filePath).To(BeARegularFile())
		})

		specFunc(filePath)
	})
}

func ConfigurationFiles(missingProcessFunc framework.MissingProcessHandlerFunc) {
	cwd, _ := os.Getwd()

	Context("", func() {
		kubelet := framework.New("kubelet", missingProcessFunc)
		BeforeEach(kubelet.BeforeEach)

		Context("", func() {
			var kubeletPodManifestPath string

			BeforeEach(func() {
				kmp, err := util.FlagValueFromProcess(kubelet.Process, "pod-manifest-path")
				Expect(err).NotTo(HaveOccurred())
				if kmp == "" {
					Skip(fmt.Sprintf("Flag --%s is unset", "pod-manifest-path"))
				}
				Expect(kmp).To(BeADirectory())
				kubeletPodManifestPath = kmp.(string)
			})

			configFilePermissionsContext(kubeletPodManifestPath, "kube-apiserver.yml", func(filePath string) {
				It("[1.4.1] Ensure that the API server pod specification file permissions are set to 644 or more restrictive [Scored]", func() {
					Expect(filePath).To(HavePermissionsNumerically("<=", os.FileMode(0644)))
				})

				It("[1.4.2] Ensure that the API server pod specification file ownership is set to root:root [Scored]", func() {
					Expect(filePath).To(BeOwnedBy("root", "root"))
				})
			})

			configFilePermissionsContext(kubeletPodManifestPath, "kube-controller-manager.yml", func(filePath string) {
				It("[1.4.3] Ensure that the controller manager pod specification file permissions are set to 644 or more restrictive [Scored]", func() {
					Expect(filePath).To(HavePermissionsNumerically("<=", os.FileMode(0644)))
				})

				It("[1.4.4] Ensure that the controller manager pod specification file ownership is set to root:root [Scored]", func() {
					Expect(filePath).To(BeOwnedBy("root", "root"))
				})
			})

			configFilePermissionsContext(kubeletPodManifestPath, "kube-scheduler.yml", func(filePath string) {
				It("[1.4.5] Ensure that the scheduler pod specification file permissions are set to 644 or more restrictive [Scored]", func() {
					Expect(filePath).To(HavePermissionsNumerically("<=", os.FileMode(0644)))
				})

				It("[1.4.6] Ensure that the scheduler pod specification file ownership is set to root:root [Scored]", func() {
					Expect(filePath).To(BeOwnedBy("root", "root"))
				})
			})

			configFilePermissionsContext(kubeletPodManifestPath, ""+
				""+
				".yaml", func(filePath string) {
				It("[1.4.7] Ensure that the etcd pod specification file permissions are set to 644 or more restrictive [Scored]", func() {
					Expect(filePath).To(HavePermissionsNumerically("<=", os.FileMode(0644)))
				})

				It("[1.4.8] Ensure that the etcd pod specification file ownership is set to root:root [Scored]", func() {
					Expect(filePath).To(BeOwnedBy("root", "root"))
				})
			})
		})

		Context("", func() {
			var cniConfDir string

			BeforeEach(func() {
				networkPlugin, err := util.FlagValueFromProcess(kubelet.Process, "network-plugin")
				Expect(err).NotTo(HaveOccurred())
				if networkPlugin == "" {
					Skip("Flag --network-plugin is unset")
				}
				if networkPlugin != "cni" {
					Skip("Flag --network-plugin is not set to cni")
				}

				ccd, err := util.FlagValueFromProcess(kubelet.Process, "cni-conf-dir")
				Expect(err).NotTo(HaveOccurred())
				if ccd == "" {
					ccd = "/etc/cni/net.d/"
				}
				Expect(ccd).To(BeADirectory())

				cniConfDir = ccd.(string)
			})

			It("[1.4.9] Ensure that the Container Network Interface file permissions are set to 644 or more restrictive [Scored]", func() {
				err := filepath.Walk(cniConfDir, func(path string, info os.FileInfo, err error) error {
					ExpectWithOffset(1, err).NotTo(HaveOccurred())
					if !info.IsDir() {
						Expect(path).To(HavePermissionsNumerically("<=", os.FileMode(0644)))
					}
					return nil
				})
				Expect(err).NotTo(HaveOccurred())
			})

			It("[1.4.10] Ensure that the Container Network Interface file ownership is set to root:root [Scored]", func() {
				err := filepath.Walk(cniConfDir, func(path string, info os.FileInfo, err error) error {
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
		etcd := framework.New("etcd", missingProcessFunc)
		BeforeEach(etcd.BeforeEach)

		Context("", func() {
			var etcdDataDir string

			BeforeEach(func() {
				edd, err := util.FlagValueFromProcess(etcd.Process, "data-dir")
				Expect(err).NotTo(HaveOccurred())
				if edd == "" {
					edd = filepath.Join(cwd, "data-dir")
				}
				Expect(edd).To(BeADirectory())

				etcdDataDir = edd.(string)
			})

			It("[1.4.11] Ensure that the etcd data directory permissions are set to 700 or more restrictive [Scored]", func() {
				expectedPerm := os.FileMode(0700) | os.ModeDir
				Expect(etcdDataDir).To(HavePermissionsNumerically("<=", expectedPerm))
			})

			It("[1.4.12] Ensure that the etcd data directory ownership is set to etcd:etcd [Scored]", func() {
				Expect(etcdDataDir).To(BeOwnedBy("etcd", "etcd"))
			})
		})
	})

	Context("", func() {
		const adminFilePath = "/etc/kubernetes/admin.conf"
		BeforeEach(func() {
			_, err := os.Stat(adminFilePath)
			if err != nil {
				if os.IsNotExist(err) {
					Skip(fmt.Sprintf("%s does not exist", adminFilePath))
				}
				Expect(err).NotTo(HaveOccurred())
			}
		})

		It("[1.4.13] Ensure that the admin.conf file permissions are set to 644 or more restrictive [Scored]", func() {
			Expect(adminFilePath).To(HavePermissionsNumerically("<=", os.FileMode(0644)))
		})

		It("[1.4.14] Ensure that the admin.conf file ownership is set to root:root [Scored]", func() {
			Expect(adminFilePath).To(BeOwnedBy("root", "root"))
		})
	})

	Context("", func() {
		var kubeconfigFilePath string

		assertKubeconfigFilePermissions := func() func() {
			return func() { Expect(kubeconfigFilePath).To(HavePermissionsNumerically("<=", os.FileMode(0644))) }
		}

		assertKubeconfigFileOwnership := func() func() {
			return func() { Expect(kubeconfigFilePath).To(BeOwnedBy("root", "root")) }
		}

		kubeconfigFlagContext := func(processName string, body func()) {
			Context("", func() {
				f := framework.New(processName, missingProcessFunc)
				BeforeEach(f.BeforeEach)

				Context("", func() {
					BeforeEach(func() {
						kubeConfigFile, fileExists, err := util.FilePathFromFlag(f.Process, "kubeconfig", cwd)
						Expect(err).NotTo(HaveOccurred())
						if !fileExists {
							Skip(fmt.Sprintf("%s does not exist", kubeConfigFile))
						}
						kubeconfigFilePath = kubeConfigFile
					})

					body()
				})
			})
		}

		kubeconfigFlagContext("kube-scheduler", func() {
			It(
				"[1.4.15] Ensure that the scheduler.conf file permissions are set to 644 or more restrictive [Scored]",
				assertKubeconfigFilePermissions(),
			)

			It(
				"[1.4.16] Ensure that the scheduler.conf file ownership is set to root:root [Scored]",
				assertKubeconfigFileOwnership(),
			)
		})

		kubeconfigFlagContext("kube-controller-manager", func() {
			It(
				"[1.4.17] Ensure that the controller-manager.conf file permissions are set to 644 or more restrictive [Scored]",
				assertKubeconfigFilePermissions(),
			)

			It(
				"[1.4.18] Ensure that the controller-manager.conf file ownership is set to root:root [Scored]",
				assertKubeconfigFileOwnership(),
			)
		})
	})

	Context("", func() {
		PIt("[1.4.19] Ensure that the Kubernetes PKI directory and file ownership is set to root:root [Scored]")

		PIt("[1.4.20] Ensure that the Kubernetes PKI certificate file permissions are set to 644 or more restrictive [Scored]")

		PIt("[1.4.21] Ensure that the Kubernetes PKI key file permissions are set to 600 [Scored]")
	})
}
