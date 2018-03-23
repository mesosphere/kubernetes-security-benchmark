package benchmark_test

import (
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/shirou/gopsutil/process"
)

func apiServer() {
	var apiServerProcess *process.Process

	processes, _ := process.Processes()
	for _, p := range processes {
		exe, _ := p.Exe()
		if filepath.Base(exe) == "kube-apiserver" {
			apiServerProcess = p
			break
		}
	}

	BeforeEach(func() {
		if apiServerProcess == nil {
			Fail("API server is not running")
		}
	})

	It("1.1.1 Ensure that the --anonymous-auth argument is set to false", func() {
		Expect(apiServerProcess.CmdlineSlice()).To(HaveFlagWithValue("--anonymous-auth", "false"))
	})

	It("1.1.2 Ensure that the --basic-auth-file argument is not set", func() {
		Expect(apiServerProcess.CmdlineSlice()).To(NotHaveFlag("--basic-auth-file"))
	})

	It("1.1.3 Ensure that the --insecure-allow-any-token argument is not set", func() {
		Expect(apiServerProcess.CmdlineSlice()).To(NotHaveFlag("--insecure-allow-any-token"))
	})

	It("1.1.4 Ensure that the --kubelet-https argument is set to true", func() {
		Expect(apiServerProcess.CmdlineSlice()).To(NotHaveFlagOrHaveFlagWithValue("--kubelet-https", "true"))
	})

	It("1.1.5 Ensure that the --insecure-bind-address argument is not set", func() {
		Expect(apiServerProcess.CmdlineSlice()).To(NotHaveFlagOrHaveFlagWithValue("--insecure-bind-address", "127.0.0.1"))
	})

	It("1.1.6 Ensure that the --insecure-port argument is set to 0", func() {
		Expect(apiServerProcess.CmdlineSlice()).To(HaveFlagWithValue("--insecure-port", "0"))
	})

	It("1.1.7 Ensure that the --secure-port argument is not set to 0", func() {
		Expect(apiServerProcess.CmdlineSlice()).To(NotHaveFlagOrHaveFlagWithDifferentValue("--secure-port", "0"))
	})

	It("1.1.8 Ensure that the --profiling argument is set to false", func() {
		Expect(apiServerProcess.CmdlineSlice()).To(HaveFlagWithValue("--profiling", "false"))
	})

	It("1.1.9 Ensure that the --repair-malformed-updates argument is set to false", func() {
		Expect(apiServerProcess.CmdlineSlice()).To(HaveFlagWithValue("--repair-malformed-updates", "false"))
	})

	It("1.1.10 Ensure that the admission control policy is not set to AlwaysAdmit", func() {
		Expect(apiServerProcess.CmdlineSlice()).To(HaveFlagThatDoesNotContainValue("--admission-control", "AlwaysAdmit"))
	})

	It("1.1.11 Ensure that the admission control policy is set to AlwaysPullImages", func() {
		Expect(apiServerProcess.CmdlineSlice()).To(HaveFlagThatContainsValue("--admission-control", "AlwaysPullImages"))
	})

	It("1.1.12 Ensure that the admission control policy is set to DenyEscalatingExec", func() {
		Expect(apiServerProcess.CmdlineSlice()).To(HaveFlagThatContainsValue("--admission-control", "DenyEscalatingExec"))
	})

	It("1.1.13 Ensure that the admission control policy is set to SecurityContextDeny", func() {
		Expect(apiServerProcess.CmdlineSlice()).To(HaveFlagThatContainsValue("--admission-control", "SecurityContextDeny"))
	})

	It("1.1.14 Ensure that the admission control policy is set to NamespaceLifecycle", func() {
		Expect(apiServerProcess.CmdlineSlice()).To(HaveFlagThatContainsValue("--admission-control", "NamespaceLifecycle"))
	})

	It("1.1.15 Ensure that the --audit-log-path argument is set as appropriate", func() {
		Expect(apiServerProcess.CmdlineSlice()).To(HaveFlagWithAnyValue("--audit-log-path"))
	})

	It("1.1.16 Ensure that the --audit-log-maxage argument is set to 30 or as appropriate", func() {
		Expect(apiServerProcess.CmdlineSlice()).To(HaveFlagThatMatchesNumerically("--audit-log-maxage", ">=", 30))
	})

	It("1.1.17 Ensure that the --audit-log-maxbackup argument is set to 10 or as appropriate", func() {
		Expect(apiServerProcess.CmdlineSlice()).To(HaveFlagThatMatchesNumerically("--audit-log-maxbackup", ">=", 10))
	})

	It("1.1.18 Ensure that the --audit-log-maxsize argument is set to 100 or as appropriate", func() {
		Expect(apiServerProcess.CmdlineSlice()).To(HaveFlagThatMatchesNumerically("--audit-log-maxsize", ">=", 100))
	})

	It("1.1.19 Ensure that the --authorization-mode argument is not set to AlwaysAllow", func() {
		Expect(apiServerProcess.CmdlineSlice()).To(HaveFlagWithDifferentValue("--authorization-mode", "AlwaysAllow"))
	})

	It("1.1.20 Ensure that the --token-auth-file parameter is not set", func() {
		Expect(apiServerProcess.CmdlineSlice()).To(NotHaveFlag("--token-auth-file"))
	})

	It("1.1.21 Ensure that the --kubelet-certificate-authority argument is set as appropriate", func() {
		Expect(apiServerProcess.CmdlineSlice()).To(HaveFlagWithAnyValue("--kubelet-certificate-authority"))
	})

	It("1.1.22 Ensure that the --kubelet-client-certificate and --kubelet-client-key arguments are set as appropriate", func() {
		Expect(apiServerProcess.CmdlineSlice()).To(HaveFlagWithAnyValue("--kubelet-client-certificate"))
		Expect(apiServerProcess.CmdlineSlice()).To(HaveFlagWithAnyValue("--kubelet-client-key"))
	})

	It("1.1.23 Ensure that the --service-account-lookup argument is set to true", func() {
		Expect(apiServerProcess.CmdlineSlice()).To(HaveFlagWithValue("--service-account-lookup", "true"))
	})

	It("1.1.24 Ensure that the admission control policy is set to PodSecurityPolicy", func() {
		Expect(apiServerProcess.CmdlineSlice()).To(HaveFlagThatContainsValue("--admission-control", "PodSecurityPolicy"))
	})

	It("1.1.25 Ensure that the --service-account-key-file argument is set as appropriate", func() {
		Expect(apiServerProcess.CmdlineSlice()).To(HaveFlagWithAnyValue("--service-account-key-file"))
	})

	It("1.1.26 Ensure that the --etcd-certfile and --etcd-keyfile arguments are set as appropriate", func() {
		Expect(apiServerProcess.CmdlineSlice()).To(HaveFlagWithAnyValue("--etcd-certfile"))
		Expect(apiServerProcess.CmdlineSlice()).To(HaveFlagWithAnyValue("--etcd-keyfile"))
	})

	It("1.1.27 Ensure that the admission control policy is set to ServiceAccount", func() {
		Expect(apiServerProcess.CmdlineSlice()).To(HaveFlagThatContainsValue("--admission-control", "ServiceAccount"))
	})

	It("1.1.28 Ensure that the --tls-cert-file and --tls-private-key-file arguments are set as appropriate", func() {
		Expect(apiServerProcess.CmdlineSlice()).To(HaveFlagWithAnyValue("--tls-cert-file"))
		Expect(apiServerProcess.CmdlineSlice()).To(HaveFlagWithAnyValue("--tls-private-key-file"))
	})

	It("1.1.29 Ensure that the --client-ca-file argument is set as appropriate", func() {
		Expect(apiServerProcess.CmdlineSlice()).To(HaveFlagWithAnyValue("--client-ca-file"))
	})

	It("1.1.30 Ensure that the --etcd-cafile argument is set as appropriate", func() {
		Expect(apiServerProcess.CmdlineSlice()).To(HaveFlagWithAnyValue("--etcd-cafile"))
	})

	It("1.1.31 Ensure that the --authorization-mode argument is set to Node", func() {
		Expect(apiServerProcess.CmdlineSlice()).To(HaveFlagThatContainsValue("--authorization-mode", "Node"))
	})

	It("1.1.32 Ensure that the admission control policy is set to NodeRestriction", func() {
		Expect(apiServerProcess.CmdlineSlice()).To(HaveFlagThatContainsValue("--admission-control", "NodeRestriction"))
	})

	PIt("1.1.34 Ensure that the encryption provider is set to aescbc")

	It("1.1.35 Ensure that the admission control policy is set to EventRateLimit", func() {
		Expect(apiServerProcess.CmdlineSlice()).To(HaveFlagThatContainsValue("--admission-control", "EventRateLimit"))
		Expect(apiServerProcess.CmdlineSlice()).To(HaveFlagWithAnyValue("--admission-control-config-file"))
	})

	It("1.1.36 Ensure that the AdvancedAuditing argument is not set to false", func() {
		Expect(apiServerProcess.CmdlineSlice()).To(NotHaveFlagOrNotContainValue("--feature-gates", "AdvancedAuditing=false"))
	})

	PIt("1.1.37 Ensure that the --request-timeout argument is set as appropriate")
}
